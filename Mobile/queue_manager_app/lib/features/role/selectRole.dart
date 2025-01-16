import 'package:flutter/material.dart';
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
  final List<String> roles = ['Queue Manager', 'Owner', 'Driver'];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(title: Text('Select Role'),  leading: IconButton(
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
                items: roles.map((String role) {
                  return DropdownMenuItem<String>(
                    
                    value: role,
                    child: Padding(
                      padding: const EdgeInsets.symmetric(horizontal: 16.0, vertical: 8.0), // Adjust padding as needed(8.0),
                      child: Text(role, style: TextStyle(fontSize: 16),),
                    ),
                  );
                }).toList(),
                onChanged: (String? newRole) {
                  setState(() {
                    selectedRole = newRole;
                  });
                },
                style: TextStyle(color: Colors.black),
                focusColor: Colors.black,
                icon: Icon(Icons.arrow_drop_down, color: Colors.black),
                dropdownColor: Colors.white,
                isExpanded: true,
              ),
            ),
          ],
        ),
      ),
    );
  }
}

