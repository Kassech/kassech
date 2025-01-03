// auth_service.dart
import 'package:dio/dio.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:queue_manager_app/core/network/network_service.dart';

class AuthenticationService {
  bool _isAuthenticated = false;

  bool get isAuthenticated => _isAuthenticated;

  void login() {
    _isAuthenticated = true;
  }

  void logout() {
    _isAuthenticated = false;
  }

  static const String _accessTokenKey = 'accessToken';
  static const String _refreshTokenKey = 'refreshToken';

  Future<void> saveTokens(String accessToken, String refreshToken) async {
    final dio = Dio();
    final response = await dio.post(
      '${networkService.baseUrl}refresh',
      data: {
        'accessToken': accessToken,
        'refreshToken': refreshToken,
      },
    );
    if (response.statusCode == 200) {
      print('Tokens saved in the database');
    } else if (response.statusCode == 403) {
      print('Could not save token: ${response.statusMessage}');
    }
    
    else {
      print('Error saving tokens in the database: ${response.statusMessage}');
    }

    final prefs = await SharedPreferences.getInstance();
    await prefs.setString(_accessTokenKey, accessToken);
    await prefs.setString(_refreshTokenKey, refreshToken);
  }

  Future<String?> getAccessToken() async {
    final prefs = await SharedPreferences.getInstance();
    return prefs.getString(_accessTokenKey);
  }

  Future<String?> getRefreshToken() async {
    final prefs = await SharedPreferences.getInstance();
    return prefs.getString(_refreshTokenKey);
  }

  Future<void> clearTokens() async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.remove(_accessTokenKey);
    await prefs.remove(_refreshTokenKey);
  }

  Future<bool> checkUserAuthentication() async {
    final token = await getAccessToken(); // Fetch the token
    if (token == null) return false; // User is not authenticated

    return true;
  }
}
