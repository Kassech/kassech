import 'package:dio/dio.dart';

import '../../../config/const/api_constants.dart';
import '../../../config/const/local_storage_constants.dart';
import '../../../core/services/api_service.dart';
import '../../../core/services/local_storage_service.dart';
import '../models/user.dart';
import '../models/user_params.dart';

class AuthRepository {
  final dio = ApiService.dio;
  final _storage = LocalStorageService();

  Future<User?> signUp(UserParams user) async {
    try {
      final formData = FormData.fromMap({
        'FirstName': user.firstName,
        'LastName': user.lastName,
        'Email': user.email,
        'Password': user.password,
        'PhoneNumber': user.phoneNumber,
        'Role': user.role,
        if (user.kebeleId != null)
          'NationalIdFile': await MultipartFile.fromFile(user.kebeleId!.path!,
              filename: user.kebeleId!.name),
        if (user.profilePicture != null)
          'Profile': await MultipartFile.fromFile(
            user.profilePicture!.path!,
            filename: user.profilePicture!.name,
          ),
        if (user.drivingLicenseFile != null)
          'DrivingLicenseFile': user.role == 2 || user.role == 4
              ? await MultipartFile.fromFile(
                  user.drivingLicenseFile!.path!,
                  filename: user.drivingLicenseFile!.name,
                )
              : null,
        if (user.insuranceDocumentFile != null)
          'InsuranceDocumentFile': user.role == 4
              ? await MultipartFile.fromFile(
                  user.insuranceDocumentFile!.path!,
                  filename: user.insuranceDocumentFile!.name,
                )
              : null,
      });

      final response = await dio.post(
        ApiConstants.register,
        data: formData,
        options: Options(
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        ),
      );

      if (response.data['accessToken'] != null) {
        await _storage.saveToken(response.data['accessToken']);
      }

      if (response.data['user'] != null) {
        await _storage.saveString(LocalStorageConstants.userKey, response.data['user'].toString());
      }

      return User.fromJson(response.data['user'] as Map<String, dynamic>);
    } on DioException catch (e) {
      rethrow;
    } catch (e) {
      rethrow;
    }
  }

  /// sign in
  Future<User?> login(
      {String? phoneNumber, String? email, required String password}) async {
    try {
      final response = await dio.post(
        ApiConstants.login,
        data: {
          'email_or_phone': phoneNumber ?? email,
          'password': password,
        },
      );

      if (response.data['accessToken'] != null) {
        await _storage.saveToken(response.data['accessToken']);
      }

      if (response.data['user'] != null) {
        await _storage.saveString(LocalStorageConstants.userKey, response.data['user'].toString());
      }

      return User.fromJson(response.data['user'] as Map<String, dynamic>);
    } on DioException catch (e) {
      rethrow;
    } catch (e) {
      rethrow;
    }
  }

  Future<void> logout() async {
    try {
      await _storage.clear();
    } catch (e) {
      rethrow;
    }
  }
}
