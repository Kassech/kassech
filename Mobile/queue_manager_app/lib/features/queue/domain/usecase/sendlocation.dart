import 'dart:convert';

import 'package:dio/dio.dart';
import 'package:geolocator/geolocator.dart';
import 'package:queue_manager_app/core/util/geo_locator.dart';

import '../../../../core/services/api_service.dart';

Future<void> _sendLocationToBackend(Position position) async {
  final apiService = ApiService();
  // const String endpoint = 'location'; // Add your endpoint here
  // final String backendUrl =
  //     '${apiService.dio_baseUrl}$endpoint'; // Use baseUrl from ApiService

  try {
    final response = await apiService.dio_instance.post(
      '${apiService.dio_baseUrl}location',
      options: Options(headers: {'Content-Type': 'application/json'}),
      data: jsonEncode({
        'latitude': position.latitude,
        'longitude': position.longitude,
      }),
    );

    if (response.statusCode == 200) {
      print('Location sent successfully: ${response.data}');
    } else {
      print('Failed to send location: ${response.statusCode}');
    }
  } catch (error) {
    print('Error sending location: $error');
  }
}
Future<void> initializeLocation() async {
  try {
    Position position = await determinePosition();
    print('Current Position: ${position.latitude}, ${position.longitude}');
    await _sendLocationToBackend(position);
  } catch (error) {
    print('Error determining position: $error');
    if (error is DioException) {
      print('Dio Error: ${error.message}');
      print('Status Code: ${error.response?.statusCode}');
      print('Response Data: ${error.response?.data}');
    }
  }
}
