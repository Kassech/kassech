import 'package:file_picker/file_picker.dart';

class UserParams {
  final String firstName;
  final String lastName;
  final String email;
  final String phoneNumber;
  final String password;
  final int role;
  final PlatformFile? kebeleId;
  final PlatformFile? profilePicture;
  final PlatformFile? drivingLicenseFile;
  final PlatformFile? insuranceDocumentFile;

  UserParams({
    this.firstName = '',
    this.lastName = '',
    this.email = '',
    this.phoneNumber = '',
    this.password = '',
    this.role = 0,
    this.kebeleId,
    this.profilePicture,
    this.drivingLicenseFile,
    this.insuranceDocumentFile,
  });

  copyWith({
    String? firstName,
    String? lastName,
    String? email,
    String? phoneNumber,
    String? password,
    int? role,
    PlatformFile? kebeleId,
    PlatformFile? profilePicture,
    PlatformFile? drivingLicenseFile,
    PlatformFile? insuranceDocumentFile,
  }) {
    return UserParams(
      firstName: firstName ?? this.firstName,
      lastName: lastName ?? this.lastName,
      email: email ?? this.email,
      phoneNumber: phoneNumber ?? this.phoneNumber,
      password: password ?? this.password,
      role: role ?? this.role,
      kebeleId: kebeleId ?? this.kebeleId,
      profilePicture: profilePicture ?? this.profilePicture,
      drivingLicenseFile: drivingLicenseFile ?? this.drivingLicenseFile,
      insuranceDocumentFile: insuranceDocumentFile ?? this.insuranceDocumentFile,
    );
  }
}
