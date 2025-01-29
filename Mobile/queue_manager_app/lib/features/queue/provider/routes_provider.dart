import 'dart:async';

import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../models/route_model.dart';
import '../repositories/routes_repository.dart';

final routesProvider =
    AsyncNotifierProvider<RoutesNotifier, List<RouteModel>?>(() {
  final repository = RoutesRepository();
  return RoutesNotifier(repository);
});

class RoutesNotifier extends AsyncNotifier<List<RouteModel>?> {
  final RoutesRepository _repository;

  RoutesNotifier(this._repository);

  @override
  Future<List<RouteModel>?> build() async {
    return await fetchRoutes();
  }

  Future<List<RouteModel>?> fetchRoutes() async {
    try {
      state = const AsyncLoading();
      final response = await _repository.fetchRoutes();
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

  Future<RouteModel?> fetchRouteById(int id) async {
    try {
      state = const AsyncLoading();
      final response = await _repository.fetchRouteById(id);
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
