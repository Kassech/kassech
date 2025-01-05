import 'package:flutter/material.dart';
import 'package:queue_manager_app/config/route/route.dart';
import 'package:queue_manager_app/features/auth/domain/usecase/api_service.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/authButton.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/mytextfield.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/password.dart';

class SigninPage extends StatelessWidget {
  SigninPage({super.key});

// Controllers for reusable widgets
  final TextEditingController phoneController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();

//private instatiation of the authservice
  final ApiService _apiService = ApiService();

// Extract the refresh token from cookies
  String? extractRefreshToken(String cookies) {
    final refreshTokenPattern = RegExp(r'refresh_token=([^;]+)');
    final match = refreshTokenPattern.firstMatch(cookies);
    return match?.group(1);
  }

  Future<void> _login(BuildContext context) async {
    String? refresh_Token = '';
    
    try {
      final response = await _apiService.login(
          phoneController.text, passwordController.text);

      if (response.statusCode == 200) {
        // Check headers for 'Set-Cookie'
        final cookies = response.headers['set-cookie'];
        if (cookies != null) {
          refresh_Token = extractRefreshToken(cookies.toString());
          print('Refresh Token: $refresh_Token');
        }
        // Store the access token and refresh token
        final accessToken = response.data['accessToken'];
        final refreshToken = refresh_Token;
        await _apiService.saveTokens(accessToken, refresh_Token.toString());
        print('Login successful');
        print(accessToken);
        print(refreshToken);

        // Send the access token to the backend
        await _apiService.sendTokensToBackend(accessToken, refresh_Token.toString());
        AppRouter.router.go('/home');
        // Get notifications
        // await _apiService.getNotifications(accessToken);
      } else {
        print('Login failed, status code: ${response.statusCode}');
      }
      
    } catch (e) {
      print('Error during login: $e');
      throw e;
    }
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
                  controller: phoneController..text = "+251984852481",
                  hintText: "+251 ___ ___ ___",
                ),
                const SizedBox(height: 20),
                MyPasswordTextField(
                  labelText: "Password",
                  controller: passwordController..text = "test123",
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
                        onPressed: () {
                          AppRouter.router.go('/signup');
                        },
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
