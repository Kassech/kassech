import 'package:flutter/material.dart';
import 'package:loading_animation_widget/loading_animation_widget.dart';
import 'package:queue_manager_app/config/route/route.dart';

class WaitPage extends StatelessWidget {
  const WaitPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        title: const Text('Select Role'),
        leading: IconButton(
          onPressed: () {
            AppRouter.router.go('/signin');
          },
          icon: const Icon(Icons.arrow_back),
        ),
        backgroundColor: Colors.white,
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
              child: LoadingAnimationWidget.twistingDots(
                leftDotColor: Color(0xFF1A1A3F),
                rightDotColor: Color(0xFFEA3799),
                size: 200,
              ),
            ),
          ],
        ),
      ),
    );
  }
}
