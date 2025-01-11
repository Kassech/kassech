import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:queue_manager_app/core/util/token_storage.dart';

final webSocketProvider = StreamProvider.autoDispose<String>((ref) async* {
  final accessToken = await storage.read(key: 'accessToken');
  final channel = WebSocketChannel.connect(
    Uri.parse('ws://10.0.2.2:5000/ws/queue_manager?token=$accessToken'),
  );

  // Automatically close the WebSocket connection when the provider is disposed
  ref.onDispose(() => channel.sink.close());

  yield* channel.stream.map((event) => event as String);
});
