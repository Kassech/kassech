import 'package:dio/dio.dart';
import 'package:flutter/foundation.dart';
import 'package:queue_manager_app/config/const/api_constants.dart';
import 'package:queue_manager_app/core/services/api_service.dart';
import 'package:queue_manager_app/core/services/local_storage_service.dart';
import 'package:queue_manager_app/features/owner/models/car_model.dart';

import '../../auth/models/user.dart';

class CarRepo {
  final dio = ApiService.dio;
  final storage = LocalStorageService();

  Future<List<Car>> fetchCars() async {
    try {
      final response = await dio.get(ApiConstants.getCars,
          options: Options(headers: {
            'Authorization': 'Bearer ${storage.getToken()}',
          }));
      return (response.data['data'] as List)
          .map((e) => Car.fromJson(e))
          .toList();
    } on DioException catch (e) {
      if (kDebugMode) {
        print("This is the $e");
      }
      rethrow;
    } catch (e) {
      throw Exception('Failed to load vehicle list');
    }
  }
}
