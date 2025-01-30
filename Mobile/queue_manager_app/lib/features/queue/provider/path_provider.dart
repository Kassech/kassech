import 'dart:async';

import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../models/path_model.dart';
import '../repositories/routes_repository.dart';

final pathProvider =
    AsyncNotifierProvider<PathNotifier, List<PathModel>?>(() {
  final repository = PathRepository();
  return PathNotifier(repository);
});

class PathNotifier extends AsyncNotifier<List<PathModel>?> {
  final PathRepository _repository;

  PathNotifier(this._repository);

  @override
  Future<List<PathModel>?> build() async {
    return await fetchPaths();
  }

  Future<List<PathModel>?> fetchPaths() async {
    try {
      state = const AsyncLoading();
      final response = await _repository.fetchPaths();
      state = AsyncData(response);
      return response;
    } on DioException catch (e) {
      state = AsyncError(e.error ?? 'Something went wrong', e.stackTrace);
      return null;
    } catch (e, stackTrace) {
      state = AsyncError('Something went wrong', stackTrace);
      return null;
    }
  }

  Future<PathModel?> fetchPathById(int id) async {
    try {
      state = const AsyncLoading();
      final response = await _repository.fetchPathById(id);
      state = AsyncData([response]);
      return response;
    } on DioException catch (e) {
      state = AsyncError(e.error ?? 'Something went wrong', e.stackTrace);
      return null;
    } catch (e, stackTrace) {
      state = AsyncError('Something went wrong', stackTrace);
      return null;
    }
  }
}
