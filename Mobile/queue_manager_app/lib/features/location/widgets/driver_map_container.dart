import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:latlong2/latlong.dart';

import '../../../core/util/helper.dart';
import '../providers/location_provider.dart';
import '../providers/route_provider.dart';
import '../repositories/route_repository.dart';

class DriverMapContainer extends ConsumerStatefulWidget {
  const DriverMapContainer(
      {super.key,
        required this.mapController,
        required this.start,
        required this.arrival});

  final MapController mapController;
  final LatLng start;
  final LatLng arrival;

  @override
  ConsumerState<DriverMapContainer> createState() => _DriverMapContainerState();
}

class _DriverMapContainerState extends ConsumerState<DriverMapContainer> {

  List<LatLng> pathRoute = [];

  @override
  void initState() {
    // TODO: implement initState
    super.initState();

    WidgetsBinding.instance.addPostFrameCallback((_) async {
      pathRoute = await ref.read(routeRepositoryProvider).getRoute(widget.start, widget.arrival);
      print('Path Route: $pathRoute');
    });
  }

  @override
  void dispose() {
    // TODO: implement dispose
    super.dispose();
    widget.mapController.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Consumer(builder: (context, ref, child) {
      final position = ref.read(locationNotifierProvider);

      ref.read(routeNotifierProvider.notifier).fetchRoute(
        LatLng(position?.latitude ?? 0, position?.longitude ?? 0),
        widget.start,
      );

      final driverRoute = ref.watch(routeNotifierProvider);

      final location = position?.latitude != null && position?.longitude != null
          ? LatLng(position!.latitude, position.longitude)
          : widget.start;

      return FlutterMap(
        mapController: widget.mapController,
        options: MapOptions(
          initialCenter: location,
          initialZoom: 15,
        ),
        children: [
          TileLayer(
            urlTemplate: "https://tile.openstreetmap.org/{z}/{x}/{y}.png",
            userAgentPackageName: 'com.example.queue_manager_app',
          ),
          MarkerLayer(
            markers: [
              Marker(
                point: widget.start,
                child: Icon(
                  Icons.location_on_sharp,
                  size: 40,
                  color: Colors.green.shade800,
                ),
              ),
              Marker(
                point: widget.arrival,
                child: Icon(
                  Icons.location_on_sharp,
                  size: 40,
                  color: Colors.blue,
                ),
              ),
              Marker(
                point:
                LatLng(position?.latitude ?? 0, position?.longitude ?? 0),
                child: Image.asset(
                  Helper.getAssetPath('taxi.png'),
                  width: 40,
                  height: 40,
                ),
              ),
            ],
          ),
          PolylineLayer(
            polylines: [
              Polyline(
                points: pathRoute,
                strokeWidth: 4.0,
                color: Colors.red.shade300,
              ),
              Polyline(
                points: driverRoute,
                strokeWidth: 4.0,
                color: Colors.blue.shade300,
              ),
            ],
          ),
        ],
      );
    });
  }
}
