import 'package:flutter/material.dart';
import 'package:queue_manager_app/config/route/route.dart';

class HomeQueueManager extends StatefulWidget {
  @override
  _HomeQueueManagerState createState() => _HomeQueueManagerState();
}

class _HomeQueueManagerState extends State<HomeQueueManager> {
  final List<Map<String, dynamic>> queues = [
    {'routeName': 'Route 1', 'routeId': 'R001', 'queueCount': 5},
    {'routeName': 'Route 2', 'routeId': 'R002', 'queueCount': 3},
    {'routeName': 'Route 3', 'routeId': 'R003', 'queueCount': 10},
  ];
  bool _isDrawerOpen = false;

  void _toggleDrawer() {
    setState(() {
      _isDrawerOpen = !_isDrawerOpen;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Queue Manager'),
      ),
      drawer: Drawer(
        child: ListView(
          padding: EdgeInsets.zero,
          children: [
           const  DrawerHeader(
              decoration: BoxDecoration(
                color: Colors.black,
              ),
              child: Column(
                children: [
                  CircleAvatar(
                    radius: 50,
                    backgroundImage:AssetImage('assets/test.jpg') ,
                    )
                  
                  ,Text(
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
              leading: Icon(Icons.home),
              title: Text('Home'),
              onTap: () {
                // Handle item tap
                Navigator.pop(context); // Close the drawer
              },
            ),
            ListTile(
              leading: Icon(Icons.person),
              title: Text('Profile'),
              onTap: () {
                // Handle item tap
                AppRouter.router.go('/profile'); // Close the drawer
              },
            ),
            ListTile(
              leading: Icon(Icons.settings),
              title: Text('Settings'),
              onTap: () {
                // Handle item tap
                Navigator.pop(context); // Close the drawer
              },
            ),
            ListTile(
              leading:Icon(Icons.logout),
              title: Text('Logout'),
              onTap: () {
                // Handle item tap
                Navigator.pop(context); // Close the drawer
              },
            ),
          ],
        ),
      ),
      body: ListView.builder(
        itemCount: queues.length,
        itemBuilder: (context, index) {
          return QueueCard(
            routeName: queues[index]['routeName'],
            routeId: queues[index]['routeId'],
            initialCount: queues[index]['queueCount'],
          );
        },
      ),
    );
  }
}

class QueueCard extends StatefulWidget {
  final String routeName;
  final String routeId;
  final int initialCount;

  QueueCard({
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
                    style: TextStyle(color: Colors.white, fontSize: 18),
                  ),
                  SizedBox(height: 4),
                  Text(
                    'Route ID: ${widget.routeId}',
                    style: TextStyle(color: Colors.grey, fontSize: 14),
                  ),
                  SizedBox(height: 4),
                  Text(
                    'Queue: $queueCount',
                    style: TextStyle(color: Colors.orange, fontSize: 16),
                  ),
                ],
              ),

              // Increment and Decrement Buttons
              Row(
                children: [
                  IconButton(
                    icon: Icon(Icons.remove, color: Colors.red),
                    onPressed: decrementQueue,
                  ),
                  IconButton(
                    icon: Icon(Icons.add, color: Colors.green),
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
