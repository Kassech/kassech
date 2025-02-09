import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/core/theme/app_colors.dart';
import 'package:queue_manager_app/features/owner/models/car_model.dart';
import '../carLocation/car_location.dart';
import '../../providers/car_list_provider.dart';

class ListOfCars extends ConsumerWidget {
  final int roleId;
  final bool isOwner;

  ListOfCars({super.key, required this.roleId, required this.isOwner});

  static const String routeName = '/listOfCarsPage';

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    Future.delayed(Duration.zero, () async {
      await ref.read(carProvider.notifier).fetchCars();
    });

    return Consumer(builder: (context, ref, _) {
      final carListAsyncValue = ref.watch(carProvider);
return Scaffold(
        backgroundColor: Colors.white,
        appBar: AppBar(
          backgroundColor: Colors.white,
          elevation: 4.0,
          leading: const Icon(Icons.menu, color: Colors.black),
          title: const Text(
            'List of Cars',
            style: TextStyle(
              color: Colors.black,
              fontWeight: FontWeight.bold,
              fontSize: 24.0,
            ),
          ),
          actions: [
            IconButton(
              icon: const Icon(Icons.search, color: AppColors.black),
              onPressed: () {
                // Add search functionality
              },
            ),
          ],
        ),
        body: carListAsyncValue.when(
          data: (listOfCars) {
            if (listOfCars == null || listOfCars.isEmpty) {
              return const Center(
                child: Text(
                  'No cars found',
                  style: TextStyle(fontSize: 18.0, color: Colors.grey),
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
                      'https://source.unsplash.com/300x200/?car&random=$index'; // Use index for randomness

                  return GestureDetector(
                    onTap: () {
                      // Handle card tap
                      context.go(CarLocation.routeName);
                    },
                    child: Card(
                      margin: const EdgeInsets.only(bottom: 16.0),
                      elevation: 4.0,
                      shape: RoundedRectangleBorder(
                        borderRadius: BorderRadius.circular(12.0),
                      ),
                      child: Padding(
                        padding: const EdgeInsets.all(16.0),
                        child: Row(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            ClipRRect(
                              borderRadius: BorderRadius.circular(8.0),
                              child: Image.network(
                                imageUrl,
                                width: 80,
                                height: 80,
                                fit: BoxFit.cover,
                                errorBuilder: (context, error, stackTrace) {
                                  return Icon(
                                    Icons.car_repair, // Fallback icon
                                    size: 80,
                                    color: Colors.grey[400],
                                  );
                                },
                              ),
                            ),
                            const SizedBox(width: 16.0),
                            Expanded(
                              child: Column(
                                crossAxisAlignment: CrossAxisAlignment.start,
                                children: [
                                  Text(
                                    car.make,
                                    style: const TextStyle(
                                      fontSize: 20.0,
                                      fontWeight: FontWeight.bold,
                                      color: AppColors.darkOnPrimary,
                                    ),
                                  ),
                                  const SizedBox(height: 8.0),
                                  Text(
                                    'License: ${car.libre}',
                                    style: TextStyle(
                                      fontSize: 14.0,
                                      color: Colors.grey[700],
                                    ),
                                  ),
                                  Text(
                                    'Year: ${car.year}',
                                    style: TextStyle(
                                      fontSize: 14.0,
                                      color: Colors.grey[700],
                                    ),
                                  ),
                                  Text(
                                    'Color: ${car.color}',
                                    style: TextStyle(
                                      fontSize: 14.0,
                                      color: Colors.grey[700],
                                    ),
                                  ),
                                  Text(
                                    'Status: ${car.status}',
                                    style: TextStyle(
                                      fontSize: 14.0,
                                      color: Colors.grey[700],
                                    ),
                                  ),
                                ],
                              ),
                            ),
                          ],
                        ),
                      ),
                    ),
                  );
                },
              ),
            );
          },
          loading: () => Center(
            child: CircularProgressIndicator(
              valueColor: AlwaysStoppedAnimation<Color>(Colors.blue[800]!),
            ),
          ),
          error: (error, stackTrace) => Center(
            child: Text(
              'Error: $error',
              style: const TextStyle(fontSize: 18.0, color: Colors.red),
            ),
          ),
        ),
      );
    
    
    });
  }
}
