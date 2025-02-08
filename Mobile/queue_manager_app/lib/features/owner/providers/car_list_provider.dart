import 'dart:async';

import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/features/owner/repositories/car_repo.dart';

import '../models/car_model.dart';

final carProvider = AsyncNotifierProvider<CarNotifier, List<Car>?>(() {
  final repository = CarRepo();
  return CarNotifier(repository);
});

class CarNotifier extends AsyncNotifier<List<Car>?> {
  final CarRepo _repository;

  CarNotifier(this._repository);

  @override
  Future<List<Car>?> build() async {
    return null;
  }

  FutureOr<List<Car>?> fetchCars() async {
    try {
      state = const AsyncLoading();
      final response = await _repository.fetchCars();
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

  // Future<Car?> fetchCarById(int id) async {
  //   try {
  //     state = const AsyncLoading();
  //     final response = await _repository.fetchCarById(id);
  //     state = AsyncData([response]);
  //     return response;
  //   } on DioException catch (e) {
  //     state = AsyncError(e.error ?? 'Something went wrong', e.stackTrace);
  //     return null;
  //   } catch (e, stackTrace) {
  //     state = AsyncError('Something went wrong', stackTrace);
  //     return null;
  //   }
  // }
}
