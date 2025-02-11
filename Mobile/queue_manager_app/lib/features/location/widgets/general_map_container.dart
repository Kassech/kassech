// screens/driver_tracking_screen.dart
import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:latlong2/latlong.dart';

import '../../auth/providers/auth_provider.dart';
import '../providers/driver_tracking_provider.dart';
import '../repositories/driver_tracking_repo.dart';

class DriverTrackingScreen extends ConsumerStatefulWidget {
  /// [startingLocation] is the point to which a polyline will be drawn from each driver.
  final LatLng? startingLocation;
  final int? pathId;
  final int? vehicleId;
  final MapController mapController;

  const DriverTrackingScreen(
      {super.key,
      required this.startingLocation,
      required this.mapController,
      this.pathId,
      this.vehicleId});

  @override
  ConsumerState<DriverTrackingScreen> createState() =>
      _DriverTrackingScreenState();
}

class _DriverTrackingScreenState extends ConsumerState<DriverTrackingScreen> {
  @override
  void initState() {
    super.initState();
    // After the widget builds, subscribe based on user role.
    Future.delayed(Duration.zero, () {
      final user = ref.read(authProvider).value;
      final repository = ref.read(driverTrackingRepositoryProvider);
      if (user != null) {
        if (user.roles.contains('QueueManager')) {
          // For QueueManagers, subscribe using a path id.
          repository.subscribeToDrivers(pathId: widget.pathId);
        } else if (user.roles.contains('Owner')) {
          // For Owners, subscribe using a vehicle id.
          repository.subscribeToDrivers(vehicleId: widget.vehicleId);
        }
      }
    });
  }

  @override
  void dispose() {
    // Optionally disconnect from WebSocket if needed.
    ref.read(driverTrackingRepositoryProvider).dispose();
    widget.mapController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final driverLocations = ref.watch(driverTrackingProvider);
    return FlutterMap(
      mapController: widget.mapController,
      options: MapOptions(
        initialCenter: widget.startingLocation ?? const LatLng(9.0331, 38.7617),
        initialZoom: 13,
      ),
      children: [
        TileLayer(
          urlTemplate: 'https://tile.openstreetmap.org/{z}/{x}/{y}.png',
          userAgentPackageName: 'com.example.queue_manager_app',
        ),
        // Markers for each driver's current location.
        MarkerLayer(
          markers: [
            if (widget.startingLocation != null)
              Marker(
                width: 40,
                height: 40,
                point: widget.startingLocation ?? const LatLng(9.0331, 38.7617),
                child: const Icon(
                  Icons.location_on,
                  color: Colors.green,
                  size: 40,
                ),
              ),
            ...driverLocations.entries.map((entry) {
              final vehicleId = entry.key;
              final position = entry.value;
              return Marker(
                width: 40,
                height: 40,
                point: position,
                child: Container(
                  decoration: BoxDecoration(
                    shape: BoxShape.circle,
                    color: Colors.blueAccent.withOpacity(0.8),
                    border: Border.all(color: Colors.white, width: 2),
                  ),
                  child: Center(
                    child: Text(
                      vehicleId.toString(),
                      style: const TextStyle(
                          color: Colors.white, fontWeight: FontWeight.bold),
                    ),
                  ),
                ),
              );
            }),
          ],
        ),
        // For each driver, draw a polyline from their location to the starting point.
        if (widget.startingLocation != null)
          PolylineLayer(
          polylines: driverLocations.entries.map((entry) {
            final position = entry.value;
            return Polyline(
              points: [position, widget.startingLocation ?? const LatLng(9.0331, 38.7617)],
              strokeWidth: 3,
              color: Colors.redAccent.withOpacity(0.6),
            );
          }).toList(),
        ),
      ],
    );
  }
}
