import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:latlong2/latlong.dart';

import '../repositories/route_repository.dart';

class RouteNotifier extends StateNotifier<List<LatLng>> {
  final RouteRepository _repository;

  RouteNotifier(this._repository) : super([]);

  Future<void> fetchRoute(LatLng start, LatLng arrival) async {
    try {
      final path = await _repository.getRoute(start, arrival);
      state = path;
    } catch (e) {
      print('Error fetching route: $e');
    }
  }
}

// Provide the notifier using Riverpod
final routeNotifierProvider =
StateNotifierProvider<RouteNotifier, List<LatLng>>((ref) {
  final repo = ref.read(routeRepositoryProvider);
  return RouteNotifier(repo);
});
