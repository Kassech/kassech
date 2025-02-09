import 'package:dio/dio.dart';
import 'package:queue_manager_app/config/const/api_constants.dart';
import 'package:queue_manager_app/core/services/api_service.dart';
import 'package:queue_manager_app/core/services/local_storage_service.dart';
import 'package:queue_manager_app/features/owner/models/car_model.dart';

class CarRepo {
  final dio = ApiService.dio;
  final storage = LocalStorageService();
  Future <List<Car>> fetchCars() async {
    try {
      final response = await dio.get(ApiConstants.getCars, options: Options(headers: {
        'Authorization': 'Bearer ${storage.getToken()}',
        
      }));
      return (response.data['data'] as List).map((e) => Car.fromJson(e)).toList();
    } 
    on DioException catch (e) {
      print("This is the "+ e.toString() + "");
      rethrow;
    }
    
    catch (e) {
            print("This is the " + e.toString() + "");

      throw Exception('Failed to load vehicle list');
    }
    
  }
}
