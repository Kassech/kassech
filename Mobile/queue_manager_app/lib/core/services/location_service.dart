import 'dart:async';
import 'dart:convert';

import 'package:flutter/foundation.dart';
import 'package:geolocator/geolocator.dart';
import 'package:queue_manager_app/core/services/local_storage_service.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

import '../../config/const/api_constants.dart';

class LocationService {
  WebSocketChannel? _channel;
  StreamSubscription<Position>? _positionSubscription;
  Position? _lastPosition;
  final LocalStorageService _localStorageService = LocalStorageService();

  final int minDistance = 1;

  Future<bool> _checkAndRequestPermissions() async {
    bool serviceEnabled = await Geolocator.isLocationServiceEnabled();
    if (!serviceEnabled) {
      bool serviceRequested = await Geolocator.openLocationSettings();
      if (!serviceRequested) {
        return false;
      }
    }

    LocationPermission permission = await Geolocator.checkPermission();

    if (permission == LocationPermission.denied) {
      permission = await Geolocator.requestPermission();
      if (permission == LocationPermission.denied) {
        return false;
      }
    }

    if (permission == LocationPermission.deniedForever) {
      await Geolocator.openAppSettings();
      return false;
    }

    if (permission == LocationPermission.whileInUse) {
      if (await Geolocator.requestPermission() != LocationPermission.always) {
        debugPrint(
            "Warning: Foreground services require 'Always' location access on Android 14+.");
      }
    }

    return permission == LocationPermission.always ||
        permission == LocationPermission.whileInUse;
  }

  Future<void> startLocationUpdates(
      int? vehicleId, int? pathId, int userId) async {
    if (!await _checkAndRequestPermissions()) {
      throw Exception('Location permissions are not granted');
    }

    final token = _localStorageService.getToken();

    try {
      _channel = WebSocketChannel.connect(
        Uri.parse('${ApiConstants.location}?token=$token'),
      );
    } catch (e) {
      if (kDebugMode) {
        print('Error connecting to WebSocket: $e');
      }
    }

    _positionSubscription = Geolocator.getPositionStream(
      locationSettings: AndroidSettings(
        accuracy: LocationAccuracy.high,
        foregroundNotificationConfig:  ForegroundNotificationConfig(
          notificationTitle: 'Location Service',
          notificationText: 'Location service is running ${_lastPosition?.latitude}, ${_lastPosition?.longitude}',
          setOngoing: true,
          enableWakeLock: true,
          enableWifiLock: true,
        ),
        distanceFilter: minDistance,
      ),
    ).listen((Position newPosition) {
      if (_lastPosition == null ||
          Geolocator.distanceBetween(
                _lastPosition!.latitude,
                _lastPosition!.longitude,
                newPosition.latitude,
                newPosition.longitude,
              ) >=
              minDistance) {
        _lastPosition = newPosition;

        print('Location update: ${newPosition.latitude}, ${newPosition.longitude}');
        final data = jsonEncode({
          if (vehicleId != null) "vehicle_id": vehicleId,
          "lat": newPosition.latitude,
          "lon": newPosition.longitude,
          if (pathId != null) "path_id": pathId,
          "created_at": DateTime.now().millisecondsSinceEpoch,
        });

        try {
          _channel?.sink.add(data);
        } catch (e) {
          if (kDebugMode) {
            print('Error sending location data: $e');
          }
        }
      }
    }, onError: (error) {
      if (kDebugMode) {
        print('Location error: $error');
      }
    });
  }

  void stopLocationUpdates() {
    _positionSubscription?.cancel();
    _channel?.sink.close();
  }
}
