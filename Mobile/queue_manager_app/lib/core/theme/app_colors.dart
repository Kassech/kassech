import 'package:flutter/material.dart';

class AppColors {
  /// Common colors
  static const white = Colors.white;
  static const black = Colors.black;
  static const lightGray = Color(0xFFE0E0E0);
  static const darkGray = Color(0xFF5A5858);
  static const blue = Color(0xFF0D47A1);

  /// Light theme colors
  static const lightPrimary = Colors.black;
  static const lightSecondary = Colors.white;
  static const lightOnPrimary = Colors.white;
  static const lightOnSecondary = Colors.black;
  static const lightScaffoldBackground = Colors.white;
  static const lightNavBarBackground = lightGray;

  /// Dark theme colors
  static const darkPrimary = Colors.white;
  static const darkSecondary = Colors.black;
  static const darkOnPrimary = Colors.black;
  static const darkOnSecondary = Colors.white;
  static const darkScaffoldBackground = Colors.black;
  static const darkNavBarBackground = darkGray;

  /// Status colors
  static const successColor = Color(0xFF00C853); // Green
  static const warningColor = Color(0xFFFFA000); // Orange
  static const errorColor = Color(0xFFD32F2F); // Red

  /// Car colors
  static const carRed = Colors.red;
  static const carBlue = Color(0xFF0D47A1); // Same as blue
  static const carGreen = Color(0xFF00C853); // Same as successColor
  static const carYellow = Colors.yellow;
  static const carOrange = Colors.orange;
  static const carPurple = Colors.purple;
  static const carBlack = Colors.black;
  static const carWhite = Colors.white;
  static const carGray = Color(0xFF9E9E9E); // G
  
 static Color getCarColor(String colorName) {
    switch (colorName.toLowerCase()) {
      case 'red':
        return AppColors.carRed;
      case 'blue':
        return AppColors.carBlue;
      case 'green':
        return AppColors.carGreen;
      case 'yellow':
        return AppColors.carYellow;
      case 'orange':
        return AppColors.carOrange;
      case 'purple':
        return AppColors.carPurple;
      case 'black':
        return AppColors.carBlack;
      case 'white':
        return AppColors.carWhite;
      case 'gray':
        return AppColors.carGray;
      default:
        return AppColors.lightGray; // Default color if no match
    }
  }
}
