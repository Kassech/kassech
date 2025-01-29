class StationModel {
  final int id;
  final String locationName;
  final double latitude;
  final double longitude;

  StationModel({
    required this.id,
    required this.locationName,
    required this.latitude,
    required this.longitude,
  });

  factory StationModel.fromJson(Map<String, dynamic> json) {
    return StationModel(
      id: json['ID'],
      locationName: json['LocationName'],
      latitude: json['Latitude'],
      longitude: json['Longitude'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'ID': id,
      'LocationName': locationName,
      'Latitude': latitude,
      'Longitude': longitude,
    };
  }
}
