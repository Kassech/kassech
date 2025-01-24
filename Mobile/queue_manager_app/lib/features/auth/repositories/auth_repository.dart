import 'package:dio/dio.dart';

import '../../../core/services/api_service.dart';
import '../models/user.dart';
import '../models/user_params.dart';


class AuthRepository {
  final dio = ApiService.dio;
  Future<Map<String, dynamic>?> signUp(UserParams user) async {
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
        '${ApiService().dio_baseUrl}register',
        data: formData,
        options: Options(
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        ),
      );

      return {
        'token': response.data['token'],
        'user': User.fromJson(response.data['user'] as Map<String, dynamic>),

      };
    } on DioException catch (e) {
      rethrow;
    } catch (e) {
      rethrow;
    }
  }

  /// sign in
}
