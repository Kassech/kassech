// lib/features/assign_drivers/providers/hire_driver_provider.dart
import 'dart:async';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import '../models/hire_driver_params.dart';
import '../repositories/drivers_repo.dart';

class HireDriverNotifier extends FamilyAsyncNotifier<void, HireDriverParams> {
  Future<void> hireDriver(HireDriverParams params) async {
    // Set the local (family-scoped) state to loading.
    state = const AsyncLoading();
    final repo = ref.read(driversRepoProvider);
    try {
      await repo.hireDriver(params.driverId, params.vehicleId);
      state = const AsyncData(null);
    } catch (error, stackTrace) {
      state = AsyncError(error, stackTrace);
    }
  }

  @override
  FutureOr<void> build(HireDriverParams arg) async {
    // Nothing to initialize.
  }
}

final hireDriverProvider = AsyncNotifierProviderFamily<HireDriverNotifier, void, HireDriverParams>(
      () => HireDriverNotifier(),
);
