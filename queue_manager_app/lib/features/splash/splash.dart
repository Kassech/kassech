import 'package:flutter/material.dart';
import 'package:lottie/lottie.dart';

class Splash extends StatefulWidget {
  const Splash({super.key});

  @override
  State<Splash> createState() => _SplashState();
}

class _SplashState extends State<Splash> {
  @override
  void initState() {
    super.initState();
    // Wait for 3 seconds and navigate to the next page
    Future.delayed(Duration(seconds: 2), () {
      Navigator.pushReplacementNamed(
          context, '/signin'); // Push replacement to move to sign-in
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
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
