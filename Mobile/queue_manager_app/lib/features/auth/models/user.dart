class User {
  final String id;
  final String firstName;
  final String lastName;
  final String email;
  final String phoneNumber;
  final String password;
  final String profilePictureUrl;
  final String drivingLicenseUrl;
  final String insuranceDocumentUrl;
  final String kebeleIdUrl;
  final int role;

  User({
    required this.id,
    required this.firstName,
    required this.lastName,
    required this.email,
    required this.phoneNumber,
    required this.password,
    required this.profilePictureUrl,
    required this.drivingLicenseUrl,
    required this.insuranceDocumentUrl,
    required this.kebeleIdUrl,
    required this.role,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json['id'] ?? '',
      firstName: json['first_name'] ?? '',
      lastName: json['last_name'] ?? '',
      email: json['Email'] ?? '',
      phoneNumber: json['PhoneNumber'] ?? '',
      password: json['Password'] ?? '',
      profilePictureUrl: json['Profile'] ?? '',
      drivingLicenseUrl: json['DrivingLicenseFile'] ?? '',
      insuranceDocumentUrl: json['InsuranceDocumentFile'] ?? '',
      kebeleIdUrl: json['NationalIdFile'] ?? '',
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
      'Profile': profilePictureUrl,
      'DrivingLicenseFile': drivingLicenseUrl,
      'InsuranceDocumentFile': insuranceDocumentUrl,
      'NationalIdFile': kebeleIdUrl,
      'Role': role,
    };
  }
}
