import 'package:flutter/material.dart';

import 'app_colors.dart';

/// Light theme data
final lightThemeData = ThemeData(
  primaryColor: AppColors.lightPrimary,
  scaffoldBackgroundColor: AppColors.lightScaffoldBackground,
  colorScheme: const ColorScheme.light(
    primary: AppColors.lightPrimary,
    onPrimary: AppColors.lightOnPrimary,
    secondary: AppColors.lightSecondary,
    onSecondary: AppColors.lightOnSecondary,
    error: AppColors.errorColor,
  ),
  appBarTheme: const AppBarTheme(
    centerTitle: true,
  ),
  inputDecorationTheme: InputDecorationTheme(
    enabledBorder: OutlineInputBorder(
      borderRadius: BorderRadius.all(Radius.circular(5)),
      borderSide: BorderSide(color: AppColors.darkGray),
    ),
    focusedBorder: OutlineInputBorder(
      borderSide: BorderSide(color: AppColors.black),
      borderRadius: BorderRadius.all(Radius.circular(5)),
    ),
    errorBorder: OutlineInputBorder(
      borderSide: BorderSide(color: AppColors.errorColor),
      borderRadius: BorderRadius.all(Radius.circular(5)),
    ),
    fillColor: AppColors.white,
    filled: true,
  ),
  elevatedButtonTheme: ElevatedButtonThemeData(
    style: ButtonStyle(
      minimumSize:
      WidgetStateProperty.all<Size>(const Size(double.infinity, 58)),
      backgroundColor: WidgetStateProperty.resolveWith<Color>(
            (Set<WidgetState> state) => AppColors.lightPrimary,
      ),
      foregroundColor: WidgetStateProperty.resolveWith<Color>(
            (Set<WidgetState> state) => AppColors.lightOnPrimary,
      ),
      shape: WidgetStateProperty.all<RoundedRectangleBorder>(
        RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(5.0),
        ),
      ),
    ),
  ),
);

/// Dark theme data
final darkThemeData = ThemeData(
  primaryColor: AppColors.darkPrimary,
  scaffoldBackgroundColor: AppColors.darkScaffoldBackground,
  colorScheme: const ColorScheme.dark(
    primary: AppColors.darkPrimary,
    onPrimary: AppColors.darkOnPrimary,
    secondary: AppColors.darkSecondary,
    onSecondary: AppColors.darkOnSecondary,
    error: AppColors.errorColor,
  ),
  appBarTheme: const AppBarTheme(
    centerTitle: true,
  ),
  inputDecorationTheme: InputDecorationTheme(
    enabledBorder: OutlineInputBorder(
      borderRadius: BorderRadius.all(Radius.circular(5)),
      borderSide: BorderSide(color: AppColors.lightGray),
    ),
    focusedBorder: OutlineInputBorder(
      borderSide: BorderSide(color: AppColors.white),
      borderRadius: BorderRadius.all(Radius.circular(5)),
    ),
    errorBorder: OutlineInputBorder(
      borderSide: BorderSide(color: AppColors.errorColor),
      borderRadius: BorderRadius.all(Radius.circular(5)),
    ),
    fillColor: AppColors.black,
    filled: true,
  ),
  elevatedButtonTheme: ElevatedButtonThemeData(
    style: ButtonStyle(
      minimumSize:
      WidgetStateProperty.all<Size>(const Size(double.infinity, 58)),
      backgroundColor: WidgetStateProperty.resolveWith<Color>(
            (Set<WidgetState> state) => AppColors.darkPrimary,
      ),
      foregroundColor: WidgetStateProperty.resolveWith<Color>(
            (Set<WidgetState> state) => AppColors.darkOnPrimary,
      ),
      shape: WidgetStateProperty.all<RoundedRectangleBorder>(
        RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(5.0),
        ),
      ),
    ),
  )
);
