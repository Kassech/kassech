import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/features/queue/provider/path_provider.dart';

class HomePage extends ConsumerWidget {
  final List<Map<String, String>> paths;

  const HomePage({required this.paths, Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    return Scaffold(
      appBar: AppBar(title: const Text('WebSocket Example')),
      body: ListView.builder(
        itemCount: paths.length,
        itemBuilder: (context, index) => PathCard(
          pathId: paths[index]['id']!,
          pathName: paths[index]['name']!,
        ),
      ),
    );
  }
}

class PathCard extends ConsumerWidget {
  final String pathId;
  final String pathName;

  const PathCard({
    required this.pathId,
    required this.pathName,
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final amount = ref.watch(
      pathNotifierProvider.select((state) => state[pathId] ?? 0),
    );

    return Card(
      child: ListTile(
        title: Text(pathName),
        subtitle: Text('Amount: $amount'),
        trailing: _ControlButtons(
          onIncrement: () => ref.read(pathNotifierProvider.notifier).updateCount(pathId, 1),
          onDecrement: () => ref.read(pathNotifierProvider.notifier).updateCount(pathId, -1),
        ),
      ),
    );
  }
}

class _ControlButtons extends StatelessWidget {
  final VoidCallback onIncrement;
  final VoidCallback onDecrement;

  const _ControlButtons({
    required this.onIncrement,
    required this.onDecrement,
  });

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisSize: MainAxisSize.min,
      children: [
        IconButton(icon: const Icon(Icons.remove), onPressed: onDecrement),
        IconButton(icon: const Icon(Icons.add), onPressed: onIncrement),
      ],
    );
  }
}
