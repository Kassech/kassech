import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:queue_manager_app/features/owner/pages/carLocation/car_location.dart';
import 'package:queue_manager_app/features/owner/pages/delegate/delegation.dart';
import 'package:queue_manager_app/features/owner/pages/list/list_of_cars.dart';
import 'package:queue_manager_app/features/queue/pages/home.dart';
import 'package:queue_manager_app/features/queue/pages/noRoutesAssigned.dart';
import 'package:queue_manager_app/features/queue/pages/notificaton_page.dart';
import 'package:queue_manager_app/features/queue/pages/profile.dart';
import 'package:queue_manager_app/features/queue/pages/qmdetails.dart';
import 'package:queue_manager_app/features/role/selectRole.dart';
import 'package:queue_manager_app/features/splash/splash.dart';

import '../../features/auth/models/user.dart';
import '../../features/auth/pages/errorpage.dart';
import '../../features/auth/pages/forgotpassword.dart';
import '../../features/auth/pages/signinpage.dart';
import '../../features/auth/pages/signuppage.dart';
import '../../features/auth/pages/waitpage.dart';
import '../../features/auth/providers/auth_provider.dart';

final goRouterProvider = Provider<GoRouter>((ref) {
  final authState = ref.watch(authProvider);
  final notifier = GoRouterRefreshNotifier(authState);

  return GoRouter(
    initialLocation: '/',
    debugLogDiagnostics: true,
    refreshListenable: notifier,
    redirect: (context, state) async {
      final user = authState.value;

      if (user == null && state.uri.toString() != '/login') {
        return '/login';
      }

      if (user != null && state.uri.toString() == '/login') {
        return '/home';
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
          builder: (context, state) {
            final roleId = state.extra as int;
            return SignUpPage(role: roleId);
          }),
      GoRoute(
        path: '/forgotpassword',
        builder: (context, state) => ForgotPassword(),
      ),
      GoRoute(
        path: '/home',
        name: 'home',
        builder: (context, state) => const HomeQueueManager(),
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
      GoRoute(
          path: '/delegation', builder: (context, state) => DelegationPage()),
      GoRoute(
          path: '/noroutes',
          name: 'noroutes',
          builder: (context, state) => const NoRoutesAssignedYet()),
      GoRoute(
          path: '/home/qmdetails',
          name: 'qmdetails',
          builder: (context, state) => const QueueManagerDetalils()),
      GoRoute(path: '*', builder: (context, state) => const ErrorPage()),
      GoRoute(
          path: '/profile', builder: (context, state) => const ProfilePage()),
      GoRoute(
          path: '/selectRole', builder: (context, state) => SelectRolePage()),
      GoRoute(path: '/wait', builder: (context, state) => const WaitPage()),

      //Owner Routes
      // GoRoute(path: '/cars', builder: (context, state) => ListOfCars(),),
      GoRoute(path: '/carlocation', builder: (context, state) => CarLocation(),),
      GoRoute(path: '/notifications', builder: (context, state) => NotificationPage(),),

      GoRoute(
        path: '/listofcars',
        builder: (context, state) {
          final extra = state.extra as Map<String, dynamic>;
          final roleId = extra['roleId'] as int;
          final isOwner = extra['isOwner'] as bool;
          return ListOfCars(roleId: roleId, isOwner: isOwner);
        },
      ),
      GoRoute(path: '/carlocation', builder: (context, state) => CarLocation()),
      GoRoute(
        path: '/notifications',
        builder: (context, state) => NotificationPage(),
      )
    ],
  );
});

class GoRouterRefreshNotifier extends ChangeNotifier {
  GoRouterRefreshNotifier(AsyncValue<User?> authState) {
    authState.when(
      data: (user) => notifyListeners(),
      error: (_, __) => notifyListeners(),
      loading: () => notifyListeners(),
    );
  }
}



