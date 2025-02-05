import '../../../core/services/location_service.dart';

class LocationRepository {
  final LocationService _locationService;

  LocationRepository(this._locationService);

  Future<void> startTracking(int? vehicleId, int? pathId, int userId) async {
    return await _locationService.startLocationUpdates(vehicleId, pathId, userId);
  }

  void stopTracking() {
    _locationService.stopLocationUpdates();
  }
}
