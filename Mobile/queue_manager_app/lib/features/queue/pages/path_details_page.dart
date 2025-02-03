import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:latlong2/latlong.dart';
import 'package:sliding_up_panel2/sliding_up_panel2.dart';

import '../../location/providers/location_provider.dart';
import '../models/path_model.dart';
import '../provider/passenger_provider.dart';

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
        titleTextStyle: themeData.textTheme.bodyMedium,
        title: Text(widget.path.route.name),
      ),
      body: SlidingUpPanel(
        panelBuilder: () => _buildPanel(themeData),
        body: Consumer(
          builder: (context, ref, child) {
            final position = ref.watch(locationNotifierProvider);
            final location = position?.latitude != null && position?.longitude != null
                ? LatLng(position!.latitude, position.longitude)
                : LatLng(widget.path.route.startingLocation.latitude, widget.path.route.startingLocation.longitude);
            print('Position: $position');
            return FlutterMap(
              options: MapOptions(
                initialCenter: location,
                initialZoom: 15,
              ),
              children: [
                openStreetMapTileLayer,
                MarkerLayer(
                  markers: [
                    Marker(
                      point: LatLng(widget.path.route.startingLocation.latitude, widget.path.route.startingLocation.longitude),
                      child: Icon(
                        Icons.location_on_sharp,
                        size: 60,
                        color: Colors.green[600],
                      ),
                    ),
                    /// driver location
                    Marker(
                      point: LatLng(position?.latitude ?? 0, position?.longitude ?? 0),
                      child: Icon(
                        Icons.directions_car,
                        size: 60,
                        color: Colors.red[600],
                      ),
                    ),
                  ],
                )
              ],
            );
          }
        ),
      ),
    );
  }

  Widget _buildPanel(ThemeData themeData) {
    final size = MediaQuery.sizeOf(context);
    return Container(
      padding: const EdgeInsets.all(16),
      width: double.infinity,
      color: themeData.navigationBarTheme.backgroundColor,
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Container(
            alignment: Alignment.center,
            height: 4,
            color: Colors.black,
            margin: EdgeInsets.only(
                bottom: 16, left: size.width * 0.3, right: size.width * 0.3),
          ),
          // Panel Content
          Expanded(
            child: SingleChildScrollView(
              controller: sc,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                spacing: 10,
                children: [
                  buildPathDetailItem('Route:', widget.path.pathName),
                  buildPathDetailItem('Distance:', '${widget.path.distanceKm} km'),
                  buildPathDetailItem('Duration:', widget.path.estimatedTime.toString()),
                  buildPathDetailItem('Start:', widget.path.route.startingLocation.locationName),
                  buildPathDetailItem('End:', widget.path.route.arrivalLocation.locationName),
                  Consumer(
                    builder: (context, ref, child) {
                      ref
                          .watch(passengerNotifierProvider.notifier)
                          .getInitialData(widget.pathId);
                      final passengerCount = ref.watch(
                        passengerNotifierProvider.select(
                                (state) => state[widget.pathId.toString()] ?? 0),
                      );
                      return Row(
                        children: [
                          Text(
                            'Passenger Count: $passengerCount',
                            style: TextStyle(
                              fontSize: 16,
                              fontWeight: FontWeight.bold,
                            ),
                          ),
                          Spacer(),
                          IconButton(
                            icon: const Icon(Icons.remove),
                            onPressed: () => ref
                                .read(passengerNotifierProvider.notifier)
                                .updateCount(widget.pathId.toString(), -1),
                          ),
                          IconButton(
                            icon: const Icon(Icons.add),
                            onPressed: () => ref
                                .read(passengerNotifierProvider.notifier)
                                .updateCount(widget.pathId.toString(), 1),
                          ),
                        ],
                      );
                    },
                  ),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }

  Widget buildPathDetailItem(String label, String value) {
    return Column(
      children: [
        Row(
          spacing: 10,
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            Text(
              label,
              style: TextStyle(
                fontSize: 16,
                fontWeight: FontWeight.bold,
              ),
            ),
            Text(
              value,
              style: TextStyle(
                fontSize: 14,
              ),
            ),
          ],
        ),
        const Divider(),
      ],
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
