import 'package:flutter/material.dart';

import '../models/path_model.dart';

class PathCard extends StatelessWidget {
  const PathCard({super.key, required this.path});

  final PathModel path;

  @override
  Widget build(BuildContext context) {
    return Card(
      child: ListTile(
        title: Text(path.pathName),
        subtitle: Text(path.route.name),
        trailing: Text('${path.distanceKm} km'),
      ),
    );
  }
}
