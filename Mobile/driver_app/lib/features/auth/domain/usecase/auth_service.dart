// // auth_service.dart

// import 'package:driver_app/core/util/token_storage.dart';

// class AuthenticationService {
//   bool _isAuthenticated = false;

//   bool get isAuthenticated => _isAuthenticated;

//   void login() {
//     _isAuthenticated = true;
//   }

//   void logout() {
//     _isAuthenticated = false;
//   }
// }
// Future<bool> checkUserAuthentication() async {
//   final token = await getAccessToken(); // Fetch the token
//   if (token == null) return false; // User is not authenticated

//   return true;
// }
// // 