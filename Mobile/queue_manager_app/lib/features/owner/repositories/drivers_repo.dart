import 'package:dio/dio.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/config/const/api_constants.dart';
import 'package:queue_manager_app/core/services/api_service.dart';
import 'package:queue_manager_app/core/services/local_storage_service.dart';
import 'package:queue_manager_app/features/owner/models/car_model.dart';

import '../../auth/models/user.dart';

class DriversRepo {
  final dio = ApiService.dio;
  final storage = LocalStorageService();

  Future<List<User>> fetchDrivers(String? name) async {
    try {
      final response = await dio.get(
        ApiConstants.getDrivers,
        queryParameters: {
          'role': 2,
          'limit': 10,
          if (name != null) 'search': name,
        },
        options: Options(
          headers: {
            'Authorization': 'Bearer ${storage.getToken()}',
          },
        ),
      );

      print(response.data['data']);
      return (response.data['data'] as List)
          .map((e) => User.fromJson(e))
          .toList();
    } on DioException catch (_) {
      rethrow;
    } catch (e) {
      throw Exception('Failed to load vehicle list');
    }
  }

  /// hires a driver
  Future<void> hireDriver(int driverId) async {
    try {
      print("Hiring driver repo ");
      await dio.post(
        ApiConstants.hireDriver,
        data: {
          'driver_id': driverId,
        },
        options: Options(
          headers: {
            'Authorization': 'Bearer ${storage.getToken()}',
          },
        ),
      );
    } on DioException catch (_) {
      print("This is the error $_");
      rethrow;
    } catch (e, s) {
      throw Exception('Failed to hire driver');
    }
  }
}

final driversRepoProvider = Provider<DriversRepo>((ref) {
  return DriversRepo();
});
