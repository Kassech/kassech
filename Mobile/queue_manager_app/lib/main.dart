import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/config/provider/webSocket.dart';
import 'package:queue_manager_app/config/route/route.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:queue_manager_app/core/theme/app_theme.dart';
import 'package:queue_manager_app/features/notification/notification_service.dart';
import 'package:queue_manager_app/features/queue/domain/usecase/sendlocation.dart';
import 'firebase_options.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  
  await Firebase.initializeApp(
    options: DefaultFirebaseOptions.currentPlatform,
  );
  WidgetsFlutterBinding.ensureInitialized();

  // Initialize location tracking on app start.
  initializeLocation();
  runApp(ProviderScope(child: MyApp()));
}

class MyApp extends ConsumerStatefulWidget {
  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends ConsumerState<MyApp> {
  final NotificationService _notificationService = NotificationService();

 
  @override
  void initState() {
    super.initState();
    // Initialize the WebSocket connection
    ref.read(webSocketProvider);
    // Initialize location services
    initializeLocation();
    // Initialize notification service
    _notificationService.initialize();
  }


  @override  
  Widget build(BuildContext context) {
    return MaterialApp.router(
      debugShowCheckedModeBanner: false,
      theme: lightThemeData,
      darkTheme: darkThemeData,
      routerConfig: AppRouter.router,
    );
  }
}

final tokenProvider = StateNotifierProvider<TokenNotifier, String?>((ref) {
  return TokenNotifier();
});

class TokenNotifier extends StateNotifier<String?> {
  TokenNotifier() : super(null);

  void updateToken(String? token) {
    state = token;
  }
}
