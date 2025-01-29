import 'package:web_socket_channel/io.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

import '../../config/const/api_constants.dart';
import 'local_storage_service.dart';

class WebSocketService {
  // Singleton instance
  static final WebSocketService _instance = WebSocketService._internal();

  WebSocketService._internal();

  // Factory constructor for accessing the singleton
  factory WebSocketService() => _instance;

  static late WebSocketChannel channel;
  final _storage = LocalStorageService();

  void connect({String? url}) async {
    try {
      final token = await _storage.getToken();
      final path = url ?? ApiConstants.status;
      print('Connecting to WebSocket at $path');
      channel = IOWebSocketChannel.connect(
        Uri.parse('$path?token=$token'),
      );
    } catch (e) {
      print('Error connecting to WebSocket: $e');
    }
  }

  // Send a message to the WebSocket
  void sendMessage(String message) {
    channel.sink.add(message);
  }

  // Disconnect from the WebSocket
  void disconnect() {
    channel.sink.close();
  }
}
