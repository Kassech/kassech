class Driver {
  final int id;
  final String firstName;
  final String lastName;
  final String email;
  final String phoneNumber;
  final bool isOnline;
  final String? profilePicture;
  final bool isVerified;
  final String lastLoginDate;
  final String roles;

  Driver({
    required this.id,
    required this.firstName,
    required this.lastName,
    required this.email,
    required this.phoneNumber,
    required this.isOnline,
    this.profilePicture,
    required this.isVerified,
    required this.lastLoginDate,
    required this.roles,
  });

  // Convert from JSON
  factory Driver.fromJson(Map<String, dynamic> json) {
    return Driver(
      id: json['ID'],
      firstName: json['FirstName'],
      lastName: json['LastName'],
      email: json['Email'],
      phoneNumber: json['PhoneNumber'],
      isOnline: json['IsOnline'],
      profilePicture: json['ProfilePicture'],
      isVerified: json['IsVerified'],
      lastLoginDate: json['LastLoginDate'],
      roles: json['roles'],
    );
  }

  // Search driver by first name, last name, or phone number
  static bool matchesSearchQuery(Driver driver, String query) {
    final searchQuery = query.toLowerCase();
    return driver.firstName.toLowerCase().contains(searchQuery) ||
        driver.lastName.toLowerCase().contains(searchQuery) ||
        driver.phoneNumber.contains(searchQuery);
  }
}
