import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';

import '../models/path_model.dart';
import '../pages/path_details_page.dart';
import '../provider/passenger_provider.dart';

class PathCard extends StatelessWidget {
  const PathCard({super.key, required this.path});

  final PathModel path;

  @override
  Widget build(BuildContext context) {
    final themeData = Theme.of(context);
    return GestureDetector(
      onTap: () {
        context.pushNamed(PathDetailsPage.routeName, extra: path.id);
      },
      child: Card(
        child: Container(
          padding: const EdgeInsets.all(10),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            spacing: 10,
            children: [
              Text(path.pathName, style: themeData.textTheme.titleLarge),
              Row(
                spacing: 5,
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Flexible(
                    flex: 1,
                    child: Text(
                      path.route.startingLocation.locationName,
                      style: themeData.textTheme.bodyMedium?.copyWith(
                        fontSize: 13,
                      ),
                    ),
                  ),
                  Expanded(
                    flex: 2,
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.center,
                      children: [
                        Flexible(
                          child: Divider(
                            color: themeData.disabledColor,
                          ),
                        ),
                        Chip(
                          shape: RoundedRectangleBorder(
                            side: BorderSide(color: themeData.disabledColor),
                            borderRadius: BorderRadius.circular(100),
                          ),
                          label: Text(
                            path.estimatedTime,
                            style: themeData.textTheme.bodySmall?.copyWith(
                              color: themeData.disabledColor,
                            ),
                          ),
                        ),
                        Flexible(
                          child: Divider(
                            color: themeData.disabledColor,
                          ),
                        ),
                      ],
                    ),
                  ),
                  Flexible(
                    flex: 1,
                    child: Text(path.route.arrivalLocation.locationName),
                  ),
                ],
              ),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Consumer(
                    builder: (context, ref, child) {
                      ref
                          .watch(passengerNotifierProvider.notifier)
                          .getInitialData(path.id);
                      final passengerCount = ref.watch(
                        passengerNotifierProvider
                            .select((state) => state[path.id.toString()] ?? 0),
                      );
                      return Row(
                        children: [
                          IconButton(
                            icon: const Icon(Icons.remove),
                            onPressed: () => ref
                                .read(passengerNotifierProvider.notifier)
                                .updateCount(path.id.toString(), -1),
                          ),
                          Text(
                            '$passengerCount',
                            style: TextStyle(fontWeight: FontWeight.bold),
                          ),
                          IconButton(
                            icon: const Icon(Icons.add),
                            onPressed: () => ref
                                .read(passengerNotifierProvider.notifier)
                                .updateCount(path.id.toString(), 1),
                          ),
                        ],
                      );
                    },
                  ),
                  Chip(
                    label: Text(
                      path.isActive ? 'Active' : 'Inactive',
                      style: TextStyle(
                        color: path.isActive ? Colors.green : Colors.red,
                      ),
                    ),
                    shape: RoundedRectangleBorder(
                      side: BorderSide(
                        color: path.isActive ? Colors.green : Colors.red,
                      ),
                      borderRadius: BorderRadius.circular(10),
                    ),
                    backgroundColor: path.isActive ? Colors.green[100] : Colors.red[100],
                  ),
                ],
              )
            ],
          ),
        ),
      ),
    );
  }
}
