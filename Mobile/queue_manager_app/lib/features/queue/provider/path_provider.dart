import 'dart:async';
import 'dart:convert';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/core/services/web_socket_service.dart';

final webSocketServiceProvider =
    Provider.family.autoDispose<WebSocketService, Uri>(
  (ref, uri) {
    print('Creating WebSocketService for $uri');
    final authToken =
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDA0MzY3MDYsInJvbGUiOlsiQWRtaW4iXSwidXNlcl9pZCI6MX0.jnxtSQ1hQw1WDGJpixr67p_m6wGWtJTyB3wec0wmBAo"; // Replace with actual token retrieval logic
    final service = WebSocketService(Uri.parse('$uri?token=$authToken'));
    ref.onDispose(() async {
      print('Disposing WebSocketService for $uri');
      await service.dispose();
    });
    return service;
  },
);

final pathNotifierProvider =
    StateNotifierProvider.autoDispose<PathNotifier, Map<String, int>>(
  (ref) {
    print('Initializing PathNotifier');
    return PathNotifier(
      ref.watch(webSocketServiceProvider(
          Uri.parse('ws://192.168.134.138:5000/ws/passengers'))),
    );
  },
);

class PathNotifier extends StateNotifier<Map<String, int>> {
  final WebSocketService _webSocketService;
  StreamSubscription<dynamic>? _subscription;

  PathNotifier(this._webSocketService) : super({}) {
    print('PathNotifier created');
    _listenToWebSocket();
  }

  void _listenToWebSocket() {
    print('Listening to WebSocket messages');
    _subscription = _webSocketService.messages.listen(
      (message) {
        print('Received message: $message');
        _handleMessage(message);
      },
      onError: (error) {
        print('WebSocket error: $error');
        _handleError(error);
      },
      onDone: () {
        print('WebSocket connection closed');
        _handleDisconnect();
      },
    );
  }

  void _handleMessage(dynamic message) {
    try {
      print('Handling message: $message');
      final data = jsonDecode(message);
      print('Parsed data: $data');
      state = {
        ...state,
        data['pathID'].toString(): int.parse(data['passengerCount'].toString())
      };
      print('Updated state: $state');
    } catch (e) {
      print('Error parsing message: $e');
    }
  }

  void _handleError(dynamic error) {
    print('Handling WebSocket error: $error');
  }

  void _handleDisconnect() {
    print('Handling WebSocket disconnection');
  }

  void updateCount(String pathId, int delta) {
    final message = {
      'action': delta > 0 ? 'increment' : 'decrement',
      'pathID': int.parse(pathId),
      'amount': delta.abs().toInt(),
    };

    print('Sending message: $message');
    _webSocketService.sendMessage(message);
  }

  @override
  void dispose() {
    print('Disposing PathNotifier');
    _subscription?.cancel();
    super.dispose();
  }
}
