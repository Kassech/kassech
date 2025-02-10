import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../../../config/const/api_constants.dart';
import '../../../core/services/api_service.dart';
import '../../../core/services/local_storage_service.dart';
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
  Future<void> hireDriver(int driverId, int vehicleId) async {
    try {
      await dio.post(
        ApiConstants.hireDriver,
        data: {
          'driver_id': driverId,
          'vehicle_id': vehicleId,
        },
        options: Options(
          headers: {
            'Authorization': 'Bearer ${storage.getToken()}',
          },
        ),
      );
    } on DioException catch (_) {
      rethrow;
    } catch (e) {
      throw Exception('Failed to hire driver');
    }
  }
}

final driversRepoProvider = Provider<DriversRepo>((ref) {
  return DriversRepo();
});
