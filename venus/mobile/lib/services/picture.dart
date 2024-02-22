import 'dart:io';
import 'package:path_provider/path_provider.dart';
import 'package:venus/services/models/folder.dart';
import 'package:venus/utils/utils.dart';

import 'package:venus/utils/image.dart';
import 'package:flutter/foundation.dart';
import 'package:path/path.dart';

import '../utils/logger.dart';
import 'database.dart';
import 'folder.dart';
import 'models/picture.dart';

Future<List<PictureModel>> searchPictures(String a) async {
  var sqlText = '''select pk, header, body, 
    simple_highlight(searches, 3, '[', ']') as highlight 
    from searches where body match jieba_query('国');''';

  var list = await DBHelper.instance.selectAsync(sqlText);

  debugPrint("list ${list.length}");

  // var searchList = List.generate(list.length, (i) {
  //   return SearchModel.fromJson(list[i]);
  // });

  return List.empty();
}

Future<List<PictureModel>> selectPicturesByFolder(FolderModel folder) async {
  if (folder.pk.isEmpty) return List.empty();

  var sqlText = '''select p.*, f.path path from pictures p 
    left join folders f on p.folder = f.pk where folder = ?
    order by basename limit 100;''';

  var list = await DBHelper.instance.selectAsync(sqlText, [folder.pk]);

  var pictureList = List.generate(list.length, (i) {
    return PictureModel.fromJson(list[i]);
  });

  return pictureList;
}

void insertPictureIfNotExists(PictureModel model, String folderPk) async {
  var commands = <String, List<Object?>>{};
  var sqlText =
  '''select pk from pictures where folder = ? and basename = ?;''';
  commands[sqlText] = [folderPk, model.basename];
  // 插入图片数据
  var sqlTextInsertPicture = '''
insert into pictures(pk, basename, folder)
values(?, ?, ?);
''';
  commands[sqlTextInsertPicture] = [model.pk, model.basename, folderPk];
  // 插入索引数据
  var sqlTextInsertSearches = '''
insert into searches(pk, header, body)
values(?, ?, ?);
''';
  commands[sqlTextInsertSearches] = [model.pk, 'picture', model.basename];
}

Future<void> scanPicturesWorker(FolderModel folderModel) async {
  if (folderModel.path.trim().isEmpty || folderModel.pk.trim().isEmpty) {
    return;
  }
  var realPath = folderModel.path;
  final dir = Directory(realPath);
  var isDirExist = await dir.exists();
  if (!isDirExist) {
    return;
  }
  var fileListStream = dir.list(recursive: false, followLinks: false);
  var filesCount = 0;
  await for (FileSystemEntity entity in fileListStream) {
    FileSystemEntityType type = await FileSystemEntity.type(entity.path);
    if (type == FileSystemEntityType.file) {
      var isPic = isImage(entity.path);

      if (!isPic) continue;
      filesCount += 1;
      var picPk = generateRandomString(16);
      var model =
          PictureModel(picPk, basename(entity.path), dirname(entity.path));
      // 保存到数据库
      insertPictureIfNotExists(model, folderModel.pk);
    }
  }
  await updateFilesCount(folderModel.pk, filesCount);
}

Future<bool> requestPermission() async {
  // final PermissionState ps = await PhotoManager.requestPermissionExtend();
  // if (ps.isAuth) {
  //   logger.d("requestPermission: $ps");
  //   return true;
  // } else {
  //   logger.d("requestPermission2222: $ps");
  //   return false;
  // }
  return true;
}

Future<List<PictureModel>> selectImages() async {

  Directory appDocDir = await getApplicationDocumentsDirectory();
  String appDocPath = appDocDir.path;
  var homeDir = appDocPath;
  debugPrint("fullPath: $homeDir");



  var imagePath = join(homeDir, "images");
  final dir = Directory(imagePath);
  var isDirExist = await dir.exists();
  if (!isDirExist) {
    return List.empty();
  }
  var fileListStream = dir.list(recursive: false, followLinks: false);

  var fileList = List<PictureModel>.empty(growable: true);
  await for (FileSystemEntity entity in fileListStream) {
    FileSystemEntityType type = await FileSystemEntity.type(entity.path);
    if (type == FileSystemEntityType.file) {
      var isPic = isImage(entity.path);

      if (!isPic) continue;
      var picPk = generateRandomString(16);
      var model = PictureModel(picPk, basename(entity.path), dirname(entity.path));
      fileList.add(model);
    }
  }
  return fileList;

}