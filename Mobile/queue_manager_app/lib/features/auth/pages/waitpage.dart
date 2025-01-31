import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import 'signinpage.dart';

class WaitPage extends StatelessWidget {
  const WaitPage({super.key});

  static const String routeName = '/waitPage';

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Select Role'),
        leading: IconButton(
          onPressed: () {
            context.go(SignInPage.routeName);
          },
          icon: const Icon(Icons.arrow_back),
        ),
      ),
      body:  Padding(
        padding: const EdgeInsets.symmetric(horizontal: 25.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text(
              'Please Wait Till Verified',
              textAlign: TextAlign.left,
              style: TextStyle(
                fontSize: 40,
                fontWeight: FontWeight.w800,
              ),
            ),
            SizedBox(height: 20),
           Center(
              child: CircularProgressIndicator()
            ),
          ],
        ),
      ),
    );
  }
}
