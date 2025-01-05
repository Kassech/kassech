// import 'package:dio/dio.dart';

// class NetworkService {
//   NetworkService._privateConstructor();
//   static final NetworkService _instance = NetworkService._privateConstructor();
//   factory NetworkService() {
//     return _instance;
//   }

//   final dio = Dio();
//   String baseUrl = 'http://10.0.2.2:5000/api/';
// }

// final networkService = NetworkService();

// Future<Response> login(String phoneNumber, String Password) async {
//   final formData = {
//     'email_or_phone': phoneNumber,
//     'Password': Password,
//   };

//   try {
//     final appService = NetworkService();

//     final response =
//         await appService.dio.post('${networkService.baseUrl}login', data: formData);
//     print(response);
//     return response;
//   } catch (e) {
//     print(e);
//     throw e;
//   }
// }



// // Future<Response> register(String phoneNumber, String Password) async {
// //   final formData = {
// //     'email_or_phone': phoneNumber,
// //     'Password': Password,
// //   };
// //   try {
// //     final response = await NetworkService()
// //         .dio
// //         .post('${networkService.baseUrl}register', data: formData);
// //     return response;
// //   } catch (e) {
// //     print(e);
// //     rethrow;
// //   }
// // }

// // Future<Response> logout() async {
// //   try {
// //     final response =
// //         await NetworkService().dio.post('${networkService.baseUrl}logout');
// //     return response;
// //   } catch (e) {
// //     print(e);
// //     rethrow;
// //   }
// // }

// // Future<Response<dynamic>> getLocation() async {
// //   try {
// //     final response =
// //         await NetworkService().dio.get('${networkService.baseUrl}location');
// //     return response;
// //   } catch (e) {
// //     print(e);
// //     rethrow;
// //   }
// // }
