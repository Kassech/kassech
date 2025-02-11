import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:queue_manager_app/features/owner/pages/trackCar/track_car.dart';
import 'package:queue_manager_app/features/owner/pages/trackOrDelegate/track_or_delegate.dart';
import '../../features/auth/models/user.dart';
import '../../features/auth/pages/errorpage.dart';
import '../../features/auth/pages/selectRole.dart';
import '../../features/auth/pages/signinpage.dart';
import '../../features/auth/pages/signuppage.dart';
import '../../features/auth/pages/waitpage.dart';
import '../../features/auth/providers/auth_provider.dart';
import '../../features/owner/models/car_model.dart';
import '../../features/owner/pages/carLocation/car_location.dart';
import '../../features/owner/pages/delegate/delegation.dart';
import '../../features/owner/pages/list/list_of_cars.dart';
import '../../features/queue/models/path_model.dart';
import '../../features/queue/pages/home.dart';
import '../../features/queue/pages/noRoutesAssigned.dart';
import '../../features/queue/pages/notificaton_page.dart';
import '../../features/queue/pages/profile.dart';
import '../../features/location/screens/path_details_page.dart';
import '../../features/queue/widgets/appDrawer.dart';
import '../../features/splash/splash.dart';
import '../../shared/widgets/custom_navigation_bar.dart';

final rootNavigatorKey = GlobalKey<NavigatorState>();
final GlobalKey<ScaffoldState> scaffoldKey = GlobalKey<ScaffoldState>();

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
        if (user.roles.contains('QueueManager') ||
            user.roles.contains('Driver')) {
          return HomePage.routeName;
        } else if (user.roles.contains('Owner')) {
          return ListOfCars.routeName;
        } else {
          return SignInPage.routeName;
        }
      }
      return null;
    },
    // errorBuilder: (context, state) => ErrorPage(error: state.error.toString()),
    routes: [
      StatefulShellRoute.indexedStack(
        builder: (context, state, navigationShell) {
          return Scaffold(
            key: scaffoldKey,
            extendBody: true,
            body: navigationShell,
            drawer: AppDrawer(),
            bottomNavigationBar: CustomNavigationBar(
              height: 60,
              selectedIndex: navigationShell.currentIndex,
              icons: [
                if (authState.value!.roles.contains('QueueManager') ||
                    authState.value!.roles.contains('Driver'))
                  Icons.home,
                if (authState.value!.roles.contains('Owner')) Icons.car_rental,
                Icons.person_2,
              ],
              onDestinationSelected: (index) => navigationShell.goBranch(index),
            ),
          );
        },
        branches: [
          if (authState.value != null)
            if ((authState.value!.roles.contains('QueueManager') ||
                authState.value!.roles.contains('Driver')))
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
                          final extra = state.extra as Map<String, dynamic>?;
                          final int pathId = extra?['pathId'];
                          final PathModel path = extra?['path'];
                          return PathDetailsPage(pathId: pathId, path: path);
                        },
                      ),
                    ],
                  ),
                ],
              ),
          if (authState.value != null &&
              authState.value!.roles.contains('Owner'))
            StatefulShellBranch(
              routes: [
                GoRoute(
                  path: ListOfCars.routeName,
                  name: ListOfCars.routeName,
                  builder: (context, state) {
                    print('state: ${state}');
                    final extra = state.extra as Map<String, dynamic>?;
                    final roleId = extra?['roleId'] as int? ?? 0;
                    final isOwner = extra?['isOwner'] as bool? ?? false;
                    return ListOfCars(roleId: roleId, isOwner: isOwner);
                  },
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
          builder: (context, state) => ErrorPage(
                state.error,
                error: state.error?.toString() ?? 'Unknown error',
              )),
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
        builder: (context, state) {
          return CarLocation(car: state.extra as Car,);
        },
      ),
      GoRoute(
        path: NotificationPage.routeName,
        name: NotificationPage.routeName,
        builder: (context, state) => NotificationPage(),
      ),
      GoRoute(
          path: TrackOrDelegate.routeName,
          name: TrackOrDelegate.routeName,
          builder: (context, state) => TrackOrDelegate()),

      GoRoute(
          path: TrackCar.routeName,
          name: TrackCar.routeName,
          builder: (context, state) => TrackCar()),

      // GoRoute(
      //   path: ListOfCars.routeName,
      //   name: ListOfCars.routeName,
      //   builder: (context, state) {
      //     final extra = state.extra as Map<String, dynamic>;
      //     final roleId = extra['roleId'];
      //     final isOwner = extra['isOwner'] as bool;
      //     return ListOfCars(roleId: roleId, isOwner: isOwner);
      //   },
      // ),
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
