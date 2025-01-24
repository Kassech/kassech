import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class ListOfCars extends StatelessWidget {
  final List<Map<String, String>> listOfCars = [
    {
      'make': 'Toyota',
      'plateNumber': 'A123456',
    },
    {
      'make': 'Honda',
      'plateNumber': 'B456321',
    },
    {
      'make': 'Ford',
      'plateNumber': 'C324513',
    },
    {
      'make': 'Chevrolet',
      'plateNumber': 'A334567',
    },
    {
      'make': 'Nissan',
      'plateNumber': 'B094563',
    },
  ];

  final int roleId;
  final bool isOwner;

  ListOfCars({super.key, required this.roleId, required this.isOwner});

  @override
  Widget build(BuildContext context) {
    if (roleId != 4 && !isOwner) {
      // Redirect to another page or show an appropriate message
      return Scaffold(
        appBar: AppBar(
          title: const Text('Access Denied'),
        ),
        body: Center(
          child: Text('You do not have permission to view this page.'),
        ),
      );
    }

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
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: ListView.builder(
          itemCount: listOfCars.length,
          itemBuilder: (context, index) {
            return GestureDetector(
              onTap: () {
                // Handle card tap
                context.go('/carlocation');
              },
              child: Card(
                margin: const EdgeInsets.only(bottom: 16.0),
                elevation: 4.0,
                color: Colors.black,
                child: ListTile(
                  title: Text(
                    listOfCars[index]['make']!,
                    style: const TextStyle(fontSize: 20.0, color: Colors.white),
                  ),
                  subtitle: Text(
                    listOfCars[index]['plateNumber']!,
                    style: TextStyle(fontSize: 14.0, color: Colors.grey[100]),
                  ),
                ),
              ),
            );
          },
        ),
      ),
    );
  }
}
