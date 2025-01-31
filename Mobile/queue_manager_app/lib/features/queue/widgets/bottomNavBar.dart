import 'package:flutter/material.dart';

class BottomNavBar extends StatefulWidget {
  final Function(int) onItemTapped;
  final int selectedIndex;
  final List<String> navTitles;
  final List<String> navRoutes;

  const BottomNavBar({
    super.key,
    required this.onItemTapped,
    required this.selectedIndex,
    required this.navTitles,
    required this.navRoutes,
  });

  @override
  State<BottomNavBar> createState() => _BottomNavBarState();
}

List<IconData> icons = const [Icons.list, Icons.map, Icons.person];

class _BottomNavBarState extends State<BottomNavBar> {
  @override
  Widget build(BuildContext context) {
    return Align(
      alignment: Alignment.bottomCenter,
      child: Container(
        child: _navBar(),
      ),
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
          bool isSelected = widget.selectedIndex == index;
          return Material(
            color: Colors.transparent,
            child: GestureDetector(
              onTap: () {
                widget.onItemTapped(index);
              },
              child: SingleChildScrollView(
                child: Column(
                  children: [
                    Container(
                      alignment: Alignment.center,
                      margin: const EdgeInsets.only(
                          top: 15, bottom: 0, left: 35, right: 35),
                      child: Icon(icon,
                          color: isSelected ? Colors.black : Colors.grey),
                    ),
                    Text(widget.navTitles[index],
                        style: TextStyle(
                            color: isSelected ? Colors.black : Colors.grey,
                            fontSize: 12))
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
