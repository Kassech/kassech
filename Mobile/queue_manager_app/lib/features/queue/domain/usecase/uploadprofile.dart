
import 'dart:io';

import 'package:image_picker/image_picker.dart';

File? _image;


pickImage() async {
  final pickedFile = await ImagePicker().pickImage(source: ImageSource.gallery);
  if (pickedFile != null) {
    _image = File(pickedFile.path);
    // Since this is a standalone function, `setState` cannot be used here.
    // Consider using a callback or state management solution to update the UI.
  }
  
}

