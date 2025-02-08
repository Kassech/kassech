class Car {
  final int id;
  final DateTime createdAt;
  final DateTime updatedAt;
  final int typeId;
  final CarType type;
  final String licenseNumber;
  final String vin;
  final String make;
  final int year;
  final String color;
  final String carPicture;
  final String bollo;
  final String insurance;
  final String libre;
  final int ownerId;
  final String status;
  final Owner owner;

  Car({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.typeId,
    required this.type,
    required this.licenseNumber,
    required this.vin,
    required this.make,
    required this.year,
    required this.color,
    required this.carPicture,
    required this.bollo,
    required this.insurance,
    required this.libre,
    required this.ownerId,
    required this.status,
    required this.owner,
  });

  factory Car.fromJson(Map<String, dynamic> json) {
    return Car(
      id: json['ID'],
      createdAt: DateTime.parse(json['CreatedAt']),
      updatedAt: DateTime.parse(json['UpdatedAt']),
      typeId: json['TypeID'],
      type: CarType.fromJson(json['Type']),
      licenseNumber: json['LicenseNumber'],
      vin: json['VIN'],
      make: json['Make'],
      year: json['Year'],
      color: json['Color'],
      carPicture: json['CarPicture'],
      bollo: json['Bollo'],
      insurance: json['Insurance'],
      libre: json['Libre'],
      ownerId: json['OwnerID'],
      status: json['Status'],
      owner: Owner.fromJson(json['Owner']),
    );
  }
}

class CarType {
  final int id;
  final DateTime createdAt;
  final DateTime updatedAt;
  final String typeName;
  final int capacity;
  final String description;

  CarType({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.typeName,
    required this.capacity,
    required this.description,
  });

  factory CarType.fromJson(Map<String, dynamic> json) {
    return CarType(
      id: json['ID'],
      createdAt: DateTime.parse(json['CreatedAt']),
      updatedAt: DateTime.parse(json['UpdatedAt']),
      typeName: json['TypeName'],
      capacity: json['Capacity'],
      description: json['Description'],
    );
  }
}

class Owner {
  final int id;
  final DateTime createdAt;
  final DateTime updatedAt;
  final String firstName;
  final String lastName;
  final String email;
  final String phoneNumber;
  final bool isOnline;
  final bool isVerified;
  final String roles;

  Owner({
    required this.id,
    required this.createdAt,
    required this.updatedAt,
    required this.firstName,
    required this.lastName,
    required this.email,
    required this.phoneNumber,
    required this.isOnline,
    required this.isVerified,
    required this.roles,
  });

  factory Owner.fromJson(Map<String, dynamic> json) {
    return Owner(
      id: json['ID'],
      createdAt: DateTime.parse(json['CreatedAt']),
      updatedAt: DateTime.parse(json['UpdatedAt']),
      firstName: json['FirstName'],
      lastName: json['LastName'],
      email: json['Email'],
      phoneNumber: json['PhoneNumber'],
      isOnline: json['IsOnline'],
      isVerified: json['IsVerified'],
      roles: json['roles'],
    );
  }
}
