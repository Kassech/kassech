import 'dart:convert';

import 'package:dio/dio.dart';
import 'package:geolocator/geolocator.dart';
import 'package:queue_manager_app/core/util/geo_locator.dart';

import '../../../../core/services/api_service.dart';

Future<void> _sendLocationToBackend(Position position) async {
  const String backendUrl = '';
  final dio = ApiService.dio;
  try {
    final response = await dio.post(
      backendUrl,
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
  }
}
