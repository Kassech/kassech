// app_router.dart
import 'package:driver_app/features/auth/domain/usecase/auth_service.dart';
import 'package:driver_app/features/auth/presentation/pages/errorpage.dart';
import 'package:driver_app/features/auth/presentation/pages/forgotpassword.dart';
import 'package:driver_app/features/auth/presentation/pages/signinpage.dart';
import 'package:driver_app/features/auth/presentation/pages/signuppage.dart';
import 'package:driver_app/features/driverhome/presentation/pages/frontpagemap.dart';
import 'package:driver_app/features/splash/splash.dart';
import 'package:go_router/go_router.dart';

class AppRouter {
  static final GoRouter router = GoRouter(
    initialLocation: '/',
    redirect: (context, state) async {
      // Check if the user is authenticated
      final isLoggedIn = await checkUserAuthentication();

      // Protect /home and any authenticated routes
      final isGoingToProtectedRoute = state.path?.startsWith('/home');
      if (isGoingToProtectedRoute != null && !isLoggedIn) {
        return '/signin'; // Redirect to sign-in if unauthenticated
      }

      return null;
    },
    errorBuilder: (context, state) => const ErrorPage(),
    routes: [
      GoRoute(
        path: '/',
        name: 'splash',
        builder: (context, state) => const Splash(),
      ),
      GoRoute(
        path: '/signin',
        name: 'signin',
        builder: (context, state) => SigninPage(),
      ),
      GoRoute(
        path: '/signup',
        name: 'signup',
        builder: (context, state) => SignUpPage(),
      ),
      GoRoute(
        path: '/forgotpassword',
        builder: (context, state) => ForgotPassword(),
      ),
       GoRoute(
        path: '/home',
        name: 'home',
        builder: (context, state) => DriverHome(),
      ),
      GoRoute(
        path: '/404',
        name: '404',
        builder: (context, state) => const ErrorPage(),
      ),
      GoRoute(
        path: '/500',
        name: '500',
        builder: (context, state) => const ErrorPage(),
      ),
      
      GoRoute(path: '*', builder: (context, state) => const ErrorPage()),
      
    ],
  );
}
