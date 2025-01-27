class ApiConstants {
  /// Base url of the API
  static const String baseUrl = 'http://192.168.0.136:5000';

  /// Auth API endpoints
  static const String login = '$baseUrl/api/login';
  static const String register = '$baseUrl/api/register';
  static const String refreshToken = '$baseUrl/api/refresh-token';

  /// Notification API endpoint
  static const String notification = '$baseUrl/api/notification';
}