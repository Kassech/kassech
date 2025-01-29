import 'dart:convert';

import 'package:web_socket_channel/web_socket_channel.dart';

import '../../../config/const/api_constants.dart';
import '../../../core/services/local_storage_service.dart';

class PassengersRepository {
  late WebSocketChannel _channel;
  final _storage = LocalStorageService();

  Future<void> connect() async {
    try {
      final token = await _storage.getToken();

      _channel = WebSocketChannel.connect(
        Uri.parse('${ApiConstants.passengers}?token=$token'),
      );
    } on WebSocketChannelException catch (_) {
      rethrow;
    } catch (e) {
      rethrow;
    }
  }

  Stream<dynamic> getPassengersCount(int id) {
    try {
      _channel.sink.add(jsonEncode({"action": "getPassengers", "pathID": id}));
      return _channel.stream.asBroadcastStream();
    } on WebSocketChannelException catch (_) {
      rethrow;
    } catch (e) {
      rethrow;
    }
  }

  void incrementPassengerCount(int id) {
    try {
      _channel.sink.add(
        jsonEncode({"action": "increment", "pathID": id, "amount": 1}),
      );
    } on WebSocketChannelException catch (_) {
      rethrow;
    } catch (e) {
      rethrow;
    }
  }

  void decrementPassengerCount(int id) {
    try {
      _channel.sink.add(
        jsonEncode({"action": "decrement", "pathID": id, "amount": 1}),
      );
    } on WebSocketChannelException catch (_) {
      rethrow;
    } catch (e) {
      rethrow;
    }
  }

  void disConnect() {
    try {
      _channel.sink.close();
    } on WebSocketChannelException catch (_) {
      rethrow;
    } catch (e) {
      rethrow;
    }
  }
}
