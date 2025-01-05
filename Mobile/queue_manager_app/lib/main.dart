import 'package:flutter/material.dart';
import 'package:queue_manager_app/config/route/route.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:queue_manager_app/core/util/notification.dart';
import 'package:queue_manager_app/features/auth/domain/usecase/api_service.dart';
import 'package:queue_manager_app/features/auth/domain/usecase/auth_service.dart';
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
  final ApiService _apiService = ApiService();

  @override
  void initState() {
    super.initState();
    _initializeApp();
  }

  Future<void> _initializeApp() async {
    // final accessToken = await _authService.getAccessToken();
    // final refreshToken = await _authService.getRefreshToken();
    // print('Access Token: $accessToken');
    // print('Refresh Token: $refreshToken');
    // await _authService.saveTokens(accessToken!, refreshToken as String);

    // ignore: unnecessary_null_comparison
    // if (accessToken != null && refreshToken != null) {
    //   await _apiService.sendTokensToBackend(accessToken, refreshToken);
    // }
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      debugShowCheckedModeBanner: false,
      routerConfig: AppRouter.router,
    );
  }
}
