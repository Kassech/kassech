import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../../features/auth/providers/auth_provider.dart';

class PermissionWrapper extends ConsumerWidget {
  final String requiredPermission;
  final Widget child;
  final Widget? fallback; // Optional: What to show if permission is denied

  const PermissionWrapper({
    super.key,
    required this.requiredPermission,
    required this.child,
    this.fallback,
  });

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final user = ref.watch(authProvider).value;
    final permissions = user?.permissions ?? [];
    return permissions.contains(requiredPermission)
        ? child
        : fallback ?? const SizedBox.shrink();
  }
}
