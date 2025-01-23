import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';

class CarLocation extends StatefulWidget {
  const CarLocation({super.key});

  @override
  State<CarLocation> createState() => _CarLocationState();
}

class _CarLocationState extends State<CarLocation> {
  @override
  Widget build(BuildContext context) {
    return  Scaffold(
      appBar: AppBar(title: const Text(
          'Car Location',
          style: TextStyle(color: Colors.black, fontWeight: FontWeight.bold),
        ),),
      body: openStreetMapTileLayer,
    );
  }
}
TileLayer get openStreetMapTileLayer => TileLayer(
      urlTemplate: "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png",
      userAgentPackageName: 'com.example.queue_manager_app',
      subdomains: const ['a', 'b', 'c'],
    );

