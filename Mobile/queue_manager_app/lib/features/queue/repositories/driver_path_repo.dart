import 'dart:convert';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../../../config/const/api_constants.dart';
import '../../../core/services/web_socket_service.dart';
import '../models/path_model.dart';

final pathRepositoryProvider = Provider<DriverPathRepo>((ref) {
  return DriverPathRepo(ref);
});

class DriverPathRepo {
  final Ref ref;
  late WebSocketService _socketService;

  DriverPathRepo(this.ref) {
    print('DriverPathRepo created');
    _socketService = WebSocketService(ApiConstants.getPathForDriver);
  }

  /// Listen to WebSocket messages and decode them into PathModel
  Stream<PathModel?> get pathStream => _socketService.messages.map((event) {
    try {
      print('Received WebSocket message: $event');
      final data = jsonDecode(event);
      return PathModel.fromJson(data); 
    } catch (e) {
      print("Error parsing WebSocket message: $e");
      return null;
    }
  });

  /// Dispose the WebSocket when no longer needed
  void dispose() {
    _socketService.dispose();
  }
}
