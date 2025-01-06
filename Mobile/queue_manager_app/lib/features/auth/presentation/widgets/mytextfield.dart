import 'package:flutter/material.dart';

class MyTextField extends StatelessWidget {
  final String labelText;
  final String hintText;
  final TextEditingController controller;

  const MyTextField(
      {super.key,
      required this.labelText,
      required this.controller,
      required this.hintText});

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 25.0),
      child: TextField(
        controller: controller,
        decoration: InputDecoration(
            labelText: labelText,
            hintText: hintText,
            enabledBorder: const OutlineInputBorder(
              borderRadius: BorderRadius.all(Radius.circular(5)),
              borderSide: BorderSide(color: Colors.black),
            ),
            focusedBorder: const OutlineInputBorder(
                borderSide: BorderSide(color: Colors.grey),
                borderRadius: BorderRadius.all(Radius.circular(5))),
            fillColor: Colors.white,
            filled: true),
      ),
    );
  }
}
