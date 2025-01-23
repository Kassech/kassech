import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/config/route/route.dart';
import 'package:queue_manager_app/features/queue/presentation/widgets/profileLists.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/authButton.dart';

class ProfilePage extends ConsumerWidget {
  const ProfilePage({super.key});

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    AsyncValue<Map<String, String>> profileData = ref.watch(profileDataProvider);

    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        backgroundColor: Colors.white,
        leading: IconButton(
            onPressed: () {
              AppRouter.router.go('/home');
            },
            icon: const Icon(Icons.arrow_back)),
        title: const Text('Profile'),
      ),
      body: SafeArea(
        child: Center(
          child: profileData.when(
            data: (data) => Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                const SizedBox(height: 20),
                CircleAvatar(
                  backgroundColor: Colors.black,
                  radius: 70,
                  child: Image.network(
                    data['profilePictureUrl'] ?? 'https://via.placeholder.com/150',
                    width: 100,
                    height: 100,
                    fit: BoxFit.cover,
                  ),
                ),
                const SizedBox(height: 20),
                Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 18.0),
                  child: Column(
                    children: [
                      ProfileField(label: 'Name', value: data['name'] ?? 'N/A'),
                      ProfileField(label: 'Phone', value: data['phone'] ?? 'N/A'),
                      ProfileField(label: 'Email', value: data['email'] ?? 'N/A'),
                      ProfileField(label: 'Vehicle Type', value: data['vehicleType'] ?? 'N/A'),
                      ProfileField(label: 'License Plate', value: data['licensePlate'] ?? 'N/A'),
                      AuthButton(
                        label: 'Logout',
                        onPressed: () {
                          AppRouter.router.go('/signin');
                        },
                      ),
                    ],
                  ),
                ),
              ],
            ),
            loading: () => const CircularProgressIndicator(),
            error: (error, stack) => Text('Error: $error'),
          ),
        ),
      ),
    );
  }
}

final profileDataProvider = FutureProvider<Map<String, String>>((ref) async {
  // Replace with actual API call
  await Future.delayed(const Duration(seconds: 2));
  return {
    'profilePictureUrl': 'https://via.placeholder.com/150',
    'name': 'John Doe',
    'phone': '+251 123 456 789',
    'email': 'jojo@example.com',
    'vehicleType': 'Sedan',
    'licensePlate': 'ABC123',
  };
});

