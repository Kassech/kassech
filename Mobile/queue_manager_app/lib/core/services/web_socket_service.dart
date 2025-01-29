import 'dart:async';
import 'dart:convert';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:web_socket_channel/status.dart' as status;

class WebSocketService {
  final Uri uri;
  late WebSocketChannel _channel;
  StreamSubscription<dynamic>? _subscription;

  WebSocketService(this.uri) {
    _connect();
  }

  void _connect() {
    _channel = WebSocketChannel.connect(uri);
  }

  void sendMessage(Map<String, dynamic> message) {
    if (_channel.closeCode != null) _connect();
    _channel.sink.add(jsonEncode(message));
  }

  Stream<dynamic> get messages => _channel.stream;

  Future<void> dispose() async {
    await _subscription?.cancel();
    if (_channel.closeCode == null) {
      _channel.sink.close(status.goingAway);
    }
  }
}
