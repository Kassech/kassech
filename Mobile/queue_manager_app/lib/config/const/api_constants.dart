class ApiConstants {
  /// Base url of the API
  static const String apiBaseUrl = 'http://192.168.114.190:5000/api';
  static const String socketBaseUrl = 'ws://192.168.114.190:5000/ws';

  /// Auth API endpoints
  static const String login = '$apiBaseUrl/login';
  static const String register = '$apiBaseUrl/register';
  static const String refreshToken = '$apiBaseUrl/refresh-token';

  /// Path API endpoints
  static const String path = '$apiBaseUrl/path';

  /// Notification API endpoint
  static const String notification = '$apiBaseUrl/notification';

  /// Socket Api endpoints
  static const String passengers = '$socketBaseUrl/passengers';
  static const String status = '$socketBaseUrl/status';
}
