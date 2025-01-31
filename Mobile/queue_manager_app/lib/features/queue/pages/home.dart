import 'package:flutter/material.dart';
import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:queue_manager_app/features/queue/pages/path_details_page.dart';
import 'package:queue_manager_app/features/queue/provider/passenger_provider.dart';
import 'package:queue_manager_app/features/queue/widgets/appDrawer.dart';
import 'package:queue_manager_app/features/queue/widgets/notification_modal.dart';

import '../../../core/permissions/app_permissions.dart';
import '../../../core/permissions/permission_wrapper.dart';
import '../../../shared/widgets/error_container.dart';
import '../models/path_model.dart';
import '../provider/path_provider.dart';
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
        drawer: AppDrawer(),
        body:  PathContainer(),
        );
  }
}

class QueueCard extends StatelessWidget {
  final PathModel path;

  const QueueCard({super.key, required this.path});

  @override
  Widget build(BuildContext context) {
    return Consumer(
      builder: (context, ref, child) {
        ref.watch(passengerNotifierProvider.notifier).getInitialData(path.id);
        final passengerCount = ref.watch(
          passengerNotifierProvider.select((state) => state[path.id.toString()] ?? 0),
        );
        return GestureDetector(
          onTap: () => context.push(PathDetailsPage.routeName),
          child: Card(
            child: Padding(
              padding: const EdgeInsets.all(16.0),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      SizedBox(
                        width: MediaQuery.sizeOf(context).width * 0.5,
                        child: Text(
                          path.pathName,
                          maxLines: 2,
                          overflow: TextOverflow.clip,
                          style: const TextStyle(fontSize: 18),
                        ),
                      ),
                      const SizedBox(height: 4),
                      Text(
                        'Route ID: ${path.id}',
                        style: const TextStyle(fontSize: 14),
                      ),
                      const SizedBox(height: 4),
                      Text(
                        'Passenger Count: $passengerCount',
                        style: const TextStyle(fontSize: 14),
                      ),
                    ],
                  ),
                  Row(
                    children: [
                      IconButton(
                        icon: const Icon(Icons.remove),
                        onPressed: () => ref
                            .read(passengerNotifierProvider.notifier)
                            .updateCount(path.id.toString(), -1),
                      ),
                      IconButton(
                        icon: const Icon(Icons.add),
                        onPressed: () => ref
                            .read(passengerNotifierProvider.notifier)
                            .updateCount(path.id.toString(), 1),
                      ),
                    ],
                  ),
                ],
              ),
            ),
          ),
        );
      },
    );
  }
}
