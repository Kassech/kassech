import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:lottie/lottie.dart';
import 'package:queue_manager_app/config/const/api_constants.dart';
import 'package:queue_manager_app/core/theme/app_colors.dart';

import '../../core/services/web_socket_service.dart';
import '../auth/pages/signinpage.dart';

class Splash extends ConsumerStatefulWidget {
  const Splash({super.key});

  static const String routeName = '/splash';

  @override
  ConsumerState<Splash> createState() => _SplashState();
}

class _SplashState extends ConsumerState<Splash> {
  @override
  void initState() {
    super.initState();
    WidgetsBinding.instance.addPostFrameCallback((_) {
      final router = GoRouter.of(context);
      Future.delayed(const Duration(seconds: 3)).then((value) {
        router.go(SignInPage.routeName);
      });
      final path =  ApiConstants.status;

      WebSocketService(path);
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.white,
      body: Center(
        child: Lottie.asset(
          'assets/splash2.json', // Corrected asset path
          width: 150,
          height: 200,
          repeat: true,
        ),
      ),
    );
  }
}
