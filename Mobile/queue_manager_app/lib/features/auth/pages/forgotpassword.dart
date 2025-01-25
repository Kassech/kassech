import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

import '../widgets/mytextfield.dart';

class ForgotPassword extends StatelessWidget {
  ForgotPassword({super.key});
  final TextEditingController phoneController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(),
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
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 20.0),
            child: ElevatedButton(
              onPressed: () {
                context.go('/');
              },
              child: const Text('Send Code'),
            ),
          )
        ],
      ),
    );
  }
}
