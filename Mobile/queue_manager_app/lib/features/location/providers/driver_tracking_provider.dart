// providers/driver_tracking_provider.dart
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:latlong2/latlong.dart';

import '../repositories/driver_tracking_repo.dart';

/// The state is a Map where the key is the vehicle_id and the value is the current LatLng.
final driverTrackingProvider = AutoDisposeStateNotifierProvider<DriverTrackingNotifier, Map<int, LatLng>>((ref) {
  final repository = ref.read(driverTrackingRepositoryProvider);
  return DriverTrackingNotifier(repository);
});

class DriverTrackingNotifier extends StateNotifier<Map<int, LatLng>> {
  final DriverTrackingRepository _repository;
  DriverTrackingNotifier(this._repository) : super({}) {
    _repository.driverUpdates.listen((data) {
      print('Received data: $data');
      // Expecting data like: {\"vehicle_id\":2,\"lat\":8.917243795189455,\"lon\":38.73322211348436,\"path_id\":null}
      final vehicleId = data['vehicle_id'];
      if (vehicleId != null && data['lat'] != null && data['lon'] != null) {
        final position = LatLng(data['lat'], data['lon']);
        state = {...state, vehicleId: position};
      }
    });
  }
}