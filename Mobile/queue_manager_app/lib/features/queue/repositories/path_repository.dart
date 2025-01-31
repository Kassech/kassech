import 'package:dio/dio.dart';
import 'package:queue_manager_app/config/const/api_constants.dart';

import '../../../core/services/api_service.dart';
import '../../../core/services/local_storage_service.dart';
import '../models/path_model.dart';

class PathRepository {
  final dio = ApiService.dio;
  final _storage = LocalStorageService();

  Future<List<PathModel>> fetchPaths() async {
    try {
      final response = await dio.get(
        ApiConstants.path,
        options: Options(
          headers: {
            'Authorization': 'Bearer ${_storage.getToken()}',
          },
        ),
      );

      return (response.data['data'] as List)
          .map((path) => PathModel.fromJson(path))
          .toList();
    } on DioException catch (_) {
      rethrow;
    } catch (e) {
      rethrow;
    }
  }

  Future<PathModel> fetchPathById(int id) async {
    try {
      final response = await dio.get(
        '${ApiConstants.path}/$id',
        options: Options(
          headers: {
            'Authorization': 'Bearer ${_storage.getToken()}',
          },
        ),
      );

      return PathModel.fromJson(response.data);
    } on DioException catch (_) {
      rethrow;
    } catch (e) {
      rethrow;
    }
  }
}
