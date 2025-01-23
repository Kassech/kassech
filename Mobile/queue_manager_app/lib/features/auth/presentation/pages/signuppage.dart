import 'package:file_picker/file_picker.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/mytextfield.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/filePicker.dart';
import 'package:queue_manager_app/config/route/route.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/password.dart';

import '../../domain/models/user_params.dart';
import '../providers/auth_provider.dart';
import '../providers/user_data_provider.dart';

class SignUpPage extends StatefulWidget {
  final int role;

  const SignUpPage({required this.role, super.key});

  @override
  State<SignUpPage> createState() => _SignUpPageState();
}

class _SignUpPageState extends State<SignUpPage> {
  final TextEditingController nameController = TextEditingController();
  final TextEditingController fathersNameController = TextEditingController();
  final TextEditingController emailController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();
  final TextEditingController phoneController = TextEditingController();
  final TextEditingController kebeleIdController = TextEditingController();
  final TextEditingController profilePictureController =
      TextEditingController();

  final formKey = GlobalKey<FormState>();

  Future<void> pickFile(String fileType, WidgetRef ref) async {
    FilePickerResult? result;
    try {
      result = await FilePicker.platform.pickFiles();
    } catch (e) {
      return;
    }

    if (result != null) {
      final file = result.files.first;
      ref.read(userDataProvider.notifier).updateUserFiles(
            kebeleId: fileType == 'kebeleId' ? file : null,
            profilePicture: fileType == 'profilePicture' ? file : null,
            drivingLicenseFile: fileType == 'drivingLicense' ? file : null,
            insuranceDocumentFile:
                fileType == 'insuranceDocument' ? file : null,
          );
    }
  }

  @override
  void dispose() {
    // TODO: implement dispose
    super.dispose();
    nameController.dispose();
    fathersNameController.dispose();
    emailController.dispose();
    passwordController.dispose();
    phoneController.dispose();
    kebeleIdController.dispose();
    profilePictureController.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.white,
        leading: IconButton(
          icon: const Icon(Icons.arrow_back),
          onPressed: () {
            AppRouter.router.go('/signin');
          },
        ),
      ),
      body: Form(
        key: formKey, // Set the form key.
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Consumer(
            builder: (context, ref, child) {
              return ListView(
                children: [
                  Column(
                    children: [
                      const Text(
                        'Sign Up',
                        textAlign: TextAlign.left,
                        style: TextStyle(
                          fontSize: 40,
                          fontWeight: FontWeight.bold,
                        ),
                      ),
                      const SizedBox(
                        height: 20,
                      ),
                      MyTextField(
                        labelText: 'First Name',
                        controller: nameController,
                        hintText: 'First Name',
                        validator: (val) {
                          return val.isEmpty ? 'Enter your first name' : null;
                        },
                      ),
                      const SizedBox(
                        height: 20,
                      ),
                      MyTextField(
                        labelText: 'Father\'s Name',
                        controller: fathersNameController,
                        hintText: 'Father\'s Name',
                        validator: (val) {
                          return val.isEmpty
                              ? 'Enter your father\'s name'
                              : null;
                        },
                      ),
                      const SizedBox(
                        height: 20,
                      ),
                      MyTextField(
                        labelText: 'Email',
                        controller: emailController,
                        hintText: 'Email',
                        validator: (val) {
                          return val.isEmpty ? 'Enter your email' : null;
                        },
                      ),
                      const SizedBox(
                        height: 20,
                      ),
                      MyPasswordTextField(
                        labelText: 'Password',
                        controller: passwordController,
                        hintText: 'Password',
                        validator: (val) {
                          return val.isEmpty ? 'Enter your password' : null;
                        },
                      ),
                      const SizedBox(
                        height: 20,
                      ),
                      MyTextField(
                        labelText: 'Phone Number',
                        controller: phoneController,
                        hintText: 'Phone Number',
                        validator: (val) {
                          return val.isEmpty ? 'Enter your phone number' : null;
                        },
                      ),
                      const SizedBox(
                        height: 20,
                      ),
                      FileSelectorWidget(
                        label: 'Kebele ID File',
                        pickFile: () => pickFile('kebeleId', ref),
                        filePath: ref.read(userDataProvider)?.kebeleId?.path,
                      ),
                      FileSelectorWidget(
                        label: 'Profile Picture',
                        pickFile: () => pickFile('profilePicture', ref),
                        filePath: ref.watch(userDataProvider)?.profilePicture?.path,
                      ),
                      if (widget.role == 2) ...[
                        FileSelectorWidget(
                          label: 'Driving License File',
                          pickFile: () => pickFile('drivingLicense', ref),
                          filePath: ref.watch(userDataProvider)?.drivingLicenseFile?.path,
                        ),
                        if (widget.role == 4) ...[
                          FileSelectorWidget(
                            label: 'Insurance Document File',
                            pickFile: () => pickFile('insuranceDocument', ref),
                            filePath: ref.watch(userDataProvider)?.insuranceDocumentFile?.path,
                          ),
                        // Row(
                        //   children: [
                        //     Radio(
                        //       value: true,
                        //       groupValue: ref.watch(userDataProvider)?.isAlsoDriver,
                        //       onChanged: (val) {
                        //         ref.read(userDataProvider.notifier).updateUserParams(
                        //               isAlsoDriver: val as bool,
                        //             );
                        //       },
                        //     ),
                        //     const Text('Is also driver')
                        //   ],
                        // ),
                        const SizedBox(
                          height: 10,
                        ),
                        ]
                      ],
                      ElevatedButton(
                        onPressed: () {
                          if (formKey.currentState!.validate()) {
                            final user = UserParams(
                              firstName: nameController.text,
                              lastName: fathersNameController.text,
                              email: emailController.text,
                              phoneNumber: phoneController.text,
                              password: passwordController.text,
                              role: widget.role,
                              kebeleId: ref.read(userDataProvider)?.kebeleId,
                              profilePicture: ref.read(userDataProvider)?.profilePicture,
                              drivingLicenseFile: ref.read(userDataProvider)?.drivingLicenseFile,
                              insuranceDocumentFile: ref.read(userDataProvider)?.insuranceDocumentFile,
                            );
                            ref.read(userDataProvider.notifier).updateUserData(user);
                            ref.read(authProvider.notifier).signUp(user);
                          }
                        },
                        child: const Text('Sign Up'),
                      ),
                    ],
                  ),
                ],
              );
            },
          ),
        ),
      ),
    );
  }
}
