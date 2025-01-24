import 'package:flutter/material.dart';
import 'package:queue_manager_app/features/queue/widgets/appDrawer.dart';

var someCondition = true;

class NoRoutesAssignedYet extends StatelessWidget {
  const NoRoutesAssignedYet({super.key});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 16.0, vertical: 8.0),
      child: Scaffold(
                backgroundColor: Colors.white,
      
        appBar: AppBar(
          backgroundColor: Colors.white,
          leading: IconButton(
            onPressed: () {
              AppDrawer();
            },
            icon: const Icon(Icons.menu),
          ),
          actions: [
            IconButton(
              onPressed: () {},
              icon: Icon(
                someCondition ? Icons.notifications : Icons.notifications_none,
              ),
            ),
          ],
        ),
        body: Center(
          child: Column(
            
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              const Padding(
                 padding:  EdgeInsets.symmetric(horizontal: 16.0),
                child:  Text(
                  'No Routes Assigned Yet',
                  textAlign: TextAlign.center,
                  style: TextStyle(
                    fontSize: 40,
                    color: Colors.black,
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ),
              const SizedBox(height: 20),
              Image.asset(
                'assets/noroutes.png',
                width: 300,
                height: 300,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
