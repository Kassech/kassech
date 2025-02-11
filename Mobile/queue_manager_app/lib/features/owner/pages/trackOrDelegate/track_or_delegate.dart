import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:queue_manager_app/core/theme/app_colors.dart';
import 'package:queue_manager_app/features/owner/pages/delegate/delegation.dart';
import 'package:queue_manager_app/features/owner/pages/trackCar/track_car.dart';

class TrackOrDelegate extends StatelessWidget {
  const TrackOrDelegate({super.key});
  static const String routeName = '/trackOrDelegate';


  @override
  Widget build(BuildContext context) {
    
    return Scaffold(
      appBar: AppBar(
        backgroundColor: AppColors.white,
        elevation: 4.0,
        leading: IconButton(
          icon: const Icon(Icons.menu, color: AppColors.black),
          onPressed: () {
            Scaffold.of(context).openDrawer();
          },
        ),
        title: const Text(
          'Delegate/Track Car',
          style: TextStyle(
            color: AppColors.black,
            fontWeight: FontWeight.bold,
            fontSize: 24.0,
          ),
        ),
        actions: [
          IconButton(
            icon: const Icon(Icons.notifications, color: AppColors.black),
            onPressed: () {
              // Add search functionality
            },
          ),
        ],
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Row(
              children: [
                // First Card - Delegate Car
                Expanded(
                  child: GestureDetector(
                    onTap: () {
                      context.go(DelegationPage.routeName);
                    },
                    child: SizedBox(
                      height: 100,
                      child: Card(
                        color: AppColors.black,
                        child: Padding(
                          padding: const EdgeInsets.all(16.0),
                          child: Center(
                            child: Text(
                              'Delegate Car',
                              style: TextStyle(
                                color: AppColors.white,
                                fontSize: 18,
                                fontWeight: FontWeight.bold,
                              ),
                            ),
                          ),
                        ),
                      ),
                    ),
                  ),
                ),
                SizedBox(width: 16), // Spacing between the cards
                // Second Card - Track Car
                Expanded(
                  child: GestureDetector(
                    onTap: () {
                      context.go(TrackCar.routeName);
                      
                    },
                    child: SizedBox(
                      height: 100,
                      child: Card(
                        color: AppColors.black,
                        child: Padding(
                          padding: const EdgeInsets.all(16.0),
                          child: Center(
                            child: Text(
                              'Track Car',
                              style: TextStyle(
                                color: AppColors.white,
                                fontSize: 18,
                                fontWeight: FontWeight.bold,
                              ),
                            ),
                          ),
                        ),
                      ),
                    ),
                  ),
                ),
              ],
            ),
          ],
        ),
      ),
    );
  }
}
