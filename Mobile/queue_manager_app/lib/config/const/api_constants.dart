class ApiConstants {
  /// Base url of the API
  static const String apiBaseUrl = 'http://192.168.134.138:5000/api';
  static const String socketBaseUrl = 'ws://192.168.134.138:5000/ws';

  /// Auth API endpoints
  static const String login = '$apiBaseUrl/login';
  static const String register = '$apiBaseUrl/register';
  static const String refreshToken = '$apiBaseUrl/refresh-token';

  /// Route API endpoints
  static const String routes = '$apiBaseUrl/routes';

  /// Notification API endpoint
  static const String notification = '$apiBaseUrl/notification';

  /// Socket Api endpoints
  static const String passengers = '$socketBaseUrl/passengers';
  static const String status = '$socketBaseUrl/status';
}