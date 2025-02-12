import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/features/auth/providers/auth_provider.dart';
import 'package:queue_manager_app/features/queue/models/path_model.dart';
import 'package:queue_manager_app/features/queue/widgets/path_card.dart';

import '../../../core/permissions/app_permissions.dart';
import '../../../core/permissions/permission_wrapper.dart';
import '../../../core/services/location_service.dart';
import '../../../shared/widgets/error_container.dart';
import '../../auth/models/user.dart';
import '../../location/providers/location_provider.dart';
import '../provider/driver_path_provider.dart';
import '../provider/path_provider.dart';

class PathContainer extends ConsumerStatefulWidget {
  const PathContainer({super.key});

  @override
  ConsumerState<PathContainer> createState() => _PathContainerState();
}

class _PathContainerState extends ConsumerState<PathContainer> {
  late final User? user;

  @override
  void initState() {
    super.initState();
    user = ref.read(authProvider).value;
    Future.delayed(Duration.zero, _initializeData);
  }

  Future<void> _initializeData() async {
    if (user == null) return;

    if (user!.roles.contains('QueueManager')) {
      ref.read(pathProvider.notifier).fetchPaths();
    }

    if (user!.permissions.contains(AppPermissions.sendLocation)) {
      await ref.read(locationNotifierProvider.notifier).startListening(
            122,
            1,
            user!.id,
          );
    }
  }

  @override
  Widget build(BuildContext context) {
    return PermissionWrapper(
      requiredPermission: AppPermissions.viewAssignedRoutes,
      fallback: const Center(
        child: Text('You do not have permission to view this page'),
      ),
      child: Consumer(
        builder: (context, ref, child) {
          if (user == null) {
            return const Center(child: Text('User not found'));
          }

          return RefreshIndicator(
            onRefresh: () async {
              if (user!.roles.contains('QueueManager')) {
                ref.read(pathProvider.notifier).fetchPaths();
              }
            },
            child: _buildContent(ref),
          );
        },
      ),
    );
  }

  Widget _buildContent(WidgetRef ref) {
    if (user!.roles.contains('Driver')) {
      final path = ref.watch(driverPathNotifierProvider);
      return path.maybeWhen(
        data: (path) => path == null
            ? const Center(child: Text('No routes found'))
            : PathCard(path: path),
        error: (error, stackTrace) => ErrorContainer(
          errorMessageText: error.toString(),
          onTapRetry: () {
            ref.read(driverPathNotifierProvider);
          },
        ),
        orElse: () => const Center(child: Text('No routes found')),
      );
    }

    if (user!.roles.contains('QueueManager')) {
      final paths = ref.watch(pathProvider);
      return paths.when(
        data: (paths) => paths == null || paths.isEmpty
            ? const Center(child: Text('No routes found'))
            : ListView.builder(
                itemCount: paths.length,
                itemBuilder: (context, index) => PathCard(path: paths[index]),
              ),
        loading: () => const Center(child: CircularProgressIndicator()),
        error: (error, stackTrace) => ErrorContainer(
          errorMessageText: error.toString(),
          onTapRetry: () {
            ref.read(pathProvider.notifier).fetchPaths();
          },
        ),
      );
    }

    return const Center(child: Text('No role assigned'));
  }
}
