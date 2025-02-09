import 'package:flutter/material.dart';
import 'package:queue_manager_app/core/theme/app_colors.dart';

class TrackOrDelegate extends StatelessWidget {
  const TrackOrDelegate({super.key});
   static const String routeName = '/trackOrDelegate';


  @override
  Widget build(BuildContext context) {
    
    return 
       Scaffold(

        appBar: AppBar(
        backgroundColor: AppColors.white, // Use black
        elevation: 4.0,
        leading: IconButton(
          icon: const Icon(Icons.menu, color: AppColors.black), // Use white
          onPressed: () {
            Scaffold.of(context).openDrawer();
          },
        ),
        title: const Text(
          'Delegate/Track Car',
          style: TextStyle(
            color: AppColors.black, // Use white
            fontWeight: FontWeight.bold,
            fontSize: 24.0,
          ),
        ),
        actions: [
          IconButton(
            icon: const Icon(Icons.notifications,
                color: AppColors.black), // Use white
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
                  // First Card - Delegate Care
                  Expanded(
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
                  SizedBox(width: 16), // Spacing between the cards
                  // Second Card - Track Car
                  Expanded(
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
                ],
              ),
            ],
          ),
        ),
      );
    
  }
}
