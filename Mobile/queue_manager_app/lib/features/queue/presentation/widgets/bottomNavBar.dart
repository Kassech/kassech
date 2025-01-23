import 'package:flutter/material.dart';

class BottomNavBar extends StatefulWidget {
  const BottomNavBar({super.key});

  @override
  State<BottomNavBar> createState() => _BottomNavBarState();
}

List<IconData> icons = const [Icons.list, Icons.map, Icons.person];

List<String> navTitles = const ['Queue', 'Map', 'Profile'];
int selectedIndex = 0;

class _BottomNavBarState extends State<BottomNavBar> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Align(alignment: Alignment.bottomCenter, child: _navBar()),
    );
  }

  Widget _navBar() {
    return Container(
      height: 65,
      margin: const EdgeInsets.only(right: 24, left: 24, bottom: 24),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(20),
        boxShadow: [
          BoxShadow(
            color: Colors.black.withAlpha(20),
            spreadRadius: 10,
            blurRadius: 20,
          ),
        ],
      ),
      child: Row(
        crossAxisAlignment: CrossAxisAlignment.center,
        mainAxisAlignment: MainAxisAlignment.center,
        children: icons.map((icon) {
          int index = icons.indexOf(icon);
          bool isSelected = selectedIndex == index;
          return Material(
            color: Colors.transparent,
            child: GestureDetector(
              onTap: () {
                setState(() {
                  selectedIndex = index;
                });
              },
              child: SingleChildScrollView(
                child: Column(
                  children: [
                    Container(
                      alignment: Alignment.center,
                      margin: const EdgeInsets.only(top: 15, bottom: 0, left: 35, right: 35),
                      child: Icon(icon, color: isSelected ? Colors.black : Colors.grey),
                    ),
                    Text(navTitles[index], style: TextStyle(color: isSelected ? Colors.black : Colors.grey, fontSize: 12))
                  ],
                ),
              ),
            ),
          );
        }).toList(),
      ),
    );
  }
}
