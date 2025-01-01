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
            icon: const Icon(Icons.arrow_back)),
        title: const Text('Profile'),
      ),
      body: SafeArea(
          child: Center(
              child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          const CircleAvatar(
            radius: 50,
            backgroundImage: AssetImage('assets/test.jpg'),
          ),
          const SizedBox(
            height: 20,
          ),
          Container(
            child: Padding(
              padding: const EdgeInsets.symmetric(horizontal: 18.0),
              child: Column(
                children: [
                  const ProfileField(label: 'Name', value: 'John Doe'),
                  const ProfileField(label: 'Phone', value: '+251 123 456 789'),
                  const ProfileField(label: 'Email', value: 'jojo@example.com'),
                  const ProfileField(label: 'Vehicle Type', value: 'Sedan'),
                  const ProfileField(label: 'License Plate', value: 'ABC123'),
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
