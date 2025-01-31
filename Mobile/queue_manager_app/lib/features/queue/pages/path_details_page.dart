import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:latlong2/latlong.dart';
import 'package:sliding_up_panel2/sliding_up_panel2.dart';

class PathDetailsPage extends StatefulWidget {
  const PathDetailsPage({super.key, required this.pathId});
  final int pathId;

  static const String routeName = '/pathDetailsPage';

  @override
  State<PathDetailsPage> createState() => _PathDetailsPageState();
}

class _PathDetailsPageState extends State<PathDetailsPage> {

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Route Details'),
      ),
      body: SlidingUpPanel(
        panelBuilder: () => _buildPanel(),
        body: FlutterMap(
          options: const MapOptions(
            initialCenter: LatLng(9.030093, 38.762791),
            initialZoom: 15,
          ),
          children: [
            openStreetMapTileLayer,
            MarkerLayer(
              markers: [
                Marker(
                    point: const LatLng(9.036151548242255, 38.7625160846566),
                    child: Icon(
                      Icons.location_on_sharp,
                      size: 60,
                      color: Colors.green[600],
                    ))
              ],
            )
          ],
        ),
      ),
    );
  }

  Widget _buildPanel() {
    ScrollController sc = ScrollController();
    return Container(
      color: Colors.white,
      padding: const EdgeInsets.all(16), // Add padding for better spacing
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          // Header
          Container(
            width: 100,
            height: 4, // Adjusted height for a better look
            color: Colors.black,
            margin: const EdgeInsets.only(bottom: 16), // Margin below the header
          ),
          // Panel Content
          Expanded(
            child: SingleChildScrollView(
              controller: sc,
              child: const 
              Column(
                crossAxisAlignment: CrossAxisAlignment.center,
                children: [
                  Text('Route Name: Route 1', style: TextStyle(fontSize: 18)),
                  SizedBox(height: 8), // Spacing between texts
                  Text('Route ID: R001', style: TextStyle(fontSize: 16)),
                  SizedBox(height: 8),
                  Text('Queue Count: 5', style: TextStyle(fontSize: 16)),
                  SizedBox(height: 16), // Extra space at the bottom
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}

TileLayer get openStreetMapTileLayer => TileLayer(
      urlTemplate: "https://tile.openstreetmap.org/{z}/{x}/{y}.png",
      userAgentPackageName: 'com.example.queue_manager_app',
    );

