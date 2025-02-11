import 'package:flutter/material.dart';

import 'app_colors.dart';

/// Light theme data
final lightThemeData = ThemeData(
  primaryColor: AppColors.lightPrimary,
  scaffoldBackgroundColor: AppColors.lightScaffoldBackground,
  colorScheme: ColorScheme.light(
    brightness: Brightness.light,
    primary: AppColors.lightPrimary,
    onPrimary: AppColors.lightOnPrimary,
    secondary: AppColors.lightSecondary,
    onSecondary: AppColors.lightOnSecondary,
    error: AppColors.errorColor,
  ),
  appBarTheme: const AppBarTheme(
    centerTitle: true,
    backgroundColor: AppColors.white
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
  navigationBarTheme: NavigationBarThemeData(
    backgroundColor: AppColors.lightPrimary,
    indicatorColor: AppColors.lightSecondary,
    surfaceTintColor: Colors.transparent,
    shadowColor: Colors.red,
    elevation: 6,
    iconTheme: WidgetStateProperty.resolveWith((states) {
      return IconThemeData(
        color: states.contains(WidgetState.selected)
            ? AppColors.lightOnPrimary
            : AppColors.lightOnSecondary,
      );
    }),
  ),
);

/// Dark theme data
final darkThemeData = ThemeData(
  primaryColor: AppColors.darkPrimary,
  scaffoldBackgroundColor: AppColors.darkScaffoldBackground,
  colorScheme: const ColorScheme.dark(
    brightness: Brightness.dark,
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
  ),
  navigationBarTheme: NavigationBarThemeData(
    backgroundColor: AppColors.darkNavBarBackground,
    indicatorColor: AppColors.darkSecondary,
    shadowColor: AppColors.darkGray,
    surfaceTintColor: Colors.transparent,
    elevation: 6,
    iconTheme: WidgetStateProperty.resolveWith((states) {
      return IconThemeData(
        color: states.contains(WidgetState.selected)
            ? AppColors.darkOnPrimary
            : AppColors.darkOnSecondary,
      );
    }),
  ),
);
