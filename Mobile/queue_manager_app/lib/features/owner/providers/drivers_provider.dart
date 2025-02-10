import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../../auth/models/user.dart';
import '../repositories/drivers_repo.dart';

class DriversNotifier extends AutoDisposeFamilyAsyncNotifier<List<User>, String?> {
  List<User> _allDrivers = [];
  String _searchQuery = '';

  /// Returns the list of drivers filtered by the current search query.
  List<User> get _filteredDrivers {
    if (_searchQuery.isEmpty) return _allDrivers;
    return _allDrivers.where((driver) {
      final fullName = '${driver.firstName} ${driver.lastName}'.toLowerCase();
      return fullName.contains(_searchQuery.toLowerCase());
    }).toList();
  }

  @override
  Future<List<User>> build(String? arg) async {
    // Fetch the drivers when the provider is first created.
    final repo = ref.read(driversRepoProvider);
    _allDrivers = await repo.fetchDrivers(arg);
    return _filteredDrivers;
  }

  /// Updates the search query and refreshes the state.
  void setSearchQuery(String query) {
    _searchQuery = query;
    state = AsyncData(_filteredDrivers);
  }

  Future<void> refreshDrivers(String name) async {
    state = const AsyncLoading();
    final repo = ref.read(driversRepoProvider);
    _allDrivers = await repo.fetchDrivers(name);
    state = AsyncData(_filteredDrivers);
  }
}

final driversNotifierProvider =
    AutoDisposeAsyncNotifierProviderFamily<DriversNotifier, List<User>, String?>(
  () => DriversNotifier(),
);
