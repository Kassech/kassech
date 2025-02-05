import 'dart:async';
import 'dart:convert';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:web_socket_channel/status.dart' as status;

import 'local_storage_service.dart';

class WebSocketService {
  final String uri;
  late WebSocketChannel _channel;
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
    try {
      if (_channel.closeCode != null) await _connect();
      print('Sending message: ${jsonEncode(message)}');
      _channel.sink.add(jsonEncode(message));
    } on Exception catch (e) {
      // TODO
      print('socket exception $e');
    }
  }

  Stream<dynamic> get messages => _channel.stream;

  Future<void> dispose() async {
    if (_channel.closeCode == null) {
      _channel.sink.close(status.normalClosure);
    }
  }
}
