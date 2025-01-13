import 'package:flutter/material.dart';
import 'package:queue_manager_app/features/queue/presentation/widgets/appDrawer.dart';

class NoRoutesAssignedYet extends StatelessWidget {
  const NoRoutesAssignedYet({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: IconButton( onPressed: () { AppDrawer(); }, icon: const Icon(Icons.menu),),
      ),
    );
  }
}