import 'route_model.dart';

class PathModel {
  final int id;
  final String pathName;
  final double distanceKm;
  final String estimatedTime;
  final bool isActive;
  final RouteModel route;

  PathModel({
    required this.id,
    required this.pathName,
    required this.distanceKm,
    required this.estimatedTime,
    required this.isActive,
    required this.route,
  });

  factory PathModel.fromJson(Map<String, dynamic> json) {
    return PathModel(
      id: json['ID'],
      pathName: json['path_name'],
      distanceKm: (json['distance_km'] as num).toDouble(),
      estimatedTime: json['estimated_time'],
      isActive: json['is_active'],
      route: RouteModel.fromJson(json['route']),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'ID': id,
      'path_name': pathName,
      'distance_km': distanceKm,
      'estimated_time': estimatedTime,
      'is_active': isActive,
      'route': route.toJson(),
    };
  }
}
