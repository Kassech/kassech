import 'package:flutter/material.dart';

class AuthButton extends StatelessWidget {
  final String label;
  final VoidCallback onPressed;
  const AuthButton({super.key, required this.label, required this.onPressed});

  @override
  Widget build(BuildContext context) {
    return Container(
        margin: const EdgeInsets.only(top: 30),
        decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(5), color: Colors.black),
        padding: const EdgeInsets.symmetric(vertical: 7, horizontal: 150),
        child: TextButton(
            onPressed: onPressed,
            child: Text(
              label,
              style: const TextStyle(fontSize: 16, color: Colors.white),
            )));
  }
}
