import 'dart:io';

import 'package:device_info_plus/device_info_plus.dart';
import 'package:dio/dio.dart';
import 'package:firebase_messaging/firebase_messaging.dart';
import 'package:queue_manager_app/core/network/network_service.dart';

class NotificationService {
  final Dio dio = Dio();

  Future<void> sendTokenAndDeviceId() async {
    final String? token = await FirebaseMessaging.instance.getToken();
    final String? device_id =
      await _getDeviceId(); 

    final String url ='${networkService.baseUrl}notification';
    print('URL: $url');
    print('Token: $token');
    print('Device ID: $device_id');

    try {
      final response = await dio.post(
        url,
        data: {
          'token': token,
          'device_id': device_id,
        },
        options: Options(
          headers: {
            'Content-Type': 'application/json',
          },
        ),
      );
      print('Response: ${response.data}');
    } catch (e) {
      print('Error: $e');
    }
  }

   Future<String?> _getDeviceId() async {
    final DeviceInfoPlugin deviceInfo = DeviceInfoPlugin();
    if (Platform.isAndroid) {
      final AndroidDeviceInfo androidInfo = await deviceInfo.androidInfo;
      return androidInfo.id; // Unique ID on Android
    } else if (Platform.isIOS) {
      final IosDeviceInfo iosInfo = await deviceInfo.iosInfo;
      return iosInfo.identifierForVendor; // Unique ID on iOS
    } else {
      throw UnsupportedError('Unsupported platform');
    }
  }
}

class NetworkService {
}


