import 'package:dio/dio.dart';
import 'package:queue_manager_app/config/const/api_constants.dart';
import 'package:queue_manager_app/core/services/api_service.dart';
import 'package:queue_manager_app/core/services/local_storage_service.dart';
import 'package:queue_manager_app/features/owner/pages/delegate/delegation.dart';

class DriverRepo {
  final dio = ApiService.dio;
  final storage = LocalStorageService();

  // Fetch all drivers from the API
  Future<List<Driver>> fetchDrivers() async {
    try {
      final response = await dio.get(
        ApiConstants.getUsers, // Adjust the endpoint if necessary
        queryParameters: {
          'role': '2', // Driver role (this could vary depending on the API)
        },
        options: Options(
          headers: {
            'Authorization': 'Bearer ${storage.getToken()}',
          },
        ),
      );
      // Map the response to a list of Driver models
      return (response.data['data'] as List)
          .map((e) => Driver.fromJson(e))
          .toList();
    } on DioException catch (e) {
      print("DioException: ${e.toString()}");
      rethrow;
    } catch (e) {
      print("Error: ${e.toString()}");
      throw Exception('Failed to load drivers');
    }
  }

  // Search drivers by their name or phone number
  Future<List<Driver>> searchDrivers(String query) async {
    try {
      final response = await dio.get(
        ApiConstants.getUsers, // Adjust the endpoint if necessary
        queryParameters: {
          'role': '2', // Driver role (adjust as needed)
          'search':
              query, // Assuming there's a 'search' query parameter for filtering
        },
        options: Options(
          headers: {
            'Authorization': 'Bearer ${storage.getToken()}',
          },
        ),
      );
      // Map the response to a list of Driver models
      return (response.data['data'] as List)
          .map((e) => Driver.fromJson(e))
          .toList();
    } on DioException catch (e) {
      print("DioException: ${e.toString()}");
      rethrow;
    } catch (e) {
      print("Error: ${e.toString()}");
      throw Exception('Failed to search drivers');
    }
  }
}
