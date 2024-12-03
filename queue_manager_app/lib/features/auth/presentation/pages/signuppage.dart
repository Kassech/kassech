import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:queue_manager_app/config/route/route.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/authButton.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/filePicker.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/mytextfield.dart';

class SignUpPage extends StatefulWidget {
  SignUpPage({super.key});
  final TextEditingController nameController = TextEditingController();

  @override
  State<SignUpPage> createState() => _SignUpPageState();
}

class _SignUpPageState extends State<SignUpPage> {
  String? _fileName;

  void _onFilePicked(String? fileName) {
    setState(() {
      _fileName = fileName;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          backgroundColor: Colors.white,
          leading: IconButton(
              onPressed: () {
                AppRouter.router.go('/signin');
              },
              icon: const Icon(Icons.arrow_back)),
        ),
        backgroundColor: Colors.white,
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
                  height: 8,
                ),
                MyTextField(
                    labelText: 'First Name',
                    controller: widget.nameController,
                    hintText: 'Name'),
                const SizedBox(
                  height: 8,
                ),
                MyTextField(
                    labelText: 'Last Name',
                    controller: widget.nameController,
                    hintText: 'Name'),
                const SizedBox(
                  height: 8,
                ),
                MyTextField(
                    labelText: 'Station',
                    controller: widget.nameController,
                    hintText: 'St.George'),
                const SizedBox(
                  height: 8,
                ),
                MyTextField(
                    labelText: 'Phone Number',
                    controller: widget.nameController,
                    hintText: '0912345678'),
                const SizedBox(
                  height: 12,
                ),
                const Text(
                  'Upload ID',
                  textAlign: TextAlign.left,
                  style: TextStyle(
                    fontSize: 40,
                    fontWeight: FontWeight.w800,
                  ),
                ),
                FileSelectorWidget(label: 'Upload Kebele ID'),
                FileSelectorWidget(label: 'Upload Queue Manager ID'),
                AuthButton(
                    label: 'Register',
                    onPressed: () {
                      AppRouter.router.go('/');
                    }),
              ],
            )),
          ),
        ));
  }
}
