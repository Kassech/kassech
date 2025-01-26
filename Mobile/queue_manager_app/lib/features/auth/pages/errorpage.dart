import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import 'signinpage.dart';

class ErrorPage extends StatelessWidget {
  const ErrorPage(error, {super.key});

  static const String routeName = '/errorPage';

  @override
  Widget build(BuildContext context) {
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
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(
              'Something went wrong',
              style: TextStyle(fontSize: 40, fontWeight: FontWeight.w700),
            )
          ],
        ),
      ),
    );
  }
}
