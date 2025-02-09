import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import 'signinpage.dart';

class ErrorPage extends StatelessWidget {
  final String error;
  const ErrorPage(GoException? exception, {required this.error, super.key});

  static const String routeName = '/errorPage';
  @override
  Widget build(BuildContext context) {
    final router = GoRouter.of(context);
    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
          icon: const Icon(Icons.arrow_back),
          onPressed: () {
            context.go(SignInPage.routeName);
          },
        ),
      ),
      body: Center(
        child: Padding(
          padding: const EdgeInsets.all(20.0),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Image.asset(
                'assets/error_illustration.jpg', // Add an error illustration image in your assets
                height: 400,
              ),
              const SizedBox(height: 20),
              const Text(
                'Oops! Something went wrong.',
                style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold),
                textAlign: TextAlign.center,
              ),
              const SizedBox(height: 10),
              const Text(
                'We apologize for the inconvenience. Please try again later or return to the sign-in page.',
                style: TextStyle(fontSize: 16, color: Colors.grey),
                textAlign: TextAlign.center,
              ),
              const SizedBox(height: 20),
              ElevatedButton(
                onPressed: () {
                  router.go(SignInPage.routeName);
                },
                style: ElevatedButton.styleFrom(
                  padding:
                      const EdgeInsets.symmetric(horizontal: 30, vertical: 15),
                  shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(10),
                  ),
                ),
                child: const Text(
                  'Return to Sign In',
                  style: TextStyle(fontSize: 14),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
