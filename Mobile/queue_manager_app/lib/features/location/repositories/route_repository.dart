import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_polyline_points/flutter_polyline_points.dart';
import 'package:latlong2/latlong.dart';

class RouteRepository {
  final Dio dio;

  RouteRepository(this.dio);

  Future<List<LatLng>> getRoute(LatLng start, LatLng arrival) async {
    final response = await dio.get(
      'http://router.project-osrm.org/route/v1/driving/${start.longitude},${start.latitude};${arrival.longitude},${arrival.latitude}?steps=true',
    );

    if (response.statusCode == 200) {
      final data = response.data;
      print('Route data: $data');

      // Extract the polyline string from the API response
      final polyline = data['routes'][0]['geometry'];

      // Decode the polyline using flutter_polyline_points
      PolylinePoints polylinePoints = PolylinePoints();
      List<PointLatLng> result = polylinePoints.decodePolyline(polyline);

      // Convert the result into LatLng objects
      List<LatLng> path = result.map((point) => LatLng(point.latitude, point.longitude)).toList();

      return path;
    } else {
      throw Exception('Failed to load route');
    }
  }
}

// Provide the repository using Riverpod
final routeRepositoryProvider = Provider<RouteRepository>((ref) {
  return RouteRepository(Dio());
});
