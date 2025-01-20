import 'package:flutter/material.dart';

class ListOfCars extends StatelessWidget {
  final List<Map<String, String>> listOfCars = [
    {'make': 'Toyota', 'plateNumber': 'A123456', },
    {'make': 'Honda', 'plateNumber': 'B456321', },
    {'make': 'Ford', 'plateNumber': 'C324513', },
    {'make': 'Chevrolet', 'plateNumber': 'A334567',},
    {'make': 'Nissan', 'plateNumber': 'B094563', },
  ];
  ListOfCars({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        backgroundColor: Colors.white,
        leading: Icon(Icons.menu),
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
            return Card(
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
                  style:TextStyle(fontSize: 14.0, color: Colors.grey[100]),
                ),
              ),
            );
          },
        ),
      ),
    );
  }
}


