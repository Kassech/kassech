import 'dart:async';
import 'dart:convert';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:web_socket_channel/status.dart' as status;

import 'local_storage_service.dart';

class WebSocketService {
  final String uri;
  late WebSocketChannel _channel;
  StreamSubscription<dynamic>? _subscription;
  final _storage = LocalStorageService();

  WebSocketService(this.uri) {
    _connect();
  }

  Future<void> _connect() async {
    final token = _storage.getToken();
    _channel = WebSocketChannel.connect(
      Uri.parse('$uri?token=$token'),
    );
  }

  Future<void> sendMessage(Map<String, dynamic> message) async {
    if (_channel.closeCode != null) await _connect();
    _channel.sink.add(jsonEncode(message));
  }

  Stream<dynamic> get messages => _channel.stream;

  Future<void> dispose() async {
    await _subscription?.cancel();
    if (_channel.closeCode == null) {
      _channel.sink.close(status.normalClosure);
    }
  }
}
