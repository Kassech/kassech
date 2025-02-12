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
    final encoded = jsonEncode(message);
    try {
      // Check if the channel is closed; if so, reconnect.
      if (_channel.closeCode != null) {
        await _connect();
      }
      print('Sending message: $encoded');
      _channel.sink.add(encoded);
    } catch (e) {
      // If sending fails because the channel is closed, reconnect and retry.
      print('Error sending message: $e. Reconnecting and trying again.');
      await _connect();
      try {
        _channel.sink.add(encoded);
      } catch (e2) {
        print('Error sending message after reconnect: $e2');
      }
    }
  }

  // Convert the stream to broadcast so it can be listened to multiple times.
  Stream<dynamic> get messages => _channel.stream.asBroadcastStream();

  Future<void> dispose() async {
    if (_channel.closeCode == null) {
      _channel.sink.close(status.normalClosure);
    }
  }
}
