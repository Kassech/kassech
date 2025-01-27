import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/config/provider/webSocket.dart';
import 'package:queue_manager_app/config/route/route.dart';
import 'package:firebase_core/firebase_core.dart';
import 'package:queue_manager_app/core/theme/app_theme.dart';
import 'package:queue_manager_app/features/notification/notification_service.dart';
import 'package:queue_manager_app/features/queue/domain/usecase/sendlocation.dart';
import 'core/services/api_service.dart';
import 'core/services/local_storage_service.dart';
import 'core/util/ui_utils.dart';
import 'firebase_options.dart';

Future<void> main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await Firebase.initializeApp(
    options: DefaultFirebaseOptions.currentPlatform,
  );

  await NotificationService().initialize();
  await LocalStorageService().init();
  ApiService();
  initializeLocation();
  runApp(ProviderScope(child: MyApp()));
}

class MyApp extends ConsumerWidget {
  const MyApp({super.key});


  @override  
  Widget build(BuildContext context, WidgetRef ref) {
    final goRouter = ref.watch(goRouterProvider);

    return MaterialApp.router(
      debugShowCheckedModeBanner: false,
      scaffoldMessengerKey: UiUtils.scaffoldMessengerKey,
      theme: lightThemeData,
      darkTheme: darkThemeData,
      routerDelegate: goRouter.routerDelegate,
      routeInformationProvider: goRouter.routeInformationProvider,
      routeInformationParser: goRouter.routeInformationParser,
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
