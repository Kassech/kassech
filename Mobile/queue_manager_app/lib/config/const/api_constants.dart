class ApiConstants {
  /// Base url of the API
  static const String baseUrl = 'http://10.0.2.2:5000';
  static const String webSocketUrlPassenger =
      'ws://localhost:5000/ws/passengers?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDA0MzY3MDYsInJvbGUiOlsiQWRtaW4iXSwidXNlcl9pZCI6MX0.jnxtSQ1hQw1WDGJpixr67p_m6wGWtJTyB3wec0wmBAo';

  /// Auth API endpoints
  static const String login = '$baseUrl/api/login';
  static const String register = '$baseUrl/api/register';

  static const String passenger = webSocketUrlPassenger;
}
