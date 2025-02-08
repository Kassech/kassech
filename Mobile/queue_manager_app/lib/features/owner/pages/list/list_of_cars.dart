import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
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
          leading: const Icon(Icons.menu),
          title: const Text(
            'List of Cars',
            style: TextStyle(color: Colors.black, fontWeight: FontWeight.bold),
          ),
        ),
        body: carListAsyncValue.when(
          data: (listOfCars) {
            if (listOfCars == null || listOfCars.isEmpty) {
              return const Center(
                child: Text('No cars found'),
              );
            }
            return Padding(
              padding: const EdgeInsets.all(20.0),
              child: ListView.builder(
                itemCount: listOfCars.length,
                itemBuilder: (context, index) {
                  final Car car = listOfCars[index];

                  return GestureDetector(
                    onTap: () {
                      // Handle card tap
                      context.go(CarLocation.routeName);
                    },
                    child: Card(
                      margin: const EdgeInsets.only(bottom: 16.0),
                      elevation: 4.0,
                      color: Colors.black,
                      child: ListTile(
                        leading: Image.network(
                          car.carPicture,
                          width: 50,
                          height: 50,
                          fit: BoxFit.cover,
                        ),
                        title: Text(
                          car.make,
                          style: const TextStyle(
                              fontSize: 20.0, color: Colors.white),
                        ),
                        subtitle: Column(
                          crossAxisAlignment: CrossAxisAlignment.start,
                          children: [
                            Text(
                              'License: ${car.libre}',
                              style: TextStyle(
                                  fontSize: 14.0, color: Colors.grey[100]),
                            ),
                            Text(
                              'Year: ${car.year}',
                              style: TextStyle(
                                  fontSize: 14.0, color: Colors.grey[100]),
                            ),
                            Text(
                              'Color: ${car.color}',
                              style: TextStyle(
                                  fontSize: 14.0, color: Colors.grey[100]),
                            ),
                            Text(
                              'Status: ${car.status}',
                              style: TextStyle(
                                  fontSize: 14.0, color: Colors.grey[100]),
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
          loading: () => Center(child: CircularProgressIndicator()),
          error: (error, stackTrace) => Center(child: Text('Error: $error')),
        ),
      );
    });
  }
}
