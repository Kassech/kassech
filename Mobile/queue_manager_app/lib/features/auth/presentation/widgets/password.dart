import 'package:flutter/material.dart';

class MyPasswordTextField extends StatefulWidget {
  final String labelText;
  final String hintText;
  final TextEditingController controller;

  const MyPasswordTextField({
    super.key,
    required this.labelText,
    required this.controller,
    required this.hintText, required String? Function(dynamic val) validator,
  });

  @override
  _MyPasswordTextFieldState createState() => _MyPasswordTextFieldState();
}

class _MyPasswordTextFieldState extends State<MyPasswordTextField> {
  bool _isObscure = true; // This variable will control password visibility

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 20.0),
      child: TextFormField(
        controller: widget.controller,
        obscureText: _isObscure, // If true, password will be obscured
        decoration: InputDecoration(
          labelText: widget.labelText,
          hintText: widget.hintText,
          enabledBorder: const OutlineInputBorder(
            borderRadius: BorderRadius.all(Radius.circular(5)),
            borderSide: BorderSide(color: Colors.black),
          ),
          focusedBorder: const OutlineInputBorder(
            borderSide: BorderSide(color: Colors.grey),
            borderRadius: BorderRadius.all(Radius.circular(5)),
          ),
          fillColor: Colors.white,
          filled: true,
          suffixIcon: IconButton(
            icon: Icon(
              _isObscure
                  ? Icons.visibility_off
                  : Icons.visibility, // Toggle icon based on visibility
              color: Colors.grey,
            ),
            onPressed: () {
              setState(() {
                _isObscure = !_isObscure; // Toggle the obscured state
              });
            },
          ),
        ),
      ),
    );
  }
}
