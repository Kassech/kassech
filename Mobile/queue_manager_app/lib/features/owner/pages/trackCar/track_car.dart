import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:latlong2/latlong.dart';
import 'package:queue_manager_app/core/theme/app_colors.dart';
import 'package:queue_manager_app/features/owner/models/car_model.dart';

class TrackCar extends StatefulWidget {
  const TrackCar({super.key});
  static const String routeName = '/trackCar';

  @override
  State<TrackCar> createState() => _TrackCarState();
}

class _TrackCarState extends State<TrackCar> {
  final MapController _mapController = MapController();
  final List<Marker> _markers = [];

  void _showCarDetails(Car car) {
    showModalBottomSheet(
      context: context,
      builder: (context) {
        return Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            mainAxisSize: MainAxisSize.min,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                car.make,
                style: const TextStyle(
                  fontSize: 22.0,
                  fontWeight: FontWeight.bold,
                  color: AppColors.black,
                ),
              ),
              const SizedBox(height: 8.0),
              Text(
                'License: ${car.licenseNumber}',
                style: const TextStyle(
                  fontSize: 16.0,
                  color: AppColors.lightGray,
                ),
              ),
              const SizedBox(height: 4.0),
              Text(
                'Year: ${car.year}',
                style: const TextStyle(
                  fontSize: 16.0,
                  color: AppColors.lightGray,
                ),
              ),
              const SizedBox(height: 4.0),
              Text(
                'Color: ${car.color}',
                style: const TextStyle(
                  fontSize: 16.0,
                  color: AppColors.lightGray,
                ),
              ),
              const SizedBox(height: 4.0),
              Text(
                'Status: ${car.status}',
                style: const TextStyle(
                  fontSize: 16.0,
                  color: AppColors.lightGray,
                ),
              ),
            ],
          ),
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    final car = ModalRoute.of(context)!.settings.arguments as Car;

    return Scaffold(
      appBar: AppBar(
        backgroundColor: AppColors.white,
        elevation: 4.0,
        leading: IconButton(
          icon: const Icon(Icons.menu, color: AppColors.black),
          onPressed: () {
            Scaffold.of(context).openDrawer();
          },
        ),
        title: const Text(
          'Track Car',
          style: TextStyle(
            color: AppColors.black,
            fontWeight: FontWeight.bold,
            fontSize: 24.0,
          ),
        ),
        actions: [
          IconButton(
            icon: const Icon(Icons.notifications, color: AppColors.black),
            onPressed: () {
              // Add notification functionality
            },
          ),
        ],
      ),
      body: FlutterMap(
        mapController: _mapController,
        options: MapOptions(
          initialCenter: LatLng(51.5, -0.09),
          minZoom: 13.0,
          onTap: (tapPosition, point) {
            setState(() {
              _markers.add(
                Marker(
                  point: point,
                  child: GestureDetector(
                    onTap: () {
                      _showCarDetails(car);
                    },
                    child: const Icon(
                      Icons.location_on,
                      color: Colors.red,
                      size: 40.0,
                    ),
                  ),
                ),
              );
            });
          },
        ),
        children: [
          TileLayer(
            urlTemplate: 'https://{s}.tile.openstreetmap.de/{z}/{x}/{y}.png',
            subdomains: ['a', 'b', 'c'],
          ),
          MarkerLayer(
            markers: _markers,
          ),
        ],
      ),
    );
  }
}
