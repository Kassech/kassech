// import 'package:dio/dio.dart';
// import 'package:driver_app/core/network/network_service.dart';


// Future<Response> login(String phoneNumber, String Password) async {
//   final formData = {
//     'email_or_phone': phoneNumber,
//     'Password': Password,
//   };

//   try {
//     final appService = NetworkService();

//     final response = await appService.dio.post('${baseUrl}login', data: formData);
//     print(response);
//     return response;
//   } catch (e) {
//     print(e);
//     throw e;
//   }
// }
