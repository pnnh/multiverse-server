// AUTO GENERATED FILE, DO NOT EDIT.
//
// Generated by `package:ffigen`.
// ignore_for_file: type=lint
import 'dart:ffi' as ffi;

class NativeLibrary {
  /// Holds the symbol lookup function.
  final ffi.Pointer<T> Function<T extends ffi.NativeType>(String symbolName)
      _lookup;

  /// The symbols are looked up in [dynamicLibrary].
  NativeLibrary(ffi.DynamicLibrary dynamicLibrary)
      : _lookup = dynamicLibrary.lookup;

  /// The symbols are looked up with [lookup].
  NativeLibrary.fromLookup(
      ffi.Pointer<T> Function<T extends ffi.NativeType>(String symbolName)
          lookup)
      : _lookup = lookup;

  int open_database(
    ffi.Pointer<ffi.Char> path,
  ) {
    return _open_database(
      path,
    );
  }

  late final _open_databasePtr =
      _lookup<ffi.NativeFunction<ffi.Int Function(ffi.Pointer<ffi.Char>)>>(
          'open_database');
  late final _open_database =
      _open_databasePtr.asFunction<int Function(ffi.Pointer<ffi.Char>)>();

  void hello_world() {
    return _hello_world();
  }

  late final _hello_worldPtr =
      _lookup<ffi.NativeFunction<ffi.Void Function()>>('hello_world');
  late final _hello_world = _hello_worldPtr.asFunction<void Function()>();

  int sum(
    int a,
    int b,
  ) {
    return _sum(
      a,
      b,
    );
  }

  late final _sumPtr =
      _lookup<ffi.NativeFunction<ffi.Int Function(ffi.Int, ffi.Int)>>('sum');
  late final _sum = _sumPtr.asFunction<int Function(int, int)>();

  int subtract(
    ffi.Pointer<ffi.Int> a,
    int b,
  ) {
    return _subtract(
      a,
      b,
    );
  }

  late final _subtractPtr = _lookup<
          ffi.NativeFunction<ffi.Int Function(ffi.Pointer<ffi.Int>, ffi.Int)>>(
      'subtract');
  late final _subtract =
      _subtractPtr.asFunction<int Function(ffi.Pointer<ffi.Int>, int)>();

  ffi.Pointer<ffi.Int> multiply(
    int a,
    int b,
  ) {
    return _multiply(
      a,
      b,
    );
  }

  late final _multiplyPtr = _lookup<
          ffi.NativeFunction<ffi.Pointer<ffi.Int> Function(ffi.Int, ffi.Int)>>(
      'multiply');
  late final _multiply =
      _multiplyPtr.asFunction<ffi.Pointer<ffi.Int> Function(int, int)>();

  void free_pointer(
    ffi.Pointer<ffi.Int> int_pointer,
  ) {
    return _free_pointer(
      int_pointer,
    );
  }

  late final _free_pointerPtr =
      _lookup<ffi.NativeFunction<ffi.Void Function(ffi.Pointer<ffi.Int>)>>(
          'free_pointer');
  late final _free_pointer =
      _free_pointerPtr.asFunction<void Function(ffi.Pointer<ffi.Int>)>();
}
