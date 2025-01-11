import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

final webSocketSinkProvider = StreamProvider.autoDispose<String>((ref) {
  final channel = WebSocketChannel.connect(
      Uri.parse('ws://10.0.2.2:5000/ws/queue_manager'));
  ref.onDispose(() {
    channel.sink.close();
  });
  return channel.stream.cast<String>();
});
