import 'dart:async';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../repositories/drivers_repo.dart';

class HireDriverNotifier extends AsyncNotifier<void> {
  Future<void> hireDriver(int driverId) async {
    // Set loading state.
    print("Hiring driver");
    state = const AsyncLoading();
    final repo = ref.read(driversRepoProvider);
    try {
      await repo.hireDriver(driverId);
      // Indicate a successful hire.
      state = const AsyncData(null);
    } catch (error, stackTrace) {
      // Capture any error encountered.
      state = AsyncError(error, stackTrace);
    }
  }

  @override
  FutureOr<void> build() async {
    // No initial build logic required.
  }
}

/// A provider exposing the HireDriverNotifier.
final hireDriverProvider =
AsyncNotifierProvider<HireDriverNotifier, void>(() => HireDriverNotifier());
