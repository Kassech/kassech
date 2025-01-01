import 'package:dio/dio.dart';

final dio = Dio();
String baseUrl = 'http://10.0.2.2:5000/api/';

void getHttp() async {
  try {
    final response = await dio.get(baseUrl);
  } catch (e) {
    print(e);
  }
}

Future<Response> login(String phoneNumber, String Password) async {
  final formData = {
    'email_or_phone': phoneNumber,
    'Password': Password,
  };
  try {
    final response = await dio.post('${baseUrl}login', data: formData);
    print(response);
    return response;
  } catch (e) {
    print(e);
    rethrow;
  }
}

Future<Response> register(String phoneNumber, String Password) async {
  final formData = {
    'email_or_phone': phoneNumber,
    'Password': Password,
  };
  try {
    final response = await dio.post('${baseUrl}register', data: formData);
    return response;
  } catch (e) {
    print(e);
    rethrow;
  }
}

Future <Response> logout() async {
  try {
    final response = await dio.post('${baseUrl}logout');
    return response;
  } catch (e) {
    print(e);
    rethrow;
  }
}

Future <Response <dynamic>> getLocation() async {
  try {
    final response = await dio.get('${baseUrl}location');
    return response;
  } catch (e) {
    print(e);
    rethrow;
  }
}