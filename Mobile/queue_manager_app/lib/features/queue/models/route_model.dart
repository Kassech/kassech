import 'station_model.dart';

class RouteModel {
  final int id;
  final String name;
  final StationModel startingLocation;
  final StationModel arrivalLocation;

  RouteModel({
    required this.id,
    required this.name,
    required this.startingLocation,
    required this.arrivalLocation,
  });

  factory RouteModel.fromJson(Map<String, dynamic> json) {
    final startingLocation = StationModel.fromJson(json['station_a']);
    final arrivalLocation = StationModel.fromJson(json['station_b']);
    final name = '${startingLocation.locationName} - ${arrivalLocation.locationName}';

    return RouteModel(
      id: json['ID'],
      name: name,
      startingLocation: startingLocation,
      arrivalLocation: arrivalLocation,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'ID': id,
      'Name': name,
      'StationA': startingLocation.toJson(),
      'StationB': arrivalLocation.toJson(),
    };
  }
}
