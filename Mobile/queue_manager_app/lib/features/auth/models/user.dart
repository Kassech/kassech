class User {
  final String firstName;
  final String lastName;
  final String email;
  final String phoneNumber;
  final String password;
  final String profilePicture;
  final String drivingLicenseFile;
  final String insuranceDocumentFile;
  final String kebeleId;
  final int role;

  User({
    required this.firstName,
    required this.lastName,
    required this.email,
    required this.phoneNumber,
    required this.password,
    required this.profilePicture,
    required this.drivingLicenseFile,
    required this.insuranceDocumentFile,
    required this.kebeleId,
    required this.role,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      firstName: json['FirstName'] ?? '',
      lastName: json['LastName'] ?? '',
      email: json['Email'] ?? '',
      phoneNumber: json['PhoneNumber'] ?? '',
      password: json['Password'] ?? '',
      profilePicture: json['Profile'] ?? '',
      drivingLicenseFile: json['DrivingLicenseFile'] ?? '',
      insuranceDocumentFile: json['InsuranceDocumentFile'] ?? '',
      kebeleId: json['NationalIdFile'] ?? '',
      role: json['Role'] ?? 0,
    );
  }

  Map<String,dynamic> toJson() {
    return {
      'FirstName': firstName,
      'LastName': lastName,
      'Email': email,
      'PhoneNumber': phoneNumber,
      'Password': password,
      'Profile': profilePicture,
      'DrivingLicenseFile': drivingLicenseFile,
      'InsuranceDocumentFile': insuranceDocumentFile,
      'NationalIdFile': kebeleId,
      'Role': role,
    };
  }
}
