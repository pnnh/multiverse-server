#include "LibraryService.h"

#include "UserService.h"

#include <QSqlQuery>
#include <QtWidgets/QApplication>
#include <iostream>
#include <qdir.h>

LibraryService::LibraryService() {
  auto appDir = UserService::EnsureApplicationDirectory("/Polaris/Index");
  dbPath = appDir + "/Library.db";

  auto createSql = QString("create table if not exists libraries("
                           "uid varchar primary key not null,"
                           "name varchar(128) not null,"
                           "path varchar(512) not null)");
  if (!services::SqliteService::execute_query(dbPath, createSql)) {
    throw std::runtime_error("create table libraries error");
  }
  auto indexSql = QString(
      "create index if not exists index_libraries_path on libraries(path);");
  if (!services::SqliteService::execute_query(dbPath, indexSql)) {
    throw std::runtime_error("create index index_libraries_path error");
  }
}

QVector<LibraryModel> LibraryService::SelectLibraries() const {
  QVector<LibraryModel> libraryList;
  auto selectSql = QString("select * from libraries");

  auto sqlIterator = services::SqliteService::execute_query(dbPath, selectSql);

  while (sqlIterator->next()) {
    auto model = LibraryModel{.uid = sqlIterator->value("uid").toString(),
                              .name = sqlIterator->value("name").toString(),
                              .path = sqlIterator->value("path").toString()};
    libraryList.push_back(model);
  }
  return libraryList;
}

void LibraryService::InsertOrUpdateLibrary(
    const QVector<LibraryModel> &libraryList) {
  std::cout << "InsertOrUpdateLibrary: " << libraryList.size() << std::endl;

  const auto insertSql =
      QString("insert into libraries(uid, name, path)"
              " values(:uid, :name, :path)"
              " on conflict (uid) do update set name = :name;");
  for (const auto &library : libraryList) {
    QMap<QString, QVariant> parameters = {{
                                              ":uid",
                                              library.uid,
                                          },
                                          {":name", library.name},
                                          {":path", library.path}};
    if (!services::SqliteService::execute_query(dbPath, insertSql,
                                                parameters)) {
      throw std::runtime_error("create table libraries error");
    }
  }
}