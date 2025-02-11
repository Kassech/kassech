import 'package:flutter/material.dart';

class CustomNavigationBar extends StatelessWidget {
  final List<IconData> icons;
  final int selectedIndex;
  final List<String>? labels;
  final double? height;
  final Color? backgroundColor;
  final Color? surfaceTintColor;
  final Color? selectedIconColor;
  final Color? unselectedIconColor;
  final Color? selectedLabelColor;
  final Color? unselectedLabelColor;
  final NavigationDestinationLabelBehavior? labelBehavior;
  final double? borderRadius;
  final double? iconSize;
  final double? labelFontSize;
  final double? elevation;
  final Color? shadowColor;
  final ValueChanged<int>? onDestinationSelected;
  final Color? indicatorColor;
  final ShapeBorder? indicatorShape;
  final bool isFloating;

  const CustomNavigationBar({
    super.key,
    required this.icons,
    required this.selectedIndex,
    this.labels,
    this.height,
    this.backgroundColor,
    this.surfaceTintColor,
    this.selectedIconColor,
    this.unselectedIconColor,
    this.selectedLabelColor,
    this.unselectedLabelColor,
    this.labelBehavior,
    this.borderRadius,
    this.iconSize,
    this.labelFontSize,
    this.elevation,
    this.shadowColor,
    this.onDestinationSelected,
    this.indicatorColor,
    this.indicatorShape,
    this.isFloating = true,
  }) : assert(
          labels == null || icons.length == labels.length,
          'Error in CustomNavigationBar: The number of icons (${icons.length}) and labels (${labels.length}) must match.',
        );

  @override
  Widget build(BuildContext context) {
    final themeData = Theme.of(context);
    return Container(
      margin: isFloating
          ? const EdgeInsets.symmetric(horizontal: 20, vertical: 20)
          : null,
      decoration: BoxDecoration(
        color: backgroundColor ?? themeData.colorScheme.surface,
        borderRadius: BorderRadius.circular(
          borderRadius != null
              ? borderRadius!
              : isFloating
                  ? 20
                  : 0,
        ),
      ),
      child: ClipRRect(
        borderRadius: BorderRadius.circular(
          borderRadius != null
              ? borderRadius!
              : isFloating
                  ? 20
                  : 0,
        ),
        child: NavigationBar(
          elevation: elevation,
          backgroundColor: backgroundColor,
          shadowColor: shadowColor,
          height: height,
          labelBehavior: labelBehavior ?? NavigationDestinationLabelBehavior.alwaysHide,
          surfaceTintColor: surfaceTintColor,
          selectedIndex: selectedIndex,
          onDestinationSelected: onDestinationSelected,
          indicatorColor: indicatorColor,
          indicatorShape: indicatorShape,
          destinations: [
            for (var i = 0; i < icons.length; i++)
              NavigationDestination(
                icon: Icon(
                  icons[i],
                  size: iconSize ?? 22,
                  color: i == selectedIndex
                      ? selectedIconColor ?? themeData.colorScheme.primary
                      : unselectedIconColor ?? (themeData.brightness == Brightness.light ? themeData.colorScheme.secondary : themeData.colorScheme.onSurface),
                ),
                label: labels != null ? labels![i] : '',
              ),
          ],
        ),
      ),
    );
  }
}
