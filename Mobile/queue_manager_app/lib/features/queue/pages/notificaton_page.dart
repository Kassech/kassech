import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:go_router/go_router.dart';
import 'package:queue_manager_app/config/route/route.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:web_socket_channel/status.dart' as status;

class NotificationPage extends StatefulWidget {
  @override
  _NotificationPageState createState() => _NotificationPageState();
}

class _NotificationPageState extends State<NotificationPage> {
  final FirebaseMessaging _firebaseMessaging = FirebaseMessaging.instance;
  final WebSocketChannel _channel = WebSocketChannel.connect(
    Uri.parse('wss://your-websocket-url'),
  );

  @override
  void initState() {
    super.initState();
    _initializeFirebaseMessaging();
    _initializeWebSocket();
  }

  void _initializeFirebaseMessaging() {
    FirebaseMessaging.onMessage.listen((RemoteMessage message) {
      _showNotificationModal(
          message.notification?.title, message.notification?.body);
    });

    FirebaseMessaging.onMessageOpenedApp.listen((RemoteMessage message) {
      // Handle notification tapped logic here
    });
  }

  void _initializeWebSocket() {
    _channel.stream.listen((message) {
      _showNotificationModal('WebSocket Notification', message);
    });
  }

  void _showNotificationModal(String? title, String? body) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: Text(title ?? 'Notification'),
          content: Text(body ?? 'You have a new notification'),
          actions: [
            TextButton(
              onPressed: () {
                Navigator.of(context).pop();
              },
              child: Text('Close'),
            ),
          ],
        );
      },
    );
  }

  @override
  void dispose() {
    _channel.sink.close(status.goingAway);
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
          icon: Icon(Icons.arrow_back),
          onPressed: () {
            context.go('/home');
          },
        ),
        title: Text('Notifications'),
      ),
      body: Center(
        child: Text('Listening for notifications...'),
      ),
    );
  }
}
