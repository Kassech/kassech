import 'package:dio/dio.dart';
import 'package:file_picker/file_picker.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:queue_manager_app/features/auth/domain/usecase/api_service.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/mytextfield.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/filePicker.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/authButton.dart';
import 'package:queue_manager_app/config/route/route.dart';
import 'package:queue_manager_app/features/auth/presentation/widgets/password.dart';
import 'package:queue_manager_app/features/queue/presentation/provider/filePickerNotifier.dart';
import 'package:queue_manager_app/features/queue/presentation/provider/profileChangeProvider.dart';

class SignUpPage extends StatefulWidget {
  final int role;
  

  SignUpPage({required this.role, super.key});

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
  final TextEditingController profilePictureController = TextEditingController();

  final dioInst = ApiService().dio_instance;
  final formKey = GlobalKey<FormState>();

  Future<void> pickFile(String fileType, WidgetRef ref) async {
    final result = await FilePicker.platform.pickFiles(type: FileType.any);
    if (result != null) {
      ref
          .read(filePickerProvider.notifier)
          .setFile(fileType, result.files.first);
    }
  }

  Future<void> uploadFiles(Map<String, dynamic>? files, WidgetRef ref) async {
    print('uploadFiles called with files: $files');

    if (files == null || files.isEmpty) {
      print('No files provided, aborting upload');
      return;
    }

    print('Starting file upload process');

    final kebeleId = files['kebeleId'];
    print('Retrieved kebeleId file: $kebeleId');

    final profilePicture = files['profilePicture'];
    print('Retrieved profilePicture file: $profilePicture');

    final drivingLicenseFile = files['drivingLicense'];
    print('Retrieved drivingLicenseFile file: $drivingLicenseFile');

    final insuranceDocumentFile = files['insuranceDocument'];
    print('Retrieved insuranceDocumentFile file: $insuranceDocumentFile');

    if (kebeleId?.path == null || profilePicture?.path == null) {
      print('Required files are missing, aborting upload');
      return;
    }

    if (widget.role == 2 || widget.role == 4) {
      if (drivingLicenseFile?.path == null) {
        print('Driving License file is missing for role ${widget.role}');
        return;
      }
    }

    if (widget.role == 4) {
      if (insuranceDocumentFile?.path == null) {
        print('Insurance Document file is missing for role ${widget.role}');
        return;
      }
    }

    print('All required files are present, preparing form data');

    try {
      final formData = FormData.fromMap({
        'FirstName': nameController.text,
        'LastName': fathersNameController.text,
        'Email': emailController.text,
        'Password': passwordController.text,
        'PhoneNumber': phoneController.text,
        'Role': widget.role,
        'NationalIdFile': await MultipartFile.fromFile(kebeleId.path, filename: kebeleId.name),

        'Profile': await MultipartFile.fromFile(
          profilePicture.path,
          filename: profilePicture.name,
        ),
        'DrivingLicenseFile': widget.role == 2 || widget.role == 4
            ? await MultipartFile.fromFile(
                drivingLicenseFile.path,
                filename: drivingLicenseFile.name,
              )
            : null,
        'InsuranceDocumentFile': widget.role == 4
            ? await MultipartFile.fromFile(
                insuranceDocumentFile.path,
                filename: insuranceDocumentFile.name,
              )
            : null,
      });

      print('Form data prepared, sending to server');

      final response = await dioInst.post(
        '${ApiService().dio_baseUrl}register',
        data: formData,
        options: Options(
          headers: {
            'Content-Type': 'multipart/form-data',
          },
        ),
      );

      if (response.statusCode == 200) {
        final profilePictureUrl = response.data['profilePictureUrl'];
        ref
            .read(profilePictureProvider.notifier)
            .setProfilePicture(profilePictureUrl);
        print('Upload successful');
      } else {
        print('Upload failed: ${response.statusCode}');
      }

    }on DioException catch (e) {
      print('Upload failed from dio: $e');
    }
 
     catch (e) {
      print('Upload failed: $e');

    }
  }

  Future<void> _signup(BuildContext context, WidgetRef ref) async {
    final files = ref.read(filePickerProvider);
    await uploadFiles(files, ref);

    // Navigate to the register page if the upload was successful
    AppRouter.router.go('/register');
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        backgroundColor:  Colors.white,
        leading: IconButton(
          icon: const Icon(Icons.arrow_back),
          onPressed: () {
            AppRouter.router.go( '/signin');
          },
        ),
      ),
      body: Form(
        key: formKey, // Set the form key.
        child: Padding(
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
                          fontWeight: FontWeight.bold,
                        ),
                      ),
                                          const SizedBox(
                        height: 20,
                      ),
        
                      MyTextField(
                        labelText: 'First Name',
                        controller: nameController,
                        hintText: 'First Name', validator: (val) { return val.isEmpty ? 'Enter your first name' : null; },
                      ),
                                          const SizedBox(
                        height: 20,
                      ),
        
                      MyTextField(
                        labelText: 'Father\'s Name',
                        controller: fathersNameController,
                        hintText: 'Father\'s Name', validator: (val) { return val.isEmpty ? 'Enter your father\'s name' : null; },
                      ),
                                          const SizedBox(
                        height: 20,
                      ),
        
                      MyTextField(
                        labelText: 'Email',
                        controller: emailController,
                        hintText: 'Email', validator: (val) { return val.isEmpty ? 'Enter your email' : null; },
                      ),
                                          const SizedBox(
                        height: 20,
                      ),
        
                      MyPasswordTextField(
                        labelText: 'Password',
                        controller: passwordController,
                        hintText: 'Password', validator: (val) { return val.isEmpty ? 'Enter your password' : null; },
                      ),
                      const SizedBox(
                        height: 20,
                      ),
                      MyTextField(
                        labelText: 'Phone Number',
                        controller: phoneController,
                        hintText: 'Phone Number', validator: (val) { return val.isEmpty ? 'Enter your phone number' : null; },
                      ),
                      const SizedBox(height: 20,),
                    
                      
                      
                      FileSelectorWidget(
                        label: 'Kebele ID File',
                        onFileSelected: (filePath) => ref
                            .read(filePickerProvider.notifier)
                            .setFile('kebeleId', PlatformFile(path: filePath, name: '', size: 0)), onPressed: (String? filePath, WidgetRef ref) {  },
                      ),
                      FileSelectorWidget(
                        label: 'Profile Picture',
                        onFileSelected: (filePath) => ref
                            .read(filePickerProvider.notifier)
                            .setFile(
                                'profilePicture', PlatformFile(path: filePath, name: '', size: 0)), onPressed: (String? filePath, WidgetRef ref) {  },
                      ),
                     
                      if (widget.role == 2) ...[
                        FileSelectorWidget(
                          label: 'Driving License File',
                          onFileSelected: (filePath) => ref
                              .read(filePickerProvider.notifier)
                              .setFile(
                                  'drivingLicense', PlatformFile(path: filePath, name: '', size: 0)), onPressed: (String? filePath, WidgetRef ref) {  },
                        ),
                        if (widget.role == 4) ...[
                          FileSelectorWidget(
                            label: 'Insurance Document File',
                            onFileSelected: (filePath) => ref
                                .read(filePickerProvider.notifier)
                                .setFile('insuranceDocument',
                                    PlatformFile(path: filePath, name: '', size: 0)), onPressed: (String? filePath, WidgetRef ref) {  },
                          ),
                        ]
                      ],
                      AuthButton(
                        label: 'SignUp',
                        onPressed: () async {
                          if (!formKey.currentState!.validate()) {
                            return;
                          }

                          await uploadFiles(files, ref);
                        },
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
