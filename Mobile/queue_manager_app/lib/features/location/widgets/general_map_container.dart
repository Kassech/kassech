// import 'package:flutter/material.dart';
// import 'package:flutter_map/flutter_map.dart';
// import 'package:latlong2/latlong.dart';
//
// class GeneralMapContainer extends StatefulWidget {
//   const GeneralMapContainer({super.key, required this.mapController, required this.start, required this.arrival});
//
//   final MapController mapController;
//   final LatLng start;
//   final LatLng arrival;
//
//   @override
//   State<GeneralMapContainer> createState() => _GeneralMapContainerState();
// }
//
// class _GeneralMapContainerState extends State<GeneralMapContainer> {
//   @override
//   Widget build(BuildContext context) {
//     return FlutterMap(
//       mapController: widget.mapController,
//       options: MapOptions(
//         initialCenter: LatLng(0, 0),
//         initialZoom: 13.0,
//       ),
//       children: [
//         TileLayerOptions(
//           urlTemplate: 'https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png',
//           subdomains: ['a', 'b', 'c'],
//         ),
//         MarkerLayerOptions(
//           markers: [
//             Marker(
//               width: 80.0,
//               height: 80.0,
//               point: widget.start,
//               builder: (ctx) => const Icon(
//                 Icons.location_on,
//                 color: Colors.red,
//                 size: 30.0,
//               ),
//             ),
//             Marker(
//               width: 80.0,
//               height: 80.0,
//               point: widget.arrival,
//               builder: (ctx) => const Icon(
//                 Icons.location_on,
//                 color: Colors.green,
//                 size: 30.0,
//               ),
//             ),
//           ],
//         ),
//       ],
//     );
//   }
// }
