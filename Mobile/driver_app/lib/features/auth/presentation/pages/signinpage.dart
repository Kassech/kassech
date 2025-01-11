import 'package:driver_app/config/route/route.dart';
// import 'package:driver_app/core/network/network_service.dart';
import 'package:driver_app/core/usecase/auth_service.dart';
import 'package:driver_app/features/auth/presentation/widgets/authButton.dart';
import 'package:driver_app/features/auth/presentation/widgets/mytextfield.dart';
import 'package:driver_app/features/auth/presentation/widgets/password.dart';
import 'package:flutter/material.dart';

class SigninPage extends StatelessWidget {
  SigninPage({super.key});
  final TextEditingController phoneController = TextEditingController();

  final TextEditingController passwordController = TextEditingController();

  Future<void> _login(BuildContext context) async {
    // Get email and password from the controllers
    final phone = phoneController.text.trim();
    final password = passwordController.text.trim();

    if (phone.isEmpty || password.isEmpty) {
      // Show error message if fields are empty
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Please enter both email and password')),
      );
      return;
    }

    try {
      final response = await login(phone, password);

      // Check if the response status is 200 (successful)
      if (response.statusCode == 200) {
        // Login successful, you can handle the 0navigation or state changes here
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Login successful!')),
        );
        AppRouter.router.go('/home');
      } else {
        // If login failed
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Invalid credentials')),
        );
      }
    } catch (e) {
      // Handle errors if the request failed (e.g., server not reachable)
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Failed to login: $e')),
      );
    }
  }

  void _goToSignUp() {
    AppRouter.router.go('/signup');
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
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
                  labelText: "Phone Number",
                  controller: phoneController,
                  hintText: "+251 ___ ___ ___",
                ),
                const SizedBox(height: 20),
                MyPasswordTextField(
                  labelText: "Password",
                  controller: passwordController,
                  hintText: "**********",
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
                          AppRouter.router.go('/forgotpassword');
                        },
                      ),
                    ],
                  ),
                ),
                AuthButton(label: "Login", onPressed: () => _login(context)),
                const SizedBox(
                  height: 20,
                ),
                const Center(
                  child: Row(
                    children: [
                      Expanded(
                          child: Divider(thickness: 1, color: Colors.black54)),
                      Padding(
                        padding: EdgeInsets.symmetric(horizontal: 8.0),
                        child: Text(
                          'OR',
                          style: TextStyle(fontSize: 16, color: Colors.black),
                        ),
                      ),
                      Expanded(
                          child: Divider(thickness: 1, color: Colors.black54)),
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
                        style: TextStyle(fontSize: 15, color: Colors.black54)),
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
    );
  }
}
