import 'package:flutter/material.dart';
import 'package:queue_manager_app/config/route/route.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/authButton.dart';
import 'package:queue_manager_app/features/queue/presentation/widgets/profileLists.dart';

class ProfilePage extends StatelessWidget {
  const ProfilePage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        leading: IconButton(
            onPressed: () {
              AppRouter.router.go('/home');
            },
            icon: Icon(Icons.arrow_back)),
        title: Text('Profile'),
      ),
      body: SafeArea(
          child: Center(
              child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          CircleAvatar(
            radius: 50,
            backgroundImage: AssetImage('assets/test.jpg'),
          ),
          SizedBox(
            height: 20,
          ),
          Container(
            child: Padding(
              padding: EdgeInsets.symmetric(horizontal: 18.0),
              child: Column(
                children: [
                  ProfileField(label: 'Name', value: 'John Doe'),
                  ProfileField(label: 'Phone', value: '+251 123 456 789'),
                  ProfileField(label: 'Email', value: 'jojo@example.com'),
                  ProfileField(label: 'Vehicle Type', value: 'Sedan'),
                  ProfileField(label: 'License Plate', value: 'ABC123'),
                  AuthButton(
                      label: 'Logout',
                      onPressed: () {
                        AppRouter.router.go('/signin');
                      })
                ],
              ),
            ),
          )
        ],
      ))),
    );
  }
}
