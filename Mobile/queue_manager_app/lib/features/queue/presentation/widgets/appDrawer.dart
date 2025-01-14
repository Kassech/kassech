import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:queue_manager_app/config/route/route.dart';
import 'package:queue_manager_app/features/auth/domain/usecase/api_service.dart';
import 'package:queue_manager_app/core/util/token_storage.dart';

class AppDrawer extends ConsumerWidget {
  final ApiService apiService = ApiService();

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return Drawer(
      child: ListView(
        padding: EdgeInsets.zero,
        children: [
          const DrawerHeader(
            decoration: BoxDecoration(
              color: Colors.black,
            ),
            child: Column(
              children: [
                CircleAvatar(
                  radius: 50,
                  backgroundImage: AssetImage('assets/test.jpg'),
                ),
                Text(
                  'Queue Manager',
                  style: TextStyle(
                    color: Colors.white,
                    fontSize: 24,
                  ),
                ),
              ],
            ),
          ),
          ListTile(
            leading: const Icon(Icons.home),
            title: const Text('Home'),
            onTap: () {
              Navigator.pop(context); // Close the drawer
            },
          ),
          ListTile(
            leading: const Icon(Icons.person),
            title: const Text('Profile'),
            onTap: () {
              AppRouter.router.go('/profile'); // Close the drawer
            },
          ),
          ListTile(
            leading: const Icon(Icons.settings),
            title: const Text('Settings'),
            onTap: () {
              Navigator.pop(context); // Close the drawer
            },
          ),
          ListTile(
            leading: const Icon(Icons.logout),
            title: const Text('Logout'),
            onTap: () async {
              try {
                final isLoggedIn = await apiService.logoutApi();
                if (isLoggedIn) {
                  await storage.delete(key: 'accessToken');
                  AppRouter.router.go('/signin');
                } else {
                  print('Logout failed');
                  ScaffoldMessenger.of(context).showSnackBar(
                    SnackBar(
                        content: Text('Logout failed: Unauthorized request')),
                  );
                }
              } catch (e) {
                print('Logout failed: $e');
                context.pop(context);
                ScaffoldMessenger.of(context).showSnackBar(
                  SnackBar(content: Text('Logout failed: $e')),
                );
              }
            },
          ),
        ],
      ),
    );
  }
}
