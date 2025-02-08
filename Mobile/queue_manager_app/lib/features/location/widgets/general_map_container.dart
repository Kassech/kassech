import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:latlong2/latlong.dart';

class GeneralMapContainer extends StatefulWidget {
  const GeneralMapContainer({super.key, required this.mapController, required this.start, required this.arrival});

  final MapController mapController;
  final LatLng start;
  final LatLng arrival;

  @override
  State<GeneralMapContainer> createState() => _GeneralMapContainerState();
}

class _GeneralMapContainerState extends State<GeneralMapContainer> {
  @override
  Widget build(BuildContext context) {
    return const Placeholder();
  }
}
