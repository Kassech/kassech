import 'dart:async';
import 'dart:convert';

import 'package:web_socket_channel/web_socket_channel.dart';

import '../../../config/const/api_constants.dart';
import '../../../core/services/local_storage_service.dart';

class PassengersRepository {
  late WebSocketChannel _channel;
  final _storage = LocalStorageService();
  final _controllers = <int, StreamController<int>>{};
  StreamSubscription<dynamic>? _socketSubscription;

  Future<void> connect() async {
    try {
      final token = await _storage.getToken();
      _channel = WebSocketChannel.connect(
        Uri.parse('${ApiConstants.passengers}?token=$token'),
      );

      // Set up global listener for all incoming messages
      _socketSubscription = _channel.stream.listen((message) {
        try {
          final data = jsonDecode(message);
          if (data['pathID'] != null && data['passengerCount'] != null) {
            final id = data['pathID'] as int;
            final count = data['passengerCount'] as int;
            _controllers[id]?.add(count);
          }
        } catch (e) {
          print('Error processing message: $e');
        }
      });
    } catch (e) {
      rethrow;
    }
  }

  Stream<int> getPassengersCount(int id) {
    try {
      // Create controller if it doesn't exist
      _controllers[id] ??= StreamController<int>.broadcast();

      // Request initial count
      _channel.sink.add(jsonEncode({"action": "getPassengers", "pathID": id}));

      return _controllers[id]!.stream;
    } catch (e) {
      rethrow;
    }
  }

  void incrementPassengerCount(int id) {
    try {
      _channel.sink.add(
        jsonEncode({"action": "increment", "pathID": id, "amount": 1}),
      );
    } catch (e) {
      rethrow;
    }
  }

  void decrementPassengerCount(int id) {
    try {
      _channel.sink.add(
        jsonEncode({"action": "decrement", "pathID": id, "amount": 1}),
      );
    } catch (e) {
      rethrow;
    }
  }

  void disConnect() {
    try {
      _socketSubscription?.cancel();
      _channel.sink.close();
      _controllers.forEach((_, controller) => controller.close());
      _controllers.clear();
    } catch (e) {
      rethrow;
    }
  }
}