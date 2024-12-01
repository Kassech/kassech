import 'package:flutter/material.dart';

class AuthButton extends StatelessWidget {
  final String label;
  const AuthButton({super.key, required this.label});

  @override
  Widget build(BuildContext context) {
    return Container(
        margin: EdgeInsets.only(top: 30),
        decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(5), color: Colors.black),
        padding: EdgeInsets.symmetric(vertical: 7, horizontal: 150),
        child: TextButton(
            onPressed: () {},
            child: Text(
              label,
              style: TextStyle(fontSize: 16, color: Colors.white),
            )));
  }
}
