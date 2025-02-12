import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../../../core/util/ui_utils.dart';
import '../providers/auth_provider.dart';
import '../widgets/mytextfield.dart';

class SignInPage extends StatefulWidget {
  const SignInPage({super.key});

  static const String routeName = '/signInPage';

  @override
  State<SignInPage> createState() => _SignInPageState();
}

class _SignInPageState extends State<SignInPage> {
  final TextEditingController phoneController = TextEditingController();

  final TextEditingController passwordController = TextEditingController();

  @override
  void dispose() {
    // TODO: implement dispose
    super.dispose();
    phoneController.dispose();
    passwordController.dispose();
  }

  @override
  Widget build(BuildContext context) {

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
                  controller: phoneController..text = "+251904400375",
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
                Consumer(
                  builder: (context, ref, child) {
                    ref.listen(authProvider, (previous, next) {
                      next.maybeWhen(
                        error: (error, stack) {
                          UiUtils.showSnackBar(
                            message: error.toString(),
                            isError: true,
                          );
                        },
                        orElse: () {},
                      );
                    });
                    final authState = ref.read(authProvider);
                    return authState.when(
                      skipError: true,
                      data: (user) {
                        return Padding(
                          padding: const EdgeInsets.only(
                              left: 20.0, right: 20.0, top: 20.0),
                          child: ElevatedButton(
                            onPressed: () async {
                              await ref.read(authProvider.notifier).login(
                                  phoneNumber: phoneController.text,
                                  password: passwordController.text);
                            },
                            child: const Text('Login'),
                          ),
                        );
                      },
                      loading: () => const CircularProgressIndicator(),
                      error: (error, stack) => SizedBox.shrink(),
                    );
                  },
                ),
                // const SizedBox(
                //   height: 20,
                // ),
                // const Center(
                //   child: Row(
                //     children: [
                //       Expanded(child: Divider(thickness: 1)),
                //       Padding(
                //         padding: EdgeInsets.symmetric(horizontal: 8.0),
                //         child: Text(
                //           'OR',
                //           style: TextStyle(fontSize: 16),
                //         ),
                //       ),
                //       Expanded(child: Divider(thickness: 1)),
                //     ],
                //   ),
                // ),
                // const SizedBox(
                //   height: 50,
                // ),
                // Row(
                //   mainAxisAlignment: MainAxisAlignment.center,
                //   children: [
                //     const Text('Don\'t have an account?',
                //         style: TextStyle(fontSize: 15)),
                //     TextButton(
                //       onPressed: () {
                //         context.go(SelectRolePage.routeName);
                //       },
                //       child: Text(
                //         'Sign Up',
                //         style: TextStyle(
                //             color: AppColors.blue,
                //             fontSize: 15,
                //             fontWeight: FontWeight.w900),
                //       ),
                //     )
                //   ],
                // )
              ],
            ),
          ),
        ),
      ),
    );
  }
}
