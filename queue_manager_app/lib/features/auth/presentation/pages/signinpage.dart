import 'package:flutter/material.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/authButton.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/mytextfield.dart';

class SigninPage extends StatelessWidget {
  SigninPage({super.key});
  final TextEditingController emailController = TextEditingController();

  final TextEditingController passwordController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    void _goToSignUp() {
      Navigator.pushNamed(context, '/signup');
    }

    return MaterialApp(
      debugShowCheckedModeBanner: false,
      home: Scaffold(
        backgroundColor: Colors.white,
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
                    style: TextStyle(fontSize: 14, color: Colors.black54),
                  ),
                  const SizedBox(height: 20),
                  MyTextField(
                    labelText: "Email",
                    controller: emailController,
                    hintText: "john@gmail.com",
                  ),
                  const SizedBox(height: 20),
                  MyTextField(
                    labelText: "Password",
                    controller: passwordController,
                    hintText: "**********",
                  ),
                  const SizedBox(height: 10),
                  Padding(
                    padding: const EdgeInsets.only(left: 260),
                    child: Row(
                      children: [
                        Text(
                          "Forgot Password?",
                          style:
                              TextStyle(fontSize: 15, color: Colors.grey[600]),
                        ),
                      ],
                    ),
                  ),
                  const AuthButton(label: "Login"),
                  const SizedBox(
                    height: 20,
                  ),
                  const Center(
                    child: Row(
                      children: [
                        Expanded(
                            child:
                                Divider(thickness: 1, color: Colors.black54)),
                        Padding(
                          padding: EdgeInsets.symmetric(horizontal: 8.0),
                          child: Text(
                            'OR',
                            style: TextStyle(fontSize: 16, color: Colors.black),
                          ),
                        ),
                        Expanded(
                            child:
                                Divider(thickness: 1, color: Colors.black54)),
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
                          style:
                              TextStyle(fontSize: 15, color: Colors.black54)),
                      TextButton(
                          onPressed: _goToSignUp,
                          child: Text(
                            'Sign Up',
                            style: TextStyle(
                                color: Colors.blue[900],
                                fontSize: 15,
                                fontWeight: FontWeight.w900),
                          ))
                    ],
                  )
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
