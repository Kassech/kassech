class User {
  final int id;
  final String firstName;
  final String lastName;
  final String email;
  final String phoneNumber;
  final String profilePictureUrl;
  final bool isVerified;
  final List<String> roles;
  final List<String> permissions;
  final int role;

  User({
    required this.id,
    required this.firstName,
    required this.lastName,
    required this.email,
    required this.phoneNumber,
    required this.profilePictureUrl,
    required this.isVerified,
    required this.roles,
    required this.permissions,
    required this.role,
  });

  factory User.fromJson(Map<String, dynamic> json) {
    return User(
      id: json['id'] ?? 0,
      firstName: json['first_name'] ?? '',
      lastName: json['last_name'] ?? '',
      email: json['email'] ?? '',
      phoneNumber: json['phone_number'] ?? '',
      profilePictureUrl: json['ProfilePicture'] ?? '',
      isVerified: json['is_verified'] ?? false,
      // roles: List<String>.from(json['roles'] ?? []),
      roles: ['Owner'],
      permissions: List<String>.from(json['permissions'] ?? []),
      role: json['role'] ?? 0,
    );
  }

  Map<String,dynamic> toJson() {
    return {
      'id': id,
      'first_name': firstName,
      'last_name': lastName,
      'email': email,
      'phone_number': phoneNumber,
      'ProfilePicture': profilePictureUrl,
      'is_verified': isVerified,
      'roles': roles,
      'permissions': permissions,
      'role': role,
    };
  }
}
