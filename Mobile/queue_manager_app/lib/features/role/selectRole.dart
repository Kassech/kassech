import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

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
      appBar: AppBar(
        title: Text('Select Role'),
        leading: IconButton(
            onPressed: () {
              context.go('/signin');
            },
            icon: Icon(Icons.arrow_back)),
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
              DropdownButtonFormField<String>(
                value: selectedRole,
                hint: Text('Choose a role'),
                items: roles.entries.map((entry) {
                  return DropdownMenuItem<String>(
                    value: entry.key,
                    child: Text(
                      entry.key,
                      // style: TextTheme.of(context).bodyMedium,
                    ),
                  );
                }).toList(),
                onChanged: (String? newRole) {
                  setState(() {
                    selectedRole = newRole;
                  });
                },
                isExpanded: true,
              ),
              const SizedBox(height: 20),
              ElevatedButton(
                onPressed: () {
                  if (selectedRole != null) {
                    context.go('/signup', extra: roles[selectedRole]);
                  }
                },
                child: Text('Next'),
              ),
            ],
          )),
    );
  }
}
