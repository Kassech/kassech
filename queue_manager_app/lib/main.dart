import 'package:flutter/material.dart';
import 'package:queue_manager_app/features/auth/presentation/pages/signinpage.dart';
import 'package:queue_manager_app/features/auth/presentation/pages/signuppage.dart';
import 'package:queue_manager_app/features/splash/splash.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      initialRoute: '/',
      home: const Splash(),
      routes: _buildRoutes(),
    );
  }

  Map<String, WidgetBuilder> _buildRoutes() {
    return {
      '/signin': (context) => SigninPage(),
      '/signup': (context) => SignUpPage(),
    };
  }
}
