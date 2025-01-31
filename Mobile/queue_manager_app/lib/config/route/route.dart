import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';

import '../../features/auth/models/user.dart';
import '../../features/auth/pages/errorpage.dart';
import '../../features/auth/pages/selectRole.dart';
import '../../features/auth/pages/signinpage.dart';
import '../../features/auth/pages/signuppage.dart';
import '../../features/auth/pages/waitpage.dart';
import '../../features/auth/providers/auth_provider.dart';
import '../../features/owner/pages/carLocation/car_location.dart';
import '../../features/owner/pages/delegate/delegation.dart';
import '../../features/owner/pages/list/list_of_cars.dart';
import '../../features/queue/pages/home.dart';
import '../../features/queue/pages/noRoutesAssigned.dart';
import '../../features/queue/pages/notificaton_page.dart';
import '../../features/queue/pages/profile.dart';
import '../../features/queue/pages/path_details_page.dart';
import '../../features/splash/splash.dart';
import '../../shared/widgets/custom_navigation_bar.dart';

final rootNavigatorKey = GlobalKey<NavigatorState>();

GoRouter? _previousRouter;

final goRouterProvider = Provider<GoRouter>((ref) {
  final authState = ref.watch(authProvider);
  final refreshListenable = GoRouterRefreshNotifier(authState);

  final router = GoRouter(
    initialLocation: _previousRouter?.state?.fullPath ?? Splash.routeName,
    navigatorKey: rootNavigatorKey,
    refreshListenable: refreshListenable,
    debugLogDiagnostics: true,
    redirect: (context, state) {
      final user = authState.value;

      const publicRoutes = [
        Splash.routeName,
        SignInPage.routeName,
        SignUpPage.routeName,
        SelectRolePage.routeName,
      ];

      if (user == null) {
        if (publicRoutes.contains(state.matchedLocation)) {
          return null;
        }
        if (state.matchedLocation == SignInPage.routeName) {
          return null;
        }
        return SignInPage.routeName;
      } else if ((state.matchedLocation == SignInPage.routeName ||
          state.matchedLocation == SignUpPage.routeName)) {
        return HomePage.routeName;
      }
      return null;
    },
    errorBuilder: (context, state) => ErrorPage(state.error),
    routes: [
      StatefulShellRoute.indexedStack(
        builder: (context, state, navigationShell) {
          return Scaffold(
            extendBody: true,
            body: navigationShell,
            bottomNavigationBar: CustomNavigationBar(
              height: 60,
              selectedIndex: navigationShell.currentIndex,
              icons: const [
                Icons.home,
                Icons.person_2,
              ],
              onDestinationSelected: (index) => navigationShell.goBranch(index),
            ),
          );
        },
        branches: [
          StatefulShellBranch(
            routes: [
              GoRoute(
                path: HomePage.routeName,
                name: HomePage.routeName,
                builder: (context, state) => const HomePage(),
                routes: [
                  GoRoute(
                    path: PathDetailsPage.routeName,
                    name: PathDetailsPage.routeName,
                    parentNavigatorKey: rootNavigatorKey,
                    builder: (context, state) {
                      final extra = state.extra;
                      return PathDetailsPage(pathId: extra as int);
                    },
                  ),
                ]
              ),
            ],
          ),
          StatefulShellBranch(
            routes: [
              GoRoute(
                path: ProfilePage.routeName,
                name: ProfilePage.routeName,
                builder: (context, state) => const ProfilePage(),
              ),
            ],
          ),
        ],
      ),
      GoRoute(
        path: Splash.routeName,
        name: Splash.routeName,
        builder: (context, state) => const Splash(),
      ),
      GoRoute(
        path: SignInPage.routeName,
        name: SignInPage.routeName,
        builder: (context, state) => SignInPage(),
      ),
      GoRoute(
        path: SignUpPage.routeName,
        name: SignUpPage.routeName,
        builder: (context, state) {
          final roleId = state.extra as int?;
          return SignUpPage(role: roleId ?? 0);
        },
      ),
      GoRoute(
          path: ErrorPage.routeName,
          name: ErrorPage.routeName,
          builder: (context, state) => ErrorPage(state.error)),
      GoRoute(
        path: DelegationPage.routeName,
        name: DelegationPage.routeName,
        builder: (context, state) => DelegationPage(),
      ),
      GoRoute(
        path: '/noroutes',
        name: 'noroutes',
        builder: (context, state) => const NoRoutesAssignedYet(),
      ),
      GoRoute(
        path: SelectRolePage.routeName,
        name: SelectRolePage.routeName,
        builder: (context, state) => SelectRolePage(),
      ),
      GoRoute(
        path: WaitPage.routeName,
        name: WaitPage.routeName,
        builder: (context, state) => const WaitPage(),
      ),

//Owner Routes
      GoRoute(
        path: CarLocation.routeName,
        name: CarLocation.routeName,
        builder: (context, state) => CarLocation(),
      ),
      GoRoute(
        path: NotificationPage.routeName,
        name: NotificationPage.routeName,
        builder: (context, state) => NotificationPage(),
      ),

      GoRoute(
        path: ListOfCars.routeName,
        name: ListOfCars.routeName,
        builder: (context, state) {
          final extra = state.extra as Map<String, dynamic>;
          final roleId = extra['roleId'] as int;
          final isOwner = extra['isOwner'] as bool;
          return ListOfCars(roleId: roleId, isOwner: isOwner);
        },
      ),
    ],
  );

  _previousRouter = router;
  ref.onDispose(router.dispose);
  return router;
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
