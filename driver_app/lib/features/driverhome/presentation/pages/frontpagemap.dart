import 'package:flutter/material.dart';
import 'package:circular_menu/circular_menu.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:latlong2/latlong.dart';

class DriverHome extends StatefulWidget {
  const DriverHome({super.key});

  @override
  State<DriverHome> createState() => _DriverHomeState();
}

class _DriverHomeState extends State<DriverHome> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(body: Stack(
      children: [
        FlutterMap(
          options: const MapOptions(
            center: LatLng(9.030093, 38.762791),
            minZoom: 15,
          ),
          children: [
            openStreetMapTileLayer,
            MarkerLayer(
              markers: [
                Marker(
                    point: LatLng(9.036151548242255, 38.7625160846566),
                    child: Icon(
                      Icons.location_on_sharp,
                      size: 60,
                      color: Colors.green[600],
                    ))
              ],
            ),
          ],
        ),
      ],
    ));
  }
}
