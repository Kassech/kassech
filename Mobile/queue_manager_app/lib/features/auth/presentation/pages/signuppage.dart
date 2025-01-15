import 'package:dio/dio.dart';
import 'package:file_picker/file_picker.dart';
import 'package:flutter/material.dart';
import 'package:queue_manager_app/config/route/route.dart';
import 'package:queue_manager_app/features/auth/domain/usecase/api_service.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/authButton.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/filePicker.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/mytextfield.dart';

class SignUpPage extends StatefulWidget {
  SignUpPage({super.key});

  @override
  State<SignUpPage> createState() => _SignUpPageState();
}

class _SignUpPageState extends State<SignUpPage> {
  final TextEditingController nameController = TextEditingController();
  final TextEditingController emailController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();
  final TextEditingController phoneController = TextEditingController();
  final TextEditingController stationController = TextEditingController();
  final TextEditingController queueManagerIdController =
      TextEditingController();
  PlatformFile? kebeleId;
  PlatformFile? profilePicture;

  Future<void> pickFile(String fileType) async {
    final result = await FilePicker.platform.pickFiles(type: FileType.any);
    if (result != null) {
      setState(() {
        if (fileType == 'kebeleId') {
          kebeleId = result.files.first;
        } else if (fileType == 'profilePicture') {
          profilePicture = result.files.first;
        }
      });
    }
  }

  Future<void> pickProfilePicture() async {
    final result = await FilePicker.platform.pickFiles(type: FileType.image);
    if (result != null) {
      setState(() {
        profilePicture = result.files.first;
      });
    }
  }

  Future<void> uploadFiles() async {
    if (kebeleId == null || profilePicture == null) return;

    final dioReg = ApiService().dio_instance;
    final dioBase = await ApiService().dio_baseUrl;
    final formData = FormData.fromMap({
      'name': nameController.text,
      'email': emailController.text,
      'password': passwordController.text,
      'phone': phoneController.text,
      'station': stationController.text,
      'kebeleIdText': await MultipartFile.fromFile(kebeleId!.path!,
          filename: kebeleId!.name),
      'queueManagerIdText': queueManagerIdController.text,
      'profilePicture': await MultipartFile.fromFile(profilePicture!.path!,
          filename: profilePicture!.name),
      'queueManagerIdFile': queueManagerIdController.text,
    });

    try {
      final response = await dioReg.post('${dioBase}register', data: formData);
      if (response.statusCode == 200) {
        print('Upload successful');
      } else {
        print('Upload failed: ${response.statusCode}');
      }
    } catch (e) {
      print('Upload failed: $e');
    }
  }

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
                    controller: nameController,
                    hintText: 'Name'),
                const SizedBox(
                  height: 8,
                ),
                MyTextField(
                    labelText: 'Last Name',
                    controller: nameController,
                    hintText: 'Name'),
                const SizedBox(
                  height: 8,
                ),
                MyTextField(
                    labelText: 'Email',
                    controller: emailController,
                    hintText: 'text@text.com'),
                const SizedBox(
                  height: 12,
                ),
                MyTextField(
                    labelText: 'Phone Number',
                    controller: phoneController,
                    hintText: '0912345678'),
                const SizedBox(
                  height: 12,
                ),
                MyTextField(
                    labelText: 'Station',
                    controller: stationController,
                    hintText: 'St.George'),
                const SizedBox(
                  height: 8,
                ),
                MyTextField(
                    labelText: 'Queue Manager ID',
                    controller: queueManagerIdController,
                    hintText: '123456'),
                const Text(
                  'Upload ID',
                  textAlign: TextAlign.left,
                  style: TextStyle(
                    fontSize: 40,
                    fontWeight: FontWeight.w800,
                  ),
                ),
                const FileSelectorWidget(label: 'Upload Kebele ID'),
                const FileSelectorWidget(label: 'Upload Profile Picture'),
                const SizedBox(
                  height: 8,
                ),
                AuthButton(
                    label: 'Register',
                    onPressed: () {
                      AppRouter.router.go('/');
                    }),
                const SizedBox(
                  height: 8,
                ),
              ],
            )),
          ),
        ));
  }
}
