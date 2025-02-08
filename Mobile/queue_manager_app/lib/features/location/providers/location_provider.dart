import 'dart:async';

import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:geolocator/geolocator.dart';

import '../../../core/services/location_service.dart';
import '../repositories/location_repository.dart';

final locationProvider = Provider<LocationService>((ref) => LocationService());
final locationRepositoryProvider = Provider<LocationRepository>((ref) => LocationRepository(ref.read(locationProvider)));

class LocationNotifier extends StateNotifier<Position?> {
  final LocationRepository _repository;
  StreamSubscription<Position>? _subscription;

  LocationNotifier(this._repository) : super(null);

  Future<void> startListening(int? vehicleId, int? pathId, int userId) async {
    await _repository.startTracking(vehicleId, pathId, userId);
    _subscription = Geolocator.getPositionStream(
      locationSettings: const LocationSettings(accuracy: LocationAccuracy.high, distanceFilter: 10),
    ).listen((position) {
      state = position;
    });
  }

  updateLocation(Position position) {
    state = position;
  }

  void stopListening() {
    _subscription?.cancel();
    _repository.stopTracking();
  }
}

final locationNotifierProvider = StateNotifierProvider<LocationNotifier, Position?>((ref) {
  final repo = ref.read(locationRepositoryProvider);
  return LocationNotifier(repo);
});
