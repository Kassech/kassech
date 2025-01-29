import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:loading_animation_widget/loading_animation_widget.dart';
import 'package:queue_manager_app/core/theme/app_colors.dart';

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
              child: LoadingAnimationWidget.twistingDots(
                leftDotColor:AppColors.black,
                rightDotColor: AppColors.darkGray,
                size: 200,
              ),
            ),
          ],
        ),
      ),
    );
  }
}
