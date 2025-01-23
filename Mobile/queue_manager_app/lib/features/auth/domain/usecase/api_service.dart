import 'package:dio/dio.dart';
import 'package:queue_manager_app/core/util/token_storage.dart';
import 'package:shared_preferences/shared_preferences.dart';

class ApiService {
  final Dio _dio = Dio(BaseOptions(
    baseUrl: 'http://10.0.2.2:5000/api/',
    connectTimeout: const Duration(seconds: 10),
    receiveTimeout: const Duration(seconds: 10),
    headers: {'Content-Type': 'application/json'},
  ));

  // Getter method for the dio instance
  Dio get dio_instance => _dio;

  // Getter method for the base URL
  String get dio_baseUrl => _dio.options.baseUrl;
 

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
            final prefs = await SharedPreferences.getInstance();
            final refreshToken = prefs.getString("refreshToken");
            originalRequest.headers['Cookie'] = 'refresh_token=$refreshToken';

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

  Future<Response> login(String phoneNumber, String password) async {
    final formData = {
      'email_or_phone': phoneNumber,
      'password': password,
    };

    try {
      final response = await _dio.post('${dio_baseUrl}login', data: formData);
      print(response);
      return response;
    } catch (e) {
      print(e);
      rethrow;
    }
  }

  Future<String?> _getAccessToken() async {
    final prefs = await SharedPreferences.getInstance();
    return prefs.getString('accessToken');
  }

  Future<String?> _refreshAccessToken() async {
    try {
      final response = await _dio.post('${dio_baseUrl}refresh');
      final newToken = response.data['accessToken'];
      final newRefreshToken = response.data['refreshToken'];
      if (newToken == null || newRefreshToken == null) {
        throw Exception('Unauthorized');
      }
      // Save the new tokens to SharedPreferences
      final prefs = await SharedPreferences.getInstance();
      await prefs.setString('accessToken', newToken);
      await prefs.setString('refreshToken', newRefreshToken);
      return newToken;
    } catch (e) {
      print('Error refreshing access token: $e');
      rethrow;
    }
  }

  bool _isLoginOrRegisterRoute(String path) {
    return path.contains('/login') || path.contains('/register');
  }

  Future<void> sendTokensToBackend(
      String accessToken, String refreshToken) async {
    try {
      final response = await _dio.post('${dio_baseUrl}notification',
          data: {'token': accessToken, "device_id": "102934"});
      print('Notification response: ${response.data}');
    } catch (e) {
      print('Error sending tokens to backend: $e');
    }
  }

  Future<void> getNotifications(String accessToken) async {
    try {
      final response = await _dio.post('${dio_baseUrl}notifications',
          data: {'token': 'abcde', "device_id": "102934"});

      print('Notifications response: ${response.data}');
    } catch (e) {
      print(e);
    }
  }

  Future<bool> refreshAccessToken() async {
    final refreshToken = await getRefreshToken();
    if (refreshToken == null) {
      print('No refresh token found');
      return false;
    }

    try {
      final response = await _dio.post('${dio_baseUrl}refresh', data: {
        'refreshToken': refreshToken,
      });

      if (response.statusCode == 200) {
        final newAccessToken = response.data['accessToken'];
        // final newRefreshToken = response.data['refreshToken'];
        await updateToken(newAccessToken);
        print('Refreshed access token: $newAccessToken');
        // print('Refreshed refresh token: $newRefreshToken');
        return true;
      } else {
        print('Failed to refresh token, status code: ${response.statusCode}');
        return false;
      }
    } catch (e) {
      print('Error refreshing token: $e');
      return false;
    }
  }

  Future<String?> getRefreshToken() async {
    final prefs = await SharedPreferences.getInstance();
    final token = prefs.getString('refreshToken');
    print('Retrieved refresh token: $token');
    return token;
  }

  // Future<void> logout(BuildContext context) async {

  //    // Call the logout API (see Step 2)
  //   final success = await logoutApi();

  //   if (success) {
  //     // Clear local token storage
  //     await clearTokens();
  //     // await SharedPreferences.getInstance().remove(key: 'accessToken');

  //     // Navigate to the login screen
  //     Navigator.pushNamedAndRemoveUntil(
  //       context,
  //       '/login', // Your login page route
  //       (route) => false, // Remove all previous routes
  //     );
  //   } else {
  //     // Show an error message
  //     ScaffoldMessenger.of(context).showSnackBar(
  //       const SnackBar(content: Text('Failed to logout. Please try again.')),
  //     );
  //   }
  //   try {
  //     await SharedPreferences.getInstance().then((prefs) async {
  //       await prefs.remove('accessToken');
  //       await prefs.remove('refreshToken');
  //     });
  //     print('Logged out');
  //   } catch (e) {
  //     print('Error during logout: $e');
  //   }
  // }
Future<bool> logoutApi() async {
    try {
      final response = await _dio.post(
        'https://yourapi.com/logout',
        options: Options(
          validateStatus: (status) {
            return status! < 500; // Accept all status codes below 500
          },
        ),
      );

      if (response.statusCode == 200) {
        return true;
      } else if (response.statusCode == 403) {
        print('Logout failed: Unauthorized request');
        return false;
      } else {
        print('Logout failed: ${response.statusCode}');
        return false;
      }
    } catch (e) {
      print('Logout failed: $e');
      return false;
    }
  }
//   Future<void> clearTokens() async {
//     try {
//       await SharedPreferences.getInstance().then((prefs) async {
//         await prefs.remove('accessToken');
//       });
//       print('Cleared stored access token');
//     } catch (e) {
//       print('Error clearing stored access token: $e');
//     }
//   }

//     if (response.statusCode == 200) {
//       return true;
//     } else {
//       print('Failed to logout, status code: ${response.statusCode}');
//       return false;
//     }
//   } on DioError catch (dioError) {
//     print('DioError during logout: ${dioError.message}');
//     return false;
//   } catch (e) {
//     print('Unexpected error during logout: $e');
//     return false;
//   }
// }
  
  Future<void> saveTokens(String accessToken, String refreshToken) async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.setString('accessToken', accessToken);
    await prefs.setString('refreshToken', refreshToken);
    print('Saved access token: $accessToken');
    print('Saved refresh token: $refreshToken');
  }

  Future<void> updateToken(String accessToken) async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.setString('accessToken', accessToken);
    // await prefs.setString('refreshToken');
    print('Saved access token: $accessToken');
    // print('Saved refresh token: $refreshToken');
  }

  isAllowedToViewCars() {}
}


// Signup page api services
Future<void> saveToken(String token) async {
  await storage.write(key: 'accessToken', value: token);
}
