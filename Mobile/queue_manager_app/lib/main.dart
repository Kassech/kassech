import 'package:flutter/material.dart';
import 'package:queue_manager_app/config/route/route.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:queue_manager_app/core/util/notification.dart';
import 'firebase_options.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await Firebase.initializeApp(
    options: DefaultFirebaseOptions.currentPlatform,
  );
  runApp(const MyApp());
}

class MyApp extends StatefulWidget {
  const MyApp({super.key});

  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  // ignore: unused_field
  final NotificationService _notificationService = NotificationService();
  // final AuthenticationService _authService = AuthenticationService();

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      debugShowCheckedModeBanner: false,
      routerConfig: AppRouter.router,
    );
  }
}
