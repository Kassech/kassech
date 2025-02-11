class ApiConstants {
  /// Base url of the API
  static const String apiBaseUrl = 'http://192.168.0.136:5000/api';
  static const String socketBaseUrl = 'ws://192.168.0.136:5000/ws';

  /// Auth API endpoints
  static const String login = '$apiBaseUrl/login';
  static const String register = '$apiBaseUrl/register';
  static const String refreshToken = '$apiBaseUrl/refresh-token';

  /// Path API endpoints
  static const String path = '$apiBaseUrl/path';

  /// Owner API endpoints
  static String carsByID(id) => '$apiBaseUrl/vehicles/$id';
  static String getDrivers = '$apiBaseUrl/users';
  static String hireDriver = '$apiBaseUrl/driver-deligation/assign';

//get cars
  static String getCars = '$apiBaseUrl/vehicles';

  /// Notification API endpoint
  static const String notification = '$apiBaseUrl/notification';

  /// Socket Api endpoints
  static const String passengers = '$socketBaseUrl/passengers';
  static const String status = '$socketBaseUrl/status';
  static const String location = '$socketBaseUrl/location';
  static const String getPathForDriver = '$socketBaseUrl/destination';
}
