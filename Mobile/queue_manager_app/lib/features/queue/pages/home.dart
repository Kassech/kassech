import 'package:flutter/material.dart';
import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:queue_manager_app/config/route/route.dart';
import 'package:queue_manager_app/features/queue/widgets/appDrawer.dart';
import 'package:queue_manager_app/features/queue/widgets/bottomNavBar.dart';
import 'package:queue_manager_app/features/queue/widgets/notification_modal.dart';

class HomeQueueManager extends StatefulWidget {
  const HomeQueueManager({super.key});

  @override
  _HomeQueueManagerState createState() => _HomeQueueManagerState();
}

class _HomeQueueManagerState extends State<HomeQueueManager> {
   final List<String> navTitles = ['Queue', 'Map', 'Profile'];
  final List<String> navRoutes = ['/home', '/home/qmdetails', '/profile'];

  final List<Map<String, dynamic>> queues = [
    {'routeName': 'Route 1', 'routeId': 'R001', 'queueCount': 5},
    {'routeName': 'Route 2', 'routeId': 'R002', 'queueCount': 3},
    {'routeName': 'Route 3', 'routeId': 'R003', 'queueCount': 10},
  ];
  bool _isDrawerOpen = false;
  int _selectedIndex = 0;

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

  void _toggleDrawer() {
    setState(() {
      _isDrawerOpen = !_isDrawerOpen;
    });
  }

  void _onItemTapped(int index) {
    setState(() {
      _selectedIndex = index;
    });
    switch (index) {
      case 0:
        AppRouter.router.go('/home');
        break;
      case 1:
        AppRouter.router.go('/map');
        break;
      case 2:
        AppRouter.router.go('/profile');
        break;
      default:
        AppRouter.router.go('/home');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Queue Manager'),
        actions: [
          IconButton(
            icon: const Icon(Icons.notifications),
            onPressed: () {
              AppRouter.router.go('/notifications');
            },
          ),
        ],
      ),
      drawer: AppDrawer(),
      body: Stack(
        children: [
          ListView.builder(
            itemCount: queues.length,
            itemBuilder: (context, index) {
              return QueueCard(
                routeName: queues[index]['routeName'],
                routeId: queues[index]['routeId'],
                initialCount: queues[index]['queueCount'],
              );
            },
          ),
          Positioned(
            bottom: 0,
            left: 0,
            right: 0,
            child: BottomNavBar(
              onItemTapped: _onItemTapped,
              selectedIndex: _selectedIndex,
              navTitles: navTitles,
              navRoutes: navRoutes,
            ),
          ),
        ],
      ),
    );
  }
}

class QueueCard extends StatefulWidget {
  final String routeName;
  final String routeId;
  final int initialCount;

  const QueueCard({
    super.key,
    required this.routeName,
    required this.routeId,
    required this.initialCount,
  });

  @override
  _QueueCardState createState() => _QueueCardState();
}

class _QueueCardState extends State<QueueCard> {
  late int queueCount;

  @override
  void initState() {
    super.initState();
    queueCount = widget.initialCount;
  }

  void incrementQueue() {
    setState(() {
      queueCount++;
    });
    // You can call a backend API here to update the queue count
  }

  void decrementQueue() {
    if (queueCount > 0) {
      setState(() {
        queueCount--;
      });
      // You can call a backend API here to update the queue count
    }
  }

  @override
  Widget build(BuildContext context) {
    return TextButton(
      onPressed: () {
        AppRouter.router.go('/home/qmdetails');
      },
      child: Card(
        color: Colors.black,
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
                  Text(
                    widget.routeName,
                    style: const TextStyle(color: Colors.white, fontSize: 18),
                  ),
                  const SizedBox(height: 4),
                  Text(
                    'Route ID: ${widget.routeId}',
                    style: const TextStyle(color: Colors.grey, fontSize: 14),
                  ),
                  const SizedBox(height: 4),
                  Text(
                    'Queue: $queueCount',
                    style: const TextStyle(color: Colors.orange, fontSize: 16),
                  ),
                ],
              ),

              // Increment and Decrement Buttons
              Row(
                children: [
                  IconButton(
                    icon: const Icon(Icons.remove, color: Colors.red),
                    onPressed: decrementQueue,
                  ),
                  IconButton(
                    icon: const Icon(Icons.add, color: Colors.green),
                    onPressed: incrementQueue,
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}
