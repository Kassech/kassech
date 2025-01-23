import 'package:flutter/material.dart';

class MyTextField extends StatefulWidget {
  final String labelText;
  final String hintText;
  final bool isPassword;
  final TextEditingController controller;

  const MyTextField(
      {super.key,
      required this.labelText,
      required this.controller,
      required this.hintText,
      required String? Function(dynamic val) validator,  this.isPassword = false});

  @override
  State<MyTextField> createState() => _MyTextFieldState();
}

class _MyTextFieldState extends State<MyTextField> {
  bool _isObscure = true;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 20.0),
      child: TextFormField(
        controller: widget.controller,
        validator: (value) {
          if (value?.isEmpty ?? true) {
            return ('please enter value');
          }
          return null;
        },
        obscureText: widget.isPassword ? _isObscure : false,
        decoration: InputDecoration(
          labelText: widget.labelText,
          hintText: widget.hintText,
          suffixIcon: widget.isPassword ? IconButton(
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
          ) : null,
        ),
      ),
    );
  }
}
