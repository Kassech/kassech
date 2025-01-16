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
        child: Center(
          child: Container(
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
                  child: Text(role),
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
        ),
      ),
    );
  }
}

