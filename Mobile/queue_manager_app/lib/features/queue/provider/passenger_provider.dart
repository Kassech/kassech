import 'dart:async';
import 'dart:convert';
import 'package:flutter/foundation.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/core/services/web_socket_service.dart';

import '../../../config/const/api_constants.dart';

final webSocketServiceProvider =
    Provider.family.autoDispose<WebSocketService, String>(
  (ref, uri) {
    final service = WebSocketService(uri);
    ref.onDispose(() async {
      service.dispose();
    });
    return service;
  },
);

final passengerNotifierProvider =
    StateNotifierProvider.autoDispose<PassengerNotifier, Map<String, int>>(
  (ref) {
    return PassengerNotifier(
      ref.watch(
        webSocketServiceProvider(
          ApiConstants.passengers,
        ),
      ),
    );
  },
);

class PassengerNotifier extends StateNotifier<Map<String, int>> {
  final WebSocketService _webSocketService;
  StreamSubscription<dynamic>? _subscription;

  PassengerNotifier(this._webSocketService) : super({}) {
    _listenToWebSocket();
  }

  void getInitialData(int pathID) {
    _webSocketService
        .sendMessage({"action": "getPassengers", "pathID": pathID});
  }

  void _listenToWebSocket() {
    _subscription = _webSocketService.messages.listen(
      (message) {
        _handleMessage(message);
      },
      onError: (error) {
        _handleError(error);
      },
      onDone: () {
        _handleDisconnect();
      },
    );
  }

  void _handleMessage(dynamic message) {
    try {
      final data = jsonDecode(message);
      state = {
        ...state,
        data['pathID'].toString(): int.parse(data['passengerCount'].toString())
      };
    } catch (e) {
      if (kDebugMode) {
        print('Error parsing message: $e');
      }
    }
  }

  void _handleError(dynamic error) {
    if (kDebugMode) {
      print('Handling WebSocket error: $error');
    }
  }

  void _handleDisconnect() {
    if (kDebugMode) {
      print('Handling WebSocket disconnection');
    }
  }

  void updateCount(String pathId, int delta) {
    final message = {
      'action': delta > 0 ? 'increment' : 'decrement',
      'pathID': int.parse(pathId),
      'amount': delta.abs().toInt(),
    };

    _webSocketService.sendMessage(message);
  }

  @override
  void dispose() {
    _subscription?.cancel();
    super.dispose();
  }
}
