import 'dart:math';

import 'package:cached_network_image/cached_network_image.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/core/util/ui_utils.dart';
import '../../../core/theme/app_colors.dart';
import '../../auth/models/user.dart';
import '../models/hire_driver_params.dart';
import '../providers/drivers_provider.dart';
import '../providers/hire_driver_provider.dart';

class DriverSearchSheet extends ConsumerStatefulWidget {
  const DriverSearchSheet({super.key, required this.vehicleId});

  final int vehicleId;

  @override
  ConsumerState<DriverSearchSheet> createState() => _DriverSearchSheetState();
}

class _DriverSearchSheetState extends ConsumerState<DriverSearchSheet> {
  final TextEditingController searchController = TextEditingController();

  @override
  void initState() {
    super.initState();
    searchController.addListener(_onSearchChanged);
  }

  @override
  void dispose() {
    searchController.removeListener(_onSearchChanged);
    searchController.dispose();
    super.dispose();
  }

  void _onSearchChanged() {
    ref
        .read(driversNotifierProvider(searchController.text).notifier)
        .refreshDrivers(searchController.text);
  }

  @override
  Widget build(BuildContext context) {
    final themeData = Theme.of(context);
    final asyncDrivers =
        ref.watch(driversNotifierProvider(searchController.text));
    return Container(
      height: MediaQuery.sizeOf(context).height * 0.7,
      padding: MediaQuery.of(context).viewInsets,
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            // Drag indicator for a modern look.
            Container(
              width: 50,
              height: 5,
              decoration: BoxDecoration(
                color: themeData.dividerColor,
                borderRadius: BorderRadius.circular(10),
              ),
            ),
            const SizedBox(height: 16.0),
            Text(
              "Select a Driver",
              style: themeData.textTheme.headlineSmall,
            ),
            const SizedBox(height: 16.0),
            // Search field with an icon and themed styling.
            TextField(
              controller: searchController,
              decoration: InputDecoration(
                hintText: 'Search drivers...',
                prefixIcon: const Icon(Icons.search),
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(12.0),
                  borderSide: BorderSide.none,
                ),
                filled: true,
                fillColor: themeData.inputDecorationTheme.fillColor ??
                    themeData.cardColor.withOpacity(0.1),
              ),
            ),
            const SizedBox(height: 16.0),
            // Build the list of drivers using the async notifier state.
            asyncDrivers.when(
              data: (drivers) {
                if (drivers.isEmpty) {
                  return Padding(
                    padding: const EdgeInsets.all(8.0),
                    child: Text('No drivers found',
                        style: themeData.textTheme.bodyMedium),
                  );
                }
                return Flexible(
                  child: ListView.separated(
                    shrinkWrap: true,
                    itemCount: drivers.length,
                    separatorBuilder: (context, index) =>
                        Divider(color: themeData.dividerColor),
                    itemBuilder: (context, index) {
                      final driver = drivers[index];
                      return DriverListItem(
                          driver: driver, vehicleId: widget.vehicleId);
                    },
                  ),
                );
              },
              error: (error, stackTrace) {
                return Text('Error loading drivers',
                    style: themeData.textTheme.bodyMedium);
              },
              loading: () => const CircularProgressIndicator(),
            ),
            const SizedBox(height: 16.0),
          ],
        ),
      ),
    );
  }
}

class DriverListItem extends ConsumerWidget {
  final User driver;
  final int vehicleId;

  const DriverListItem({
    super.key,
    required this.driver,
    required this.vehicleId,
  });

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final themeData = Theme.of(context);
    final hireState = ref.watch(
      hireDriverProvider(
        HireDriverParams(driverId: driver.id, vehicleId: vehicleId),
      ),
    );

    return ListTile(
      key: ValueKey(driver.id),
      leading: CircleAvatar(
        backgroundImage: NetworkImage(driver.profilePictureUrl),
      ),
      title: Text('${driver.firstName} ${driver.lastName}'),
      trailing: FilledButton(
        onPressed: hireState is AsyncLoading
            ? null
            : () async {
                await ref
                    .read(hireDriverProvider(
                      HireDriverParams(
                          driverId: driver.id, vehicleId: vehicleId),
                    ).notifier)
                    .hireDriver(
                      HireDriverParams(
                          driverId: driver.id, vehicleId: vehicleId),
                    );

                hireState.maybeWhen(
                  error: (error, stackTrace) {
                    UiUtils.showOverlay(
                      context,
                      error.toString(),
                      themeData.colorScheme.error,
                    );
                  },
                  data: (_) {
                    UiUtils.showOverlay(
                      context,
                      'Driver hired',
                      AppColors.successColor,
                    );
                  },
                  orElse: () {},
                );
              },
        child: hireState is AsyncLoading
            ? SizedBox(
                width: 16,
                height: 16,
                child: CircularProgressIndicator(
                  strokeWidth: 2,
                ),
              )
            : const Text("Hire"),
      ),
    );
  }
}
