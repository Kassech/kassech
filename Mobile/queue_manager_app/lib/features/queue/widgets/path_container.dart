import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/features/queue/widgets/path_card.dart';

import '../../../core/permissions/app_permissions.dart';
import '../../../core/permissions/permission_wrapper.dart';
import '../../../shared/widgets/error_container.dart';
import '../provider/path_provider.dart';

class PathContainer extends ConsumerStatefulWidget {
  const PathContainer({super.key});

  @override
  ConsumerState<PathContainer> createState() => _PathContainerState();
}

class _PathContainerState extends ConsumerState<PathContainer> {
  @override
  void initState() {
    // TODO: implement initState
    super.initState();
    Future.delayed(Duration.zero, () {
      ref.read(pathProvider.notifier).fetchPaths();
    });
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
          final paths = ref.watch(pathProvider);
          return RefreshIndicator(
            onRefresh: () async {
              ref.read(pathProvider.notifier).fetchPaths();
            },
            child: paths.when(
                data: (path) {
                  if (path == null || path.isEmpty) {
                    return const Center(
                      child: Text('No routes found'),
                    );
                  }
                  return ListView.builder(
                    itemCount: path.length,
                    // padding: const EdgeInsets.only(top: 10),
                    itemBuilder: (context, index) {
                      return PathCard(
                        path: path[index],
                      );
                    },
                  );
                },
                loading: () => const Center(child: CircularProgressIndicator()),
                error: (error, stackTrace) {
                  return ErrorContainer(
                    errorMessageText: error.toString(),
                    onTapRetry: () {
                      ref.read(pathProvider.notifier).fetchPaths();
                    },
                  );
                }
            ),
          );
        },
      ),
    );
  }
}
