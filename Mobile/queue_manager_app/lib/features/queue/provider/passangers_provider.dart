import 'dart:async';
import 'dart:convert';

import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../repositories/passengers_repository.dart';

final passengerControllerProvider = AutoDisposeStreamNotifierProvider<PassengersNotifier, List<Map<String, dynamic>>>(() {
  final repository = PassengersRepository();
  repository.connect();
  return PassengersNotifier(repository);
});


class PassengersNotifier extends AutoDisposeStreamNotifier<List<Map<String, dynamic>>> {
final PassengersRepository _repository;
final Map<int, StreamSubscription<int>> _subscriptions = {};

PassengersNotifier(this._repository);

@override
Stream<List<Map<String, dynamic>>> build() {
  // Initial state is an empty list
  return Stream.value([]);
}

void _updateSubscriptions(List<Map<String, dynamic>> currentList) {
  final currentIds = currentList.map((map) => map['id'] as int).toSet();

  // Cancel subscriptions for removed IDs
  _subscriptions.keys.where((id) => !currentIds.contains(id)).toList().forEach((id) {
    _subscriptions[id]?.cancel();
    _subscriptions.remove(id);
  });

  // Add new subscriptions for current IDs
  currentIds.where((id) => !_subscriptions.containsKey(id)).forEach((id) {
    final stream = _repository.getPassengersCount(id);
    _subscriptions[id] = stream.listen((count) {
      state = AsyncData([
        for (final passenger in state.value ?? [])
          if (passenger['id'] == id) {...passenger, 'count': count} else passenger
      ]);
    });
  });
}

void addPassenger(int id) {
  if (!state.value!.any((passenger) => passenger['id'] == id)) {
    state = AsyncData([...state.value!, {'id': id, 'count': 0}]);
    _updateSubscriptions(state.value!);
  }
}

void removePassenger(int id) {
  state = AsyncData(state.value!.where((passenger) => passenger['id'] != id).toList());
  _updateSubscriptions(state.value!);
}

void incrementPassengerCount(int id) {
  _repository.incrementPassengerCount(id);
}

void decrementPassengerCount(int id) {
  _repository.decrementPassengerCount(id);
}

@override
void dispose() {
  _subscriptions.values.forEach((sub) => sub.cancel());
  _subscriptions.clear();
  // super.dispose();
}

void disConnect() {
  _repository.disConnect();
}
}