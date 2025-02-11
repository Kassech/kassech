import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/core/theme/app_colors.dart';
import 'package:queue_manager_app/features/owner/models/car_model.dart';
import 'package:queue_manager_app/shared/widgets/error_container.dart';
import 'package:queue_manager_app/features/owner/pages/trackOrDelegate/track_or_delegate.dart';
import 'package:queue_manager_app/features/queue/widgets/appDrawer.dart';
import '../../providers/car_list_provider.dart';

class ListOfCars extends ConsumerWidget {
  final int roleId;
  final bool isOwner;

  const ListOfCars({super.key, required this.roleId, required this.isOwner});

  static const String routeName = '/listOfCarsPage';

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final themeData = Theme.of(context);
    Future.delayed(Duration.zero, () async {
      await ref.read(carProvider.notifier).fetchCars();
    });

    return Consumer(builder: (context, ref, _) {
      final carListAsyncValue = ref.watch(carProvider);
      return Scaffold(
        backgroundColor: AppColors.white, // Use darkScaffoldBackground
        appBar: AppBar(
          elevation: 4.0,
            leading: IconButton(
            icon: const Icon(Icons.menu, color: AppColors.black), // Use white
            onPressed: () {
              Scaffold.of(context).openDrawer();
            },
            ),
          title: const Text(
            'List of Cars',
            style: TextStyle(// Use white
              fontWeight: FontWeight.bold,
              fontSize: 24.0,
            ),
          ),
          actions: [
            IconButton(
              icon: const Icon(Icons.notifications),
              // Use white
              onPressed: () {
                // Add search functionality
              },
            ),
          ],
        ),
        drawer: const AppDrawer(),
        body: carListAsyncValue.when(
          data: (listOfCars) {
            if (listOfCars == null || listOfCars.isEmpty) {
              return const Center(
                child: Text(
                  'No cars found',
                  style: TextStyle(
                      fontSize: 18.0,
                  ), // Use lightGray
                ),
              );
            }
            return Padding(
              padding: const EdgeInsets.all(16.0),
              child: ListView.builder(
                itemCount: listOfCars.length,
                itemBuilder: (context, index) {
                  final Car car = listOfCars[index];
                  final imageUrl =
                      'https://source.unsplash.com/300x200/?car&random=$index';

                  // Status-based color
                  final statusColor = car.status == 'active'
                      ? AppColors.successColor
                      : car.status == 'maintenance'
                          ? AppColors.warningColor
                          : AppColors.errorColor;

                  return GestureDetector(
                    onTap: () {
                      context.push(CarLocation.routeName, extra: car);
                    },
                    child: Container(
                      decoration: BoxDecoration(
                        border: Border(
                          bottom: BorderSide(
                            color: themeData.disabledColor,
                            width: 1.0,
                          ),
                        ),
                      ),
                      padding: const EdgeInsets.all(16.0),
                      child: Row(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          // Car Image
                          ClipRRect(
                            borderRadius: BorderRadius.circular(8.0),
                            child: Image.network(
                              imageUrl,
                              width: 100,
                              height: 100,
                              fit: BoxFit.cover,
                              errorBuilder: (context, error, stackTrace) {
                                return Container(
                                  width: 100,
                                  height: 100,
                                  color: themeData.brightness == Brightness.dark
                                      ? AppColors.darkGray
                                      : AppColors.lightGray,
                                  child: const Icon(
                                    Icons.car_repair,
                                    size: 50,
                                  ),
                                );
                              },
                            ),
                          ),
                          const SizedBox(width: 16.0),
                          // Car Details
                          Expanded(
                            child: Column(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                // Car Make
                                Text(
                                  car.make,
                                  style: const TextStyle(
                                    fontSize: 22.0,
                                    fontWeight: FontWeight.bold,
                                  ),
                                ),
                                const SizedBox(height: 8.0),
                                // License Plate
                                Text(
                                  'License: ${car.licenseNumber}',
                                  style: TextStyle(
                                    fontSize: 16.0,
                                  ),
                                ),
                                const SizedBox(height: 4.0),
                                // Year
                                Text(
                                  'Year: ${car.year}',
                                  style: TextStyle(
                                    fontSize: 16.0,
                                  ),
                                ),
                                const SizedBox(height: 4.0),
                                // Color
                                Row(
                                  children: [
                                    Text(
                                      'Color: ',
                                      style: TextStyle(
                                        fontSize: 16.0,
                                      ),
                                    ),
                                    Container(
                                      width: 16,
                                      height: 16,
                                      decoration: BoxDecoration(
                                        color:
                                            AppColors.getCarColor(car.color),
                                        // Use car color from AppColors
                                        shape: BoxShape.circle,
                                      ),
                                    ),
                                  ],
                                ),
                                const SizedBox(height: 4.0),
                                // Status
                                Row(
                                  children: [
                                    Text(
                                      'Status: ',
                                      style: TextStyle(
                                        fontSize: 16.0,
                                        color: AppColors
                                            .darkGray, // Use lightGray
                                      ),
                                    ),
                                    Container(
                                      padding: const EdgeInsets.symmetric(
                                        horizontal: 8.0,
                                        vertical: 4.0,
                                      ),
                                      decoration: BoxDecoration(
                                        color: statusColor.withOpacity(0.2),
                                        borderRadius:
                                            BorderRadius.circular(12.0),
                                      ),
                                      child: Text(
                                        car.status,
                                        style: TextStyle(
                                          fontSize: 14.0,
                                          fontWeight: FontWeight.bold,
                                          color: statusColor,
                                        ),
                                      ),
                                    ),
                                  ],
                                ),
                              ],
                            ),
                          ),
                        ],
                      ),
                    ),
                  );
                },
              ),
            );
          },
          loading: () => Center(
            child: CircularProgressIndicator(
              valueColor:
                  AlwaysStoppedAnimation<Color>(AppColors.blue), // Use blue
            ),
          ),
          error: (error, stackTrace) => ErrorContainer(
            errorMessageText: error.toString(),
            onTapRetry: () {
              ref.read(carProvider.notifier).fetchCars();
            },
          ),
        ),
      );
    });
  }
}
