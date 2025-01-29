import 'dart:async';
import 'dart:convert';

import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../repositories/passengers_repository.dart';

final passengerStreamProvider = AutoDisposeStreamNotifierProviderFamily<PassengersNotifier, int, int>(()  {
  final repository = PassengersRepository();
  repository.connect();

  return PassengersNotifier(repository);
});

class PassengersNotifier extends AutoDisposeFamilyStreamNotifier<int, int> {
  final PassengersRepository _repository;

  PassengersNotifier(this._repository);

  @override
  Stream<int> build(int arg) async* {
    final response = _repository.getPassengersCount(arg);

    await for (final value in response) {
      final data = jsonDecode(value);
      _passengerCountController.add(data['passengerCount']);
      yield data['passengerCount'];
    }
  }

  final _passengerCountController = StreamController<int>.broadcast();
  Stream<int> get passengerCountStream => _passengerCountController.stream;

  void incrementPassengerCount(int id) {
    _repository.incrementPassengerCount(id);
  }

  void decrementPassengerCount(int id) {
    _repository.decrementPassengerCount(id);
  }

  void disConnect() {
    _repository.disConnect();
  }
}