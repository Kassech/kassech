class HireDriverParams {
  final int driverId;
  final int vehicleId;

  HireDriverParams({
    required this.driverId,
    required this.vehicleId,
  });

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
          other is HireDriverParams &&
              runtimeType == other.runtimeType &&
              driverId == other.driverId &&
              vehicleId == other.vehicleId;

  @override
  int get hashCode => driverId.hashCode ^ vehicleId.hashCode;
}
