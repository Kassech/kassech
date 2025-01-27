import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:lottie/lottie.dart';
import 'package:queue_manager_app/core/theme/app_colors.dart';

import '../auth/pages/signinpage.dart';

class Splash extends StatefulWidget {
  const Splash({super.key});

  static const String routeName = '/splash';

  @override
  State<Splash> createState() => _SplashState();
}

class _SplashState extends State<Splash> {
  @override
  void initState() {
    super.initState();
    final ctx = context;
    Future.delayed(
      const Duration(seconds: 3),
      () {
        print('Splash');
        ctx.go(SignInPage.routeName);
      }
    );
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
