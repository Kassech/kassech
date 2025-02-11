import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../models/path_model.dart';
import '../repositories/driver_path_repo.dart';

final driverPathNotifierProvider = StateNotifierProvider<PathNotifier, AsyncValue<PathModel?>>((ref) {
  final repo = ref.watch(pathRepositoryProvider);
  return PathNotifier(repo);
});

class PathNotifier extends StateNotifier<AsyncValue<PathModel?>> {
  final DriverPathRepo repository;

  PathNotifier(this.repository) : super(const AsyncValue.loading()) {
    _listenToPathUpdates();
  }

  void _listenToPathUpdates() {
    repository.pathStream.listen((path) {
      if (path != null) {
        state = AsyncValue.data(path);
      }
    }, onError: (error) {
      state = AsyncValue.error(error, StackTrace.current);
    });
  }

  @override
  void dispose() {
    repository.dispose();
    super.dispose();
  }
}
