import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:latlong2/latlong.dart';
import 'package:sliding_up_panel2/sliding_up_panel2.dart';

import '../models/path_model.dart';

class PathDetailsPage extends StatefulWidget {
  const PathDetailsPage({super.key, required this.pathId, required this.path});

  final PathModel path;
  final int pathId;

  static const String routeName = '/pathDetailsPage';

  @override
  State<PathDetailsPage> createState() => _PathDetailsPageState();
}

class _PathDetailsPageState extends State<PathDetailsPage> {
  ScrollController sc = ScrollController();

  @override
  Widget build(BuildContext context) {
    final themeData = Theme.of(context);
    return Scaffold(
      appBar: AppBar(
        titleTextStyle: themeData.textTheme.bodySmall?.copyWith(),
        title: Row(
          spacing: 5,
          children: [
            Flexible(
              child: Text(widget.path.route.startingLocation.locationName),
            ),
            CustomPaint(
              size: Size(50, 30),
              painter: ArrowPainter(
                color: themeData.disabledColor,
              ),
            ),
            Flexible(
              child: Text(widget.path.route.arrivalLocation.locationName),
            ),
          ],
        ),
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
                  ),
                )
              ],
            )
          ],
        ),
      ),
    );
  }

  Widget _buildPanel() {
    return Container(
      padding: const EdgeInsets.all(16), // Add padding for better spacing
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          // Header
          Container(
            width: 100,
            height: 4, // Adjusted height for a better look
            color: Colors.black,
            margin:
                const EdgeInsets.only(bottom: 16), // Margin below the header
          ),
          // Panel Content
          Expanded(
            child: SingleChildScrollView(
              controller: sc,
              child: const Column(
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

class ArrowPainter extends CustomPainter {
  final Color color;

  ArrowPainter({required this.color});
  @override
  void paint(Canvas canvas, Size size) {
    final Paint paint = Paint()
      ..color = color
      ..strokeWidth = 1
      ..style = PaintingStyle.stroke;

    final Offset start = Offset(0, size.height / 2);
    final Offset end = Offset(size.width - 10, size.height / 2);

    // Draw the main line
    canvas.drawLine(start, end, paint);

    // Draw arrowhead
    final Paint arrowPaint = Paint()
      ..color = color
      ..strokeWidth = 1
      ..style = PaintingStyle.stroke;

    final Offset arrowTip = Offset(size.width, size.height / 2);
    final Offset arrowLeft = Offset(size.width - 10, size.height / 2 - 5);
    final Offset arrowRight = Offset(size.width - 10, size.height / 2 + 5);

    canvas.drawLine(arrowLeft, arrowTip, arrowPaint);
    canvas.drawLine(arrowRight, arrowTip, arrowPaint);
  }

  @override
  bool shouldRepaint(CustomPainter oldDelegate) => false;
}
