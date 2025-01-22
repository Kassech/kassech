import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/main.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

final webSocketProvider = StreamProvider.autoDispose<String>((ref) {
  final token = ref.watch(tokenProvider);

  if (token == null) {
    return const Stream<String>.empty(); // No token, no stream
  }

  final channel = WebSocketChannel.connect(
    Uri.parse('ws://10.0.2.2:5000/ws/queue_manager?token=$token'),
  );

  print('WebSocket connected with token: $token');

  ref.onDispose(() => channel.sink.close());

  return channel.stream.map((event) => event as String);
});
