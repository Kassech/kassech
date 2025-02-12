import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:latlong2/latlong.dart';
import 'package:queue_manager_app/features/location/widgets/general_map_container.dart';

import '../../models/car_model.dart';
import '../../widgets/driver_search_sheet.dart';

class CarLocation extends StatefulWidget {
  const CarLocation({super.key, required this.car});

  final Car car;

  static const String routeName = '/carLocationPage';

  @override
  State<CarLocation> createState() => _CarLocationState();
}

class _CarLocationState extends State<CarLocation> {
  final MapController mapController = MapController();
  @override
  Widget build(BuildContext context) {
    final themeData = Theme.of(context);
    return Scaffold(
      appBar: AppBar(
        title: Text(
          'Car Location',
          style: themeData.textTheme.headlineMedium
              ?.copyWith(fontWeight: FontWeight.bold),
        ),
        centerTitle: true,
        elevation: 0,
        backgroundColor: themeData.appBarTheme.backgroundColor,
      ),
      body: Column(
        children: [
          // Map Section
          Expanded(
            flex: 2,
            child: buildMap(context),
          ),
          // Car details section
          Expanded(
            flex: 3,
            child: Padding(
              padding: const EdgeInsets.all(16.0),
              child: buildCarDetails(themeData),
            ),
          ),
        ],
      ),
    );
  }

  Widget buildMap(BuildContext context) {
    // Using ClipRRect to round the top corners of the map area
    return ClipRRect(
      borderRadius: const BorderRadius.vertical(top: Radius.circular(16.0)),
      child: DriverTrackingScreen(startingLocation: null, mapController: mapController, vehicleId: widget.car.id,)
    );
  }

  Widget buildCarDetails(ThemeData themeData) {
    return Card(
      elevation: 4,
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(16.0),
      ),
      color: themeData.cardColor,
      child: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Header: Car image and make
            Row(
              children: [
                // Display the car picture if available, with a fallback icon
                ClipRRect(
                  borderRadius: BorderRadius.circular(12.0),
                  child: Image.network(
                    widget.car.carPicture,
                    width: 60,
                    height: 60,
                    fit: BoxFit.cover,
                    errorBuilder: (context, error, stackTrace) => Container(
                      width: 60,
                      height: 60,
                      color: themeData.dividerColor,
                      child: Icon(
                        Icons.directions_car,
                        color: themeData.iconTheme.color,
                      ),
                    ),
                  ),
                ),
                const SizedBox(width: 16),
                Expanded(
                  child: Text(
                    widget.car.make,
                    style: themeData.textTheme.headlineSmall?.copyWith(
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                ),
              ],
            ),
            const SizedBox(height: 20),
            Divider(color: themeData.dividerColor),
            const SizedBox(height: 16),
            // First row: License and Year
            Row(
              children: [
                Expanded(
                  child: _buildInfoItem(
                    themeData,
                    icon: Icons.confirmation_number,
                    label: 'License',
                    info: widget.car.licenseNumber,
                  ),
                ),
                const SizedBox(width: 16),
                Expanded(
                  child: _buildInfoItem(
                    themeData,
                    icon: Icons.date_range,
                    label: 'Year',
                    info: widget.car.year.toString(),
                  ),
                ),
              ],
            ),
            const SizedBox(height: 12),
            // Second row: VIN and Color
            Row(
              children: [
                Expanded(
                  child: _buildInfoItem(
                    themeData,
                    icon: Icons.vpn_key,
                    label: 'VIN',
                    info: widget.car.vin,
                  ),
                ),
                const SizedBox(width: 16),
                Expanded(
                  child: _buildInfoItem(
                    themeData,
                    icon: Icons.palette,
                    label: 'Color',
                    info: widget.car.color,
                  ),
                ),
              ],
            ),
            const SizedBox(height: 12),
            // Owner information (full width)
            _buildInfoItem(
              themeData,
              icon: Icons.person,
              label: 'Owner',
              info:
                  '${widget.car.owner.firstName} ${widget.car.owner.lastName}',
            ),
            const SizedBox(height: 20),
            // Action button centered at the bottom
            Center(
              child: ElevatedButton.icon(
                onPressed: () => showChangeDriverSheet(context, widget.car.id),
                icon: const Icon(Icons.person_outline),
                label: const Text("Change Driver"),
                style: ElevatedButton.styleFrom(
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(12.0),
                  ),
                  padding: const EdgeInsets.symmetric(
                    horizontal: 24.0,
                    vertical: 12.0,
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }

// Helper method to build each row item with an icon and descriptive text
  Widget _buildInfoItem(
    ThemeData themeData, {
    required IconData icon,
    required String label,
    required String info,
  }) {
    return Row(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Icon(
          icon,
          size: 20,
          color: themeData.iconTheme.color,
        ),
        const SizedBox(width: 8),
        Expanded(
          child: RichText(
            text: TextSpan(
              text: '$label: ',
              style: themeData.textTheme.bodyMedium?.copyWith(
                fontWeight: FontWeight.w600,
              ),
              children: [
                TextSpan(
                  text: info,
                  style: themeData.textTheme.bodyMedium,
                ),
              ],
            ),
          ),
        ),
      ],
    );
  }

  void showChangeDriverSheet(BuildContext context, int vehicleId) {
    final themeData = Theme.of(context);
    showModalBottomSheet(
      context: context,
      isScrollControlled: true,
      backgroundColor: themeData.bottomSheetTheme.backgroundColor ??
          themeData.scaffoldBackgroundColor,
      shape: const RoundedRectangleBorder(
        borderRadius: BorderRadius.vertical(top: Radius.circular(24.0)),
      ),
      builder: (context) {
        return DriverSearchSheet(vehicleId: vehicleId);
      },
    );
  }
}
