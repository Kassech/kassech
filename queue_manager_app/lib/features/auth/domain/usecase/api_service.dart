import 'package:dio/dio.dart';
import 'package:shared_preferences/shared_preferences.dart';

class ApiService {
  final Dio _dio = Dio(BaseOptions(
    baseUrl: 'http://localhost:5000/api',
    connectTimeout: const Duration(seconds: 10),
    receiveTimeout: const Duration(seconds: 10),
    headers: {'Content-Type': 'application/json'},
  ));

  ApiService() {
    // Add interceptors
    _dio.interceptors.add(InterceptorsWrapper(
      onRequest: (options, handler) async {
        // Attach access token to headers
        final token = await _getAccessToken();
        if (token != null) {
          options.headers['Authorization'] = 'Bearer $token';
        }
        return handler.next(options);
      },
      onResponse: (response, handler) {
        // Handle successful responses
        return handler.next(response);
      },
      onError: (error, handler) async {
        final originalRequest = error.requestOptions;

        // If status code is 401 and it's not a login/register request
        if (error.response?.statusCode == 401 &&
            !_isLoginOrRegisterRoute(originalRequest.path) &&
            !originalRequest.extra.containsKey('_retry')) {
          originalRequest.extra['_retry'] = true;

          try {
            // Refresh the token
            final newToken = await _refreshAccessToken();
            if (newToken != null) {
              // Set the new token in the header
              originalRequest.headers['Authorization'] = 'Bearer $newToken';

              // Retry the request with the new token
              final response = await _dio.fetch(originalRequest);
              return handler.resolve(response);
            }
          } catch (error) {
            // If token refresh fails, reject the request
            return handler.reject(error as DioException);
          }
        }

        // For other errors, reject the request
        return handler.next(error);
      },
    ));
  }

  Future<String?> _getAccessToken() async {
    final prefs = await SharedPreferences.getInstance();
    return prefs.getString('accessToken');
  }

  Future<String?> _refreshAccessToken() async {
    try {
      final response = await _dio.post('/refresh');
      final newToken = response.data['access_token'];
      if (newToken == null) {
        throw Exception('Unauthorized');
      }
      // Save the new token to SharedPreferences
      final prefs = await SharedPreferences.getInstance();
      await prefs.setString('accessToken', newToken);
      return newToken;
    } catch (e) {
      print('Error refreshing access token: $e');
      throw e;
    }
  }

  bool _isLoginOrRegisterRoute(String path) {
    return path.contains('/login') || path.contains('/register');
  }

  Dio get dio => _dio; // Expose Dio instance for making requests
}
