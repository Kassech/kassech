// app_router.dart
import 'package:go_router/go_router.dart';
import 'package:queue_manager_app/features/Owner/carLocation/car_location.dart';
import 'package:queue_manager_app/features/Owner/delegate/delegation.dart';
import 'package:queue_manager_app/features/Owner/list/list_of_cars.dart';
import 'package:queue_manager_app/features/auth/domain/usecase/authentication_service.dart';
import 'package:queue_manager_app/features/auth/presentation/pages/errorpage.dart';
import 'package:queue_manager_app/features/auth/presentation/pages/forgotpassword.dart';
import 'package:queue_manager_app/features/auth/presentation/pages/signinpage.dart';
import 'package:queue_manager_app/features/auth/presentation/pages/signuppage.dart';
import 'package:queue_manager_app/features/auth/presentation/pages/waitpage.dart';
import 'package:queue_manager_app/features/queue/presentation/pages/home.dart';
import 'package:queue_manager_app/features/queue/presentation/pages/noRoutesAssigned.dart';
import 'package:queue_manager_app/features/queue/presentation/pages/notificaton_page.dart';
import 'package:queue_manager_app/features/queue/presentation/pages/profile.dart';
import 'package:queue_manager_app/features/queue/presentation/pages/qmdetails.dart';
import 'package:queue_manager_app/features/role/selectRole.dart';
import 'package:queue_manager_app/features/splash/splash.dart';

class AppRouter {
  static final GoRouter router = GoRouter(
    initialLocation: '/',
    redirect: (context, state) async {
      final isLoggedIn = await AuthenticationService().isAuthenticated();

      final protectedRoutes = [
        '/home',
        '/profile',
        '/home/qmdetails',
        '/noroutes'
      ];
      final isGoingToProtectedRoute = protectedRoutes.contains(state.path);
      state.path?.startsWith('/home');

      if (isGoingToProtectedRoute && !isLoggedIn) {
        // Redirect to sign-in if unauthenticated
        return '/signin';
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
}
