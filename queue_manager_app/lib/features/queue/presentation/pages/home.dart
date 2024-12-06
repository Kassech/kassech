import 'package:flutter/material.dart';
import 'package:flutter_osm_plugin/flutter_osm_plugin.dart';

class HomeQueueManager extends StatefulWidget {
  const HomeQueueManager({super.key});

  @override
  State<HomeQueueManager> createState() => _HomeQueueManagerState();
}

class _HomeQueueManagerState extends State<HomeQueueManager> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(leading: Icon(Icons.menu),),
      backgroundColor: Colors.white,
      body: SafeArea(child: SingleChildScrollView(
        child: ListTile(),
      ),),
    );
  }
}
