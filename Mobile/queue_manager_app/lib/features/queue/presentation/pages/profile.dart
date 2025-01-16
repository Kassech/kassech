import 'dart:io';

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/config/route/route.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/authButton.dart';
import 'package:queue_manager_app/features/queue/presentation/provider/profileChangeProvider.dart';
import 'package:queue_manager_app/features/queue/presentation/widgets/profileLists.dart';
import 'package:image_picker/image_picker.dart';
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
          const SizedBox(
            height: 20,
          ),
          CircleAvatar(
            backgroundColor: Colors.black,
            radius: 70,
            child: Stack(
              children: [
                CircleAvatar(
                  radius: 100,
                  child: Image.asset(
                    'assets/dummyProfile.jpg',
                    width: 100,
                    height: 100,
                    fit: BoxFit.cover,
                  ),
                ),
                Positioned(
                  bottom: 0,
                  right: 0,
                  child: IconButton(
                    icon: CircleAvatar(backgroundColor: Colors.black, child: Icon(Icons.camera_alt, color: Colors.white,)),
                    onPressed: () {
                      // Add functionality to pick a new profile picture
                      final ref = ProviderScope.containerOf(context);
                      final ImagePicker _picker = ImagePicker();
                      _picker.pickImage(source: ImageSource.gallery).then((value) {
                        if (value != null) {
                          ref.read(profilePictureProvider)!.state = value.path;
                        }
                      });
                    },
                  ),
                ),
              ],
            ),),
           const SizedBox(
            height: 20,
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

extension on File {
  set state(String state) {}
}
