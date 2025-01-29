import 'package:dio/dio.dart';

import '../../../core/services/api_service.dart';
import '../../../core/services/local_storage_service.dart';
import '../models/route_model.dart';

class RoutesRepository {
  final dio = ApiService.dio;
  final _storage = LocalStorageService();

  Future<List<RouteModel>> fetchRoutes() async {
    try {
      final response = await dio.get(
        '/routes',
        options: Options(
          headers: {
            'Authorization': 'Bearer ${await _storage.getToken()}',
          },
        ),
      );

      return (response.data as List)
          .map((route) => RouteModel.fromJson(route))
          .toList();
    } on DioException catch (_) {
      rethrow;
    } catch (e) {
      rethrow;
    }
  }

  Future<RouteModel> fetchRouteById(int id) async {
    try {
      final response = await dio.get(
        '/routes/$id',
        options: Options(
          headers: {
            'Authorization': 'Bearer ${await _storage.getToken()}',
          },
        ),
      );

      return RouteModel.fromJson(response.data);
    } on DioException catch (_) {
      rethrow;
    } catch (e) {
      rethrow;
    }
  }
}
