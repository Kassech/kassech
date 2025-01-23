import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:queue_manager_app/config/route/route.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: SelectRolePage(),
    );
  }
}

class SelectRolePage extends StatefulWidget {
  @override
  _SelectRolePageState createState() => _SelectRolePageState();
}

class _SelectRolePageState extends State<SelectRolePage> {
  String? selectedRole;
  final Map<String, dynamic> roles = {
    'Queue Manager': 3,
    'Owner': 4,
    'Driver': 2,
  };

  List<int> roleNumbers = [3, 4, 2];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        title: Text('Select Role'),
        leading: IconButton(
            onPressed: () {
              AppRouter.router.go('/signin');
            },
            icon: Icon(Icons.arrow_back)),
        backgroundColor: Colors.white,
      ),
      body: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 25.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              const Text(
                'Select a Role',
                textAlign: TextAlign.left,
                style: TextStyle(
                  fontSize: 40,
                  fontWeight: FontWeight.w800,
                ),
              ),
              const SizedBox(height: 20),
              Container(
                decoration: BoxDecoration(
                  border: Border.all(color: Colors.black),
                  borderRadius: BorderRadius.circular(5),
                ),
                child: DropdownButton<String>(
                  value: selectedRole,
                  hint: Text('Choose a role'),
                  items: roles.entries.map((entry) {
                    return DropdownMenuItem<String>(
                      value: entry.key,
                      child: Padding(
                        padding: const EdgeInsets.symmetric(
                            horizontal: 16.0,
                            vertical: 8.0), // Adjust padding as needed(8.0),
                        child: Text(
                          entry.key,
                          style: TextStyle(fontSize: 16),
                        ),
                      ),
                    );
                  }).toList(),
                  onChanged: (String? newRole) {
                    setState(() {
                      selectedRole = newRole;
                      print('Selected Role: ${roles[selectedRole]}');
                    });
                  },
                  style: TextStyle(color: Colors.black),
                  // focusColor: Colors.black,
                  icon: Icon(Icons.arrow_drop_down, color: Colors.black),
                  dropdownColor: Colors.white,
                  isExpanded: true,
                ),
              ),
              const SizedBox(height: 20),
              ElevatedButton(
                onPressed: () {
                  if (selectedRole != null) {
                    AppRouter.router.go('/signup', extra: roles[selectedRole]);
                  }
                },
                child: Text('Next'),
              ),
            ],
          )),
    );
  }
}
