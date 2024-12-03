import 'package:go_router/go_router.dart';
import 'package:queue_manager_app/features/auth/presentation/pages/forgotpassword.dart';
import 'package:queue_manager_app/features/auth/presentation/pages/signinpage.dart';
import 'package:queue_manager_app/features/auth/presentation/pages/signuppage.dart';
import 'package:queue_manager_app/features/queue/presentation/pages/home.dart';
import 'package:queue_manager_app/features/splash/splash.dart';

class AppRouter {
  static final GoRouter router = GoRouter(
    initialLocation: '/',
    routes: [
      GoRoute(path: '/', builder: (context, state) => const Splash()),
      GoRoute(path: '/signin', builder: (context, state) => SigninPage()),
      GoRoute(path: '/signup', builder: (context, state) => SignUpPage()),
      GoRoute(
          path: '/home', builder: (context, state) => const HomeQueueManager()),
      GoRoute(
          path: '/forgotpassword',
          builder: (context, state) => ForgotPassword())
    ],
  );
}
