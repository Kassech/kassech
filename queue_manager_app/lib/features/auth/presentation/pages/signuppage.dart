import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/mytextfield.dart';

class SignUpPage extends StatefulWidget {
  SignUpPage({super.key});
  final TextEditingController nameController = TextEditingController();

  @override
  State<SignUpPage> createState() => _SignUpPageState();
}

class _SignUpPageState extends State<SignUpPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: SafeArea(
      child: Center(
        child: SingleChildScrollView(
            child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text(
              'Register',
              textAlign: TextAlign.left,
              style: TextStyle(
                fontSize: 40,
                fontWeight: FontWeight.w800,
              ),
            ),
            const SizedBox(
              height: 5,
            ),
            MyTextField(
                labelText: 'First Name',
                controller: widget.nameController,
                hintText: 'Name'),

            const SizedBox(
              height: 5,
            ),
            MyTextField(
                labelText: 'Last Name',
                controller: widget.nameController,
                hintText: 'Name'),
            const SizedBox(
              height: 5,
            ),
            MyTextField(
                labelText: 'Station',
                controller: widget.nameController,
                hintText: 'St.George'),
            const SizedBox(
              height: 5,
            ),
            MyTextField(
                labelText: 'Phone Number',
                controller: widget.nameController,
                hintText: '0912345678'),

            //Managed routes: from backend
          ],
        )),
      ),
    ));
  }
}
