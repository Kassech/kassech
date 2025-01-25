import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:queue_manager_app/features/auth/providers/auth_provider.dart';

import '../widgets/mytextfield.dart';
import 'forgotpassword.dart';
import 'selectRole.dart';

class SignInPage extends ConsumerWidget {
  SignInPage({super.key});

  static const String routeName = '/signInPage';

  final TextEditingController phoneController = TextEditingController();

  final TextEditingController passwordController = TextEditingController();

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final authState = ref.read(authProvider);
    return Scaffold(
      body: SafeArea(
        child: Center(
          child: SingleChildScrollView(
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                const Text(
                  'Login',
                  textAlign: TextAlign.left,
                  style: TextStyle(
                    fontSize: 40,
                    fontWeight: FontWeight.w800,
                  ),
                ),
                const SizedBox(
                  height: 5,
                ),
                const Text(
                  'Enter your phone number, email and password',
                  style: TextStyle(fontSize: 14),
                ),
                const SizedBox(height: 20),
                MyTextField(
                  labelText: "Phone Number",
                  validator: (val) =>
                      val.isEmpty ? 'Enter your phone number' : null,
                  controller: phoneController..text = "+251984852481",
                  hintText: "+251 ___ ___ ___",
                ),
                const SizedBox(height: 20),
                MyTextField(
                  labelText: "Password",
                  validator: (val) =>
                      val.isEmpty ? 'Enter your password' : null,
                  controller: passwordController..text = "test123",
                  hintText: "**********",
                  isPassword: true,
                ),
                const SizedBox(height: 10),
                Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 20.0),
                  child: Row(
                    children: [
                      TextButton(
                        child: Text(
                          "Forgot Password?",
                          style:
                              TextStyle(fontSize: 15, color: Colors.grey[800]),
                        ),
                        onPressed: () {
                          context.push(ForgotPassword.routeName);
                        },
                      ),
                    ],
                  ),
                ),
                // Padding(
                //   padding:
                //       const EdgeInsets.only(left: 20.0, right: 20.0, top: 20.0),
                //   child: ElevatedButton(
                //     onPressed: () {
                //       ref.read(authProvider.notifier).login(
                //           phoneNumber: phoneController.text,
                //           password: passwordController.text);
                //     },
                //     child: const Text('Login'),
                //   ),
                // ),
                authState.when(
                  // skipError: true,
                  data: (user) {
                    return Padding(
                      padding: const EdgeInsets.only(
                          left: 20.0, right: 20.0, top: 20.0),
                      child: ElevatedButton(
                        onPressed: () {
                          ref.read(authProvider.notifier).login(
                              phoneNumber: phoneController.text,
                              password: passwordController.text);
                        },
                        child: const Text('Login'),
                      ),
                    );
                  },
                  loading: () => const CircularProgressIndicator(),
                  error: (error, stack) {
                    WidgetsBinding.instance.addPostFrameCallback((_) {
                      /// Show error dialog
                      showDialog(
                        context: context,
                        builder: (context) {
                          return AlertDialog(
                            title: const Text('Error'),
                            content: Text(error.toString()),
                            actions: [
                              TextButton(
                                onPressed: () {
                                  Navigator.of(context).pop();
                                },
                                child: const Text('OK'),
                              )
                            ],
                          );
                        },
                      );
                    });
                    return Padding(
                      padding: const EdgeInsets.only(
                          left: 20.0, right: 20.0, top: 20.0),
                      child: ElevatedButton(
                        onPressed: () {
                          ref.read(authProvider.notifier).login(
                              phoneNumber: phoneController.text,
                              password: passwordController.text);
                        },
                        child: const Text('Login'),
                      ),
                    );
                  },
                ),
                const SizedBox(
                  height: 20,
                ),
                const Center(
                  child: Row(
                    children: [
                      Expanded(child: Divider(thickness: 1)),
                      Padding(
                        padding: EdgeInsets.symmetric(horizontal: 8.0),
                        child: Text(
                          'OR',
                          style: TextStyle(fontSize: 16),
                        ),
                      ),
                      Expanded(child: Divider(thickness: 1)),
                    ],
                  ),
                ),
                const SizedBox(
                  height: 50,
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    const Text('Don\'t have an account?',
                        style: TextStyle(fontSize: 15)),
                    TextButton(
                      onPressed: () {
                        context.go(SelectRolePage.routeName);
                      },
                      child: Text(
                        'Sign Up',
                        style: TextStyle(
                            color: Colors.blue[900],
                            fontSize: 15,
                            fontWeight: FontWeight.w900),
                      ),
                    )
                  ],
                )
              ],
            ),
          ),
        ),
      ),
    );
  }
}
