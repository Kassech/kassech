// repositories/driver_tracking_repository.dart
import 'dart:convert';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../../../config/const/api_constants.dart';
import '../../../core/services/web_socket_service.dart';
import '../../queue/provider/passenger_provider.dart';

final driverTrackingRepositoryProvider = Provider<DriverTrackingRepository>((ref) {
  final webSocketService = ref.read(webSocketServiceProvider(ApiConstants.location));
  return DriverTrackingRepository(webSocketService);
});

class DriverTrackingRepository {
  final WebSocketService _webSocketService;

  DriverTrackingRepository(this._webSocketService);

  /// Stream of driver update messages parsed as a Map.
  Stream<Map<String, dynamic>> get driverUpdates => _webSocketService.messages.asBroadcastStream().map((message) {
    try {
      print('Received message: $message');
      return jsonDecode(message);
    } catch (e) {
      print('Error decoding message: $e');
      return {};
    }
  });

  /// Subscribe to driver updates.
  /// Depending on the provided parameters, sends the proper subscription message.
  void subscribeToDrivers({int? pathId, int? vehicleId}) {
    if (pathId != null) {
      _webSocketService.sendMessage({
        "action": "subscribe",
        "type": "path",
        "path_id": pathId,
      });
    }
    if (vehicleId != null) {
      _webSocketService.sendMessage({
        "action": "subscribe",
        "type": "vehicle",
        "vehicle_id": vehicleId,
      });
    }
  }

  /// Dispose the WebSocket when no longer needed
  void dispose() {
    _webSocketService.dispose();
  }
}
