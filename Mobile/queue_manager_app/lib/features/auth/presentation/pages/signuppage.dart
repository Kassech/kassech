import 'package:dio/dio.dart';
import 'package:file_picker/file_picker.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/config/route/route.dart';
import 'package:queue_manager_app/features/auth/domain/usecase/api_service.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/authButton.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/filePicker.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/mytextfield.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/password.dart';
import 'package:queue_manager_app/features/queue/presentation/provider/filePickerNotifier.dart';
import 'package:queue_manager_app/features/queue/presentation/provider/profileChangeProvider.dart';

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
  final TextEditingController queueManagerIdController =
      TextEditingController();

  Future<void> pickFile(String fileType, WidgetRef ref) async {
    final result = await FilePicker.platform.pickFiles(type: FileType.any);
    if (result != null) {
      ref
          .read(filePickerProvider.notifier)
          .setFile(fileType, result.files.first);
    }
  }

  final dioInst = ApiService().dio_instance;

  Future<void> uploadFiles(WidgetRef ref) async {
    final files = ref.read(filePickerProvider);
    final kebeleId = files['kebeleId'];
    final profilePicture = files['profilePicture'];
    final queueManagerIdFile = files['queueManagerId'];
    final drivingLicenseFile = files['drivingLicense'];
    final insuranceDocumentFile = files['insuranceDocument'];

    if (kebeleId == null ||
        profilePicture == null ||
        (widget.role == 3 && queueManagerIdFile == null)) return;

    final formData = FormData.fromMap({
      'FirstName': nameController.text,
      'LastName': fathersNameController.text,
      'Email': emailController.text,
      'Password': passwordController.text,
      ''
          'PhoneNumber': phoneController.text,
      'Role': widget.role,
      'NationalIdFile':
          await MultipartFile.fromFile(kebeleId.path!, filename: kebeleId.name),
      'Profile': await MultipartFile.fromFile(profilePicture.path!,
          filename: profilePicture.name),
      if (widget.role == 3)
        if (widget.role == 2)
          'DrivingLicenseFile': await MultipartFile.fromFile(
              drivingLicenseFile!.path!,
              filename: drivingLicenseFile.name),
      if (widget.role == 2)
        'InsuranceDocumentFile': await MultipartFile.fromFile(
            insuranceDocumentFile!.path!,
            filename: insuranceDocumentFile.name),
    });

    try {
      final response = await dioInst.post('${ApiService().dio_baseUrl}register',
          data: formData);
      if (response.statusCode == 200) {
        final profilePictureUrl = response.data['profilePictureUrl'];
        ref
            .read(profilePictureProvider.notifier)
            .setProfilePicture(profilePictureUrl);
        print('Upload successful');
      } else {
        print('Upload failed: ${response.statusCode}');
      }
    } catch (e) {
      print('Upload failed: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        leading: IconButton(onPressed: (){AppRouter.router.go('/signin');}, icon: Icon(Icons.arrow_back)),
        backgroundColor: Colors.white,
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Consumer(
          builder: (context, ref, child) {
            final files = ref.watch(filePickerProvider);
            return ListView(
              children: [
                Column(
                  children: [
                    const Text(
                      'Sign Up',
                      textAlign: TextAlign.left,
                      style: TextStyle(
                        fontSize: 40,
                        fontWeight: FontWeight.w800,
                      ),
                    ),
                    const SizedBox(height: 20),
                    MyTextField(
                      labelText: 'Name',
                      controller: nameController,
                      hintText: 'Name',
                    ),
                    const SizedBox(height: 10),
                    MyTextField(
                      labelText: 'Fathers Name',
                      controller: fathersNameController,
                      hintText: 'Name',
                    ),
                    const SizedBox(height: 10),
                    MyTextField(
                      labelText: 'Email',
                      controller: emailController,
                      hintText: 'Email',
                    ),
                    const SizedBox(height: 10),
                    MyPasswordTextField(
                      labelText: 'Password',
                      controller: passwordController,
                      hintText: 'Password',
                      // obscureText: true,
                    ),
                    const SizedBox(height: 10),
                    MyTextField(
                      labelText: 'Phone',
                      controller: phoneController,
                      hintText: 'Phone',
                    ),
                    const SizedBox(height: 10),
                    FileSelectorWidget(
                      onPressed: () => pickFile('kebeleId', ref),
                      label: 'Upload Kebele ID',
                    ),
                    if (files['kebeleId'] != null)
                      Text(
                          'Selected Kebele ID file: ${files['kebeleId']!.name}'),
                    FileSelectorWidget(
                      onPressed: () => pickFile('profilePicture', ref),
                      label: 'Upload Profile Picture',
                    ),
                    if (files['profilePicture'] != null)
                      Text(
                          'Selected Profile Picture: ${files['profilePicture']!.name}'),
                    if (widget.role == 3) ...[
                      FileSelectorWidget(
                        onPressed: () => pickFile('queueManagerId', ref),
                        label: 'Upload Queue Manager ID',
                      ),
                      if (files['queueManagerId'] != null)
                        Text(
                            'Selected Queue Manager ID file: ${files['queueManagerId']!.name}'),
                    ],
                    if (widget.role == 2) ...[
                      FileSelectorWidget(
                        onPressed: () => pickFile('drivingLicense', ref),
                        label: 'Upload Driving License',
                      ),
                      if (files['drivingLicense'] != null)
                        Text(
                            'Selected Driving License file: ${files['drivingLicense']!.name}'),
                      FileSelectorWidget(
                        onPressed: () => pickFile('insuranceDocument', ref),
                        label: 'Upload Insurance Document',
                      ),
                      if (files['insuranceDocument'] != null)
                        Text(
                            'Selected Insurance Document file: ${files['insuranceDocument']!.name}'),
                    ],
                    AuthButton(
                      label: 'SignUp',
                      onPressed: () {
                        '${ApiService().dio_baseUrl}register';
                      },
                    ),
                  ],
                ),
              ],
            );
          },
        ),
      ),
    );
  }
}
