import 'package:flutter/material.dart';
import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:queue_manager_app/features/queue/pages/qmdetails.dart';
import 'package:queue_manager_app/features/queue/widgets/appDrawer.dart';
import 'package:queue_manager_app/features/queue/widgets/notification_modal.dart';

import '../../../core/permissions/app_permissions.dart';
import '../../../core/permissions/permission_wrapper.dart';
import '../models/route_model.dart';
import '../provider/passangers_provider.dart';
import '../provider/routes_provider.dart';
import 'notificaton_page.dart';

class HomeQueueManager extends StatefulWidget {
  const HomeQueueManager({super.key});

  static const String routeName = '/homePage';

  @override
  State<HomeQueueManager> createState() => _HomeQueueManagerState();
}

class _HomeQueueManagerState extends State<HomeQueueManager> {

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
                context.go(NotificationPage.routeName);
              },
            ),
          ],
        ),
        drawer: AppDrawer(),
        body: PermissionWrapper(
          requiredPermission: AppPermissions.viewAssignedRoutes,
          fallback: const Center(
            child: Text('You do not have permission to view this page'),
          ),
          child: Consumer(
            builder: (context, ref, child) {
              final routes = ref.watch(routesProvider);
              return routes.when(
                data: (route) {
                  if (route == null) {
                    return const Center(
                      child: Text('No routes found'),
                    );
                  }
                  return ListView.builder(
                    itemCount: route.length,
                    itemBuilder: (context, index) {
                      return QueueCard(
                        route: route[index],
                      );
                    },
                  );
                },
                loading: () => const Center(child: CircularProgressIndicator()),
                error: (error, stackTrace) => Center(
                  child: Text('Error: $error'),
                ),
              );
            },
          ),
        )
        // body: Stack(
        //   children: [
        //     ListView.builder(
        //       itemCount: queues.length,
        //       itemBuilder: (context, index) {
        //         return QueueCard(
        //           routeName: queues[index]['routeName'],
        //           routeId: queues[index]['routeId'],
        //           initialCount: queues[index]['queueCount'],
        //         );
        //       },
        //     ),
        //     Positioned(
        //       bottom: 0,
        //       left: 0,
        //       right: 0,
        //       child: BottomNavBar(
        //         onItemTapped: _onItemTapped,
        //         selectedIndex: _selectedIndex,
        //         navTitles: navTitles,
        //         navRoutes: navRoutes,
        //       ),
        //     ),
        //   ],
        // ),
        );
  }
}

class QueueCard extends StatelessWidget {
  final RouteModel route;

  const QueueCard({
    super.key,
    required this.route,
  });

  @override
  Widget build(BuildContext context) {
    return Consumer(
      builder: (context, ref, child) {
        final passengerQueue = ref.watch(passengerStreamProvider(route.id));
        return GestureDetector(
          onTap: () {
            context.go(QueueManagerDetails.routeName);
          },
          child: Card(
            // margin: EdgeInsets.symmetric(horizontal: 16, vertical: 8),
            child: Padding(
              padding: const EdgeInsets.all(16.0),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  // Route Information
                  Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      SizedBox(
                        width: MediaQuery.sizeOf(context).width * 0.5,
                        child: Text(
                          route.name,
                          maxLines: 2,
                          overflow: TextOverflow.clip,
                          style: const TextStyle(fontSize: 18),
                        ),
                      ),
                      const SizedBox(height: 4),
                      Text(
                        'Route ID: ${route.id}',
                        style: const TextStyle(fontSize: 14),
                      ),
                      const SizedBox(height: 4),
                      StreamBuilder(
                        stream: ref
                            .watch(passengerStreamProvider(route.id).notifier)
                            .passengerCountStream,
                        builder: (context, snapshot) {
                          if (snapshot.hasData) {
                            return Text(
                              'Passenger Count: ${snapshot.data}',
                              style: const TextStyle(fontSize: 14),
                            );
                          }
                          return const Text(
                            'Loading...',
                            style: TextStyle(fontSize: 14),
                          );
                        },
                      ),
                      passengerQueue.when(
                        skipLoadingOnReload: true,
                        skipLoadingOnRefresh: true,
                        data: (data) {
                          return Text(
                            'Passenger Count: $data',
                            style: const TextStyle(fontSize: 14),
                          );
                        },
                        loading: () => Text(
                          'Loading...',
                          style: const TextStyle(fontSize: 14),
                        ),
                        error: (error, stackTrace) {
                          return Text(
                            'Error: $error',
                            style: const TextStyle(fontSize: 14),
                          );
                        },
                      ),
                    ],
                  ),

                  // Increment and Decrement Buttons
                  Row(
                    children: [
                      IconButton(
                        icon: const Icon(Icons.remove),
                        onPressed: () {
                          ref
                              .read(passengerStreamProvider(route.id).notifier)
                              .decrementPassengerCount(route.id);
                        },
                      ),
                      IconButton(
                        icon: const Icon(Icons.add),
                        onPressed: () {
                          ref
                              .read(passengerStreamProvider(route.id).notifier)
                              .incrementPassengerCount(route.id);
                        },
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
