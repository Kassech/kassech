import 'package:flutter/material.dart';
import 'package:queue_manager_app/config/route/route.dart';

import '../widgets/authButton.dart';
import '../widgets/mytextfield.dart';

class ForgotPassword extends StatelessWidget {
  ForgotPassword({super.key});
  final TextEditingController phoneController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.white,
        leading: IconButton(
            onPressed: () {
              AppRouter.router.go('/signin');
            },
            icon: const Icon(Icons.arrow_back)),
      ),
      backgroundColor: const Color.fromRGBO(255, 255, 255, 1),
      body: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          const Text(
            'Forgot Password?',
            textAlign: TextAlign.left,
            style: TextStyle(
              fontSize: 40,
              fontWeight: FontWeight.w800,
            ),
          ),
          const SizedBox(
            height: 20,
          ),
          MyTextField(
              labelText: 'Enter your phone number',
              controller: phoneController,
              validator: (val) => val.isEmpty ? 'Enter your phone number' : null,
              hintText: "+251 ___ ___ ___"),
          const SizedBox(
            height: 20,
          ),
          AuthButton(
              label: 'Send Code',
              onPressed: () {
                AppRouter.router.go('/');
              }),
        ],
      ),
    );
  }
}
