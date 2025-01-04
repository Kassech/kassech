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

  static const String _accessTokenKey = 'access_token';
  static const String _refreshTokenKey = 'refresh_token';

  Future<void> saveTokens(String access_token, String refresh_token) async {
    final dio = Dio();
    try {
      final response = await dio.post(
        '${networkService.baseUrl}refresh',
        data: {
          'access_token': access_token,
          'refresh_token': refresh_token,
        },
      );
      if (response.statusCode == 200) {
        print('Tokens saved in the database');
      } else if (response.statusCode == 403) {
        print('Could not save token: ${response.statusMessage}');
      } else {
        print('Error saving tokens in the database: ${response.statusMessage}');
      }
    } catch (e) {
      print('Network error: $e');
    }

    final prefs = await SharedPreferences.getInstance();
    await prefs.setString(_accessTokenKey, access_token);
    await prefs.setString(_refreshTokenKey, refresh_token);
  }

  Future<String?> getAccessToken() async {
    final prefs = await SharedPreferences.getInstance();
    print(prefs.getString(_accessTokenKey));
    return prefs.getString(_accessTokenKey);
  }

  Future<String?> getRefreshToken() async {
    Dio dio = Dio();

    try{
      Response response = await dio.get('${networkService.baseUrl}refresh');
      if(response.statusCode == 200){

        String? refreshToken = response.headers['refresh_token']?.first;
        if (refreshToken != null) {
          print('Refresh token: $refreshToken');
         
        }else{
          print('Refresh token not found in headers');
        }

    }
    }catch(e){
      print('failed to get refresh token');
    }
    
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

    // Optionally, add token validation logic here
    try {
      final dio = Dio();
      final response = await dio.get(
        '${networkService.baseUrl}refresh',
        options: Options(
          headers: {'Cookie': 'refresh_token=$token'},
        ),
      );
      if (response.statusCode == 200) {
        // Access the refresh token from the cookies
        String? refreshToken;
        String? ftmKey;
        response.headers.forEach((name, values) {
          if (name.toLowerCase() == 'set-cookie') {
            values.forEach((value) {
              if (value.startsWith('refresh_token=')) {
                refreshToken = value.split(';')[0].split('=')[1];
              } else if (value.startsWith('ftm-key=')) {
                ftmKey = value.split(';')[0].split('=')[1];
              }
            });
          }
        });

        if (refreshToken != null && ftmKey != null) {
          print('Refresh Token: $refreshToken');
          print('FTM Key: $ftmKey');
          return true; // Token is valid
        } else {
          print('Refresh Token or FTM Key not found in cookies');
          return false; // Token is invalid
        }
      } else {
        return false; // Token is invalid
      }
    } catch (e) {
      print('Network error: $e');
      return false;
    }
  }

}
