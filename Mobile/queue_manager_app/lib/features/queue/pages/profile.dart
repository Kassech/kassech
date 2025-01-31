import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/features/auth/providers/auth_provider.dart';
import 'package:queue_manager_app/features/queue/widgets/profileLists.dart';

class ProfilePage extends ConsumerWidget {
  const ProfilePage({super.key});

  static const String routeName = '/profilePage';

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final user = ref.watch(authProvider).value;

    return Scaffold(
      extendBody: true,
      appBar: AppBar(
        title: const Text('Profile'),
      ),
      body: Column(
        children: [
          const SizedBox(height: 20),
          CircleAvatar(
            radius: 70,
            child: Image.network(
              user?.profilePictureUrl ?? 'https://via.placeholder.com/150',
              width: 100,
              height: 100,
              fit: BoxFit.cover,
              errorBuilder: (context, error, stackTrace) {
                return const Icon(Icons.person, size: 100);
              },
            ),
          ),
          const SizedBox(height: 20),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 18.0),
            child: Column(
              children: [
                ProfileField(label: 'Name', value: '${user?.firstName} ${user?.lastName}'),
                ProfileField(label: 'Phone', value: user?.phoneNumber ?? 'N/A'),
                ProfileField(label: 'Email', value: user?.email ?? 'N/A'),
                // ProfileField(label: 'Vehicle Type', value: user?.firstName ?? 'N/A'),
                // ProfileField(label: 'License Plate', value: user?.firstName ?? 'N/A'),
                ElevatedButton(
                  onPressed: () {
                    ref.read(authProvider.notifier).logout();
                  },
                  child: const Text('Logout'),
                )
              ],
            ),
          ),
        ],
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

