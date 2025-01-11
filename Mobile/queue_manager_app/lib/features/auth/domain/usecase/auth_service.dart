// import 'package:shared_preferences/shared_preferences.dart';
// import 'api_service.dart';

// class AuthService {
//   final ApiService _apiService = ApiService();

//   Future<String?> getAccessToken() async {
//     final prefs = await SharedPreferences.getInstance();
//     final token = prefs.getString('accessToken');
//     print('Retrieved access token: $token');
//     return token;
//   }

//   Future<String?> getRefreshToken() async {
//     final prefs = await SharedPreferences.getInstance();
//     final token = prefs.getString('refreshToken');
//     print('Retrieved refresh token: $token');
//     return token;
//   }

//   Future<void> saveTokens(String accessToken, String refreshToken) async {
//     final prefs = await SharedPreferences.getInstance();
//     await prefs.setString('accessToken', accessToken);
//     await prefs.setString('refreshToken', refreshToken);
//     print('Saved access token: $accessToken');
//     print('Saved refresh token: $refreshToken');
//   }

//   Future<bool> refreshAccessToken() async {
//     final refreshToken = await getRefreshToken();
//     if (refreshToken == null) {
//       print('No refresh token found');
//       return false;
//     }

//     try {
//       final response = await _apiService.dio_instance
//           .post('${_apiService.dio_baseUrl}refresh', data: {
//         'refreshToken': refreshToken,
//       });

//       if (response.statusCode == 200) {
//         final newAccessToken = response.data['accessToken'];
//         final newRefreshToken = response.data['refreshToken'];
//         await saveTokens(newAccessToken, newRefreshToken);
//         print('Refreshed access token: $newAccessToken');
//         print('Refreshed refresh token: $newRefreshToken');
//         return true;
//       } else {
//         print('Failed to refresh token, status code: ${response.statusCode}');
//         return false;
//       }
//     } catch (e) {
//       print('Error refreshing token: $e');
//       return false;
//     }
//   }
// }
