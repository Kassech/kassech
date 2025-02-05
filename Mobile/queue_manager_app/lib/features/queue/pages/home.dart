import 'package:flutter/material.dart';
import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:go_router/go_router.dart';

import '../../../config/route/route.dart';
import '../widgets/notification_modal.dart';
import '../widgets/path_container.dart';
import 'notificaton_page.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  static const String routeName = '/homePage';

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {

  @override
  void initState() {
    super.initState();
    _initializeFirebaseMessaging();
  }

  void _initializeFirebaseMessaging() {
    FirebaseMessaging.onMessage.listen((RemoteMessage message) {
      _showNotificationModal(message.notification?.title ?? 'Notification',
          message.notification?.body ?? 'You have a new notification');
    });

    FirebaseMessaging.onMessageOpenedApp.listen((RemoteMessage message) {
      // Handle notification tapped logic here
    });
  }

  void _showNotificationModal(String title, String body) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return NotificationModal(title: title, body: body);
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: Builder(
          builder: (context) {
            return IconButton(
              icon: const Icon(Icons.menu),
              onPressed: () {
                scaffoldKey.currentState!.openDrawer();
              },
            );
          },
        ),
        title: const Text('Routes'),
        actions: [
          IconButton(
            icon: const Icon(Icons.notifications),
            onPressed: () {
              context.push(NotificationPage.routeName);
            },
          ),
        ],
      ),
      body: PathContainer(),
    );
  }
}
