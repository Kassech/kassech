import 'dart:io';

import 'package:cookie_jar/cookie_jar.dart';
import 'package:dio/dio.dart';
import 'package:dio_cookie_manager/dio_cookie_manager.dart';
import 'package:flutter/foundation.dart';
import 'package:path_provider/path_provider.dart';

import '../../config/const/api_constants.dart';
import 'local_storage_service.dart';

class ApiService {
  static final Dio dio = Dio(
    BaseOptions(
      baseUrl: ApiConstants.apiBaseUrl,
      connectTimeout: const Duration(seconds: 10),
      receiveTimeout: const Duration(seconds: 10),
      headers: {'Content-Type': 'application/json'},
    ),
  );

  bool _isRefreshing = false;
  final _storage = LocalStorageService();

  static late PersistCookieJar _cookieJar;

  ApiService() {
    initializeDio();
  }

  /// Clear all cookies
  static Future<void> clearCookies() async {
    await _cookieJar.deleteAll();
  }

  /// Initialize the Dio
  Future<void> initializeDio() async {
    final cookiePath = await getCookiePath();
    _cookieJar = PersistCookieJar(storage: FileStorage(cookiePath));
    dio.interceptors.add(CookieManager(_cookieJar));

    dio.interceptors.add(InterceptorsWrapper(
      onRequest: (options, handler) async {
        if (options.headers.containsKey('Authorization')) {
          final accessToken = await _storage.getToken();
          if (accessToken != null) {
            options.headers['Authorization'] = 'Bearer $accessToken';
          }
        }
        return handler.next(options);
      },
      onResponse: (response, handler) {
        return handler.next(response);
      },
      onError: (DioException error, handler) async {
        if (error.response?.statusCode == 403) {
          /// Handle 403 error

          return handler.reject(error);
        }

        Response? errorResponse = error.response;
        String? errorMessage = '';

        if (errorResponse != null && errorResponse.data != null) {
          errorMessage = errorResponse.data['error'] ?? errorResponse.data['message'] ?? errorResponse.data;
        }

        bool isTokenExpired = error.response?.statusCode == 401 &&
            errorMessage!.toLowerCase().contains('token expired');

        if (isTokenExpired && !_isRefreshing) {
          print('Refreshing token...');
          _isRefreshing = true;
          try {
            final newToken = await _refreshTokenRequest();

            if (newToken != null) {
              await _storage.saveToken(newToken);

              final clonedRequest = await _retryRequest(error.requestOptions);
              return handler.resolve(clonedRequest);
            }
          } on DioException catch (error) {
            if (kDebugMode) {
              print('Token refresh error: ${error.response}');
            }
            return handler.reject(error);
          } catch (e) {
            if (kDebugMode) {
              print('Token refresh failed: $e');
            }
            return handler.reject(error);
          } finally {
            _isRefreshing = false;
          }
        }

        print('error after 401');
        /// handle other errors
        errorMessage = handleDioError(error);
        if (errorMessage != null && errorMessage.isNotEmpty) {
          print('error after $errorMessage');
          return handler.reject(
            DioException(
              requestOptions: error.requestOptions,
              response: error.response,
              error: errorMessage,
            ),
          );
        }

        print('error after ss $errorMessage');
        return handler.next(error);
      },
    ));
  }

  /// Get the path to store the cookies
  Future<String> getCookiePath() async {
    final directory = await getApplicationDocumentsDirectory();
    return '${directory.path}/cookies';
  }

  /// Refresh the access token
  Future<String?> _refreshTokenRequest() async {
    try {
      final response = await dio.post(ApiConstants.refreshToken);
      if (response.data != null && response.data['accessToken'] != null) {
        final newAccessToken = response.data['accessToken'];
        dio.options.headers['Authorization'] = 'Bearer $newAccessToken';
      }
      return response.data['accessToken'];
    } on DioException catch (error) {
      if (kDebugMode) {
        print('Token refresh error: ${error.response}');
      }
      return null;
    } catch (e) {
      if (kDebugMode) {
        print('Token refresh error: $e');
      }
      return null;
    }
  }

  /// Retry the original request after refreshing the token
  Future<Response> _retryRequest(RequestOptions requestOptions) async {
    dynamic data = requestOptions.data;

    if (data is FormData) {
      data = await cloneFormData(data);
    }

    final options = Options(
      method: requestOptions.method,
      headers: requestOptions.headers,
    );

    return dio.request(
      requestOptions.path,
      data: data,
      queryParameters: requestOptions.queryParameters,
      options: options,
    );
  }

  /// Clone the FormData object
  Future<FormData> cloneFormData(FormData original) async {
    final Map<String, dynamic> newMap = {};

    for (var entry in original.fields) {
      newMap[entry.key] = entry.value;
    }

    for (var fileEntry in original.files) {
      final MultipartFile file = fileEntry.value.clone();

      if (file.filename != null) {
        newMap[fileEntry.key] = file;
      }
    }

    return FormData.fromMap(newMap);
  }

  /// Handle errors and exceptions
  static String? handleDioError(DioException error) {
    String? errorMessage;

    switch (error.type) {
      case DioExceptionType.connectionTimeout:
        errorMessage = 'Connection timed out. Please try again later.';
        break;
      case DioExceptionType.sendTimeout:
        errorMessage = 'Request timed out while sending data. Please retry.';
        break;
      case DioExceptionType.receiveTimeout:
        errorMessage = 'Server response timed out. Check your connection.';
        break;
      case DioExceptionType.badResponse:
        if (error.response == null) {
          errorMessage = 'No response received from the server.';
        } else {
          if (error.response?.data is String) {
            errorMessage = error.response?.data.toString();
          } else if (error.response?.data is Map<String, dynamic>) {
            errorMessage = error.response?.data['message'] ??
                error.response?.data['error'];
          } else {
            errorMessage = error.response?.data.toString();
          }
        }
        break;
      case DioExceptionType.cancel:
        errorMessage = 'Request was cancelled. Please retry if needed.';
        break;
      case DioExceptionType.unknown:
        errorMessage =
            'Unexpected error occurred. Please check your connection.';
        break;
      default:
        if (error.error is SocketException) {
          errorMessage = 'No internet connection. Please check your network.';
        } else {
          errorMessage = 'Something went wrong. Please try again later.';
        }
    }

    if (errorMessage != null) {
      if (error.response != null &&
          error.response!.data
              .toString()
              .toLowerCase()
              .contains('token expired')) {
        return null;
      }
    }
    return errorMessage;
  }

  Future<void> sendTokensToBackend(
      String accessToken, String refreshToken) async {
    try {
      final response = await dio.post(ApiConstants.notification,
          data: {'token': accessToken, "device_id": "102934"});
      print('Notification response: ${response.data}');
    } catch (e) {
      print('Error sending tokens to backend: $e');
    }
  }

  Future<void> getNotifications(String accessToken) async {
    try {
      final response = await dio.post(ApiConstants.notification,
          data: {'token': 'abcde', "device_id": "102934"});

      print('Notifications response: ${response.data}');
    } catch (e) {
      print(e);
    }
  }
}
