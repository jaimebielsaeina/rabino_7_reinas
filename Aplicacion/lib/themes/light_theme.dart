import 'package:flutter/material.dart';

class LightTheme {
  static final lightTheme = ThemeData(
    appBarTheme:  const AppBarTheme(
      color: Colors.indigo,
      centerTitle: true,
      titleTextStyle: TextStyle(
        color: Colors.white,
        fontSize: 25
      ),
        iconTheme: IconThemeData(
        color: Colors.white
      ),
        actionsIconTheme: IconThemeData(
        color: Colors.white
      ),
    ),
    tabBarTheme: const TabBarTheme(
      labelColor: Colors.amber,
      unselectedLabelColor: Colors.white,
      indicatorColor: Colors.amber,
    ),
    bottomNavigationBarTheme: const BottomNavigationBarThemeData(
      showUnselectedLabels: false,
      backgroundColor: Colors.indigo,
      selectedItemColor: Colors.amber,
      unselectedItemColor: Colors.white,
    ),
    colorSchemeSeed: Colors.indigo,
    brightness: Brightness.light,
    useMaterial3: true,
  );
}