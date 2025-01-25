import 'package:flutter/material.dart';
import 'package:dio/dio.dart';

class DelegationPage extends StatefulWidget {
  const DelegationPage({super.key});

  static const String routeName = '/delegationPage';

  @override
  State<DelegationPage> createState() => _DelegationPageState();
}

class _DelegationPageState extends State<DelegationPage> {
  final TextEditingController _searchController = TextEditingController();
  List<Driver> _drivers = [];
  List<Driver> _filteredDrivers = [];
  bool _isLoading = true;

  @override
  void initState() {
    super.initState();
    _fetchDrivers();
  }

  Future<void> _fetchDrivers() async {
    try {
      final response = await Dio().get('https://yourapi.com/drivers');
      final List<dynamic> data = response.data;
      setState(() {
        _drivers = data.map((json) => Driver.fromJson(json)).toList();
        _filteredDrivers = _drivers;
        _isLoading = false;
      });
    } catch (e) {
      setState(() {
        _isLoading = false;
      });
    }
  }

  void _filterDrivers(String query) {
    final filtered = _drivers.where((driver) {
      final nameLower = driver.name.toLowerCase();
      final phoneLower = driver.phone.toLowerCase();
      final searchLower = query.toLowerCase();
      return nameLower.contains(searchLower) ||
          phoneLower.contains(searchLower);
    }).toList();

    setState(() {
      _filteredDrivers = filtered;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Delegation'),
      ),
      body: _isLoading
          ? Center(child: CircularProgressIndicator())
          : Column(
              children: [
                Padding(
                  padding: const EdgeInsets.all(16.0),
                  child: TextField(
                    controller: _searchController,
                    decoration: InputDecoration(
                      labelText: 'Search by name or phone number',
                      border: OutlineInputBorder(),
                    ),
                    onChanged: _filterDrivers,
                  ),
                ),
                Expanded(
                  child: ListView.builder(
                    itemCount: _filteredDrivers.length,
                    itemBuilder: (context, index) {
                      final driver = _filteredDrivers[index];
                      return ListTile(
                        title: Text(driver.name),
                        subtitle: Text(driver.phone),
                      );
                    },
                  ),
                ),
              ],
            ),
    );
  }
}

class Driver {
  final String name;
  final String phone;

  Driver({required this.name, required this.phone});

  factory Driver.fromJson(Map<String, dynamic> json) {
    return Driver(
      name: json['name'],
      phone: json['phone'],
    );
  }
}
