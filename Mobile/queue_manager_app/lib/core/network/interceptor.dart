// import 'package:dio/dio.dart';
// import 'package:queue_manager_app/core/util/token_storage.dart';

// class ApiClient {
//   late Dio dio;

//   ApiClient(String baseUrl) {
//     dio = Dio(
//       BaseOptions(
//         baseUrl: baseUrl,
//         connectTimeout: const Duration(seconds: 10),
//         receiveTimeout: const Duration(seconds: 10),
//       ),
//     );

//     dio.interceptors.add(InterceptorsWrapper(
//       onRequest: (options, handler) async {
//         // Add the access token to headers
//         final accessToken = await getAccessToken();
//         if (accessToken != null) {
//           options.headers['Authorization'] = 'Bearer $accessToken';
//         }
//         handler.next(options); // Proceed with the request
//       },
//       onResponse: (response, handler) {
//         // Handle successful responses
//         handler.next(response);
//       },
//       onError: (DioException error, handler) async {
//         // Handle token expiration
//         if (error.response?.statusCode == 401) {
//           // Try refreshing the token
//           final refreshToken = await getRefreshToken();
//           if (refreshToken != null) {
//             try {
//               final newTokens = await _refreshAccessToken(refreshToken);
//               if (newTokens != null) {
//                 final accessToken = newTokens['accessToken'];
//                 final refreshToken = newTokens['refreshToken'];

//                 // Save new tokens
//                 await saveTokens(accessToken, refreshToken);

//                 // Retry the failed request with the new access token
//                 final retryOptions = error.requestOptions;
//                 retryOptions.headers['Authorization'] = 'Bearer $accessToken';
//                 final response = await dio.fetch(retryOptions);
//                 return handler.resolve(response);
//               }
//             } catch (e) {
//               // Handle refresh token failure (e.g., logout user)
//               await clearTokens();
//               return handler.reject(error);
//             }
//           }
//         }
//         handler.next(error); // Pass the error if it's not 401 or refresh fails
//       },
//     ));
//   }

//   // Function to refresh access token
//   Future<Map<String, String>?> _refreshAccessToken(String refreshToken) async {
//     try {
//       final response = await dio.post('/auth/refresh', data: {
//         'refreshToken': refreshToken,
//       });

//       if (response.statusCode == 200) {
//         return {
//           'accessToken': response.data['accessToken'],
//           'refreshToken': response.data['refreshToken'],
//         };
//       }
//     } catch (e) {
//       // Log error if needed
//     }
//     return null; // Return null on failure
//   }
// }
