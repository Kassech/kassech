import 'package:dio/dio.dart';

final dio = Dio();
String baseUrl = 'http://192.168.169.158:5000/api/';

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
    final response =
        await dio.post('${baseUrl}login', data: formData);
    print(response);
    return response;
  } catch (e) {
    print(e);
    throw e;
  }
}
