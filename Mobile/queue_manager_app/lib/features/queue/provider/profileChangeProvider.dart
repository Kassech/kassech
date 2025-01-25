import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:image_picker/image_picker.dart';
import 'dart:io';

final profilePictureProvider =
    StateNotifierProvider<ProfilePictureNotifier, File?>((ref) {
  return ProfilePictureNotifier();
});

class ProfilePictureNotifier extends StateNotifier<File?> {
  ProfilePictureNotifier() : super(null);

  void setProfilePicture(String? profilePictureUrl) {
    state = profilePictureUrl != null ? File(profilePictureUrl) : null;
  }

  Future<void> updateProfilePicture() async {
    final picker = ImagePicker();
    final pickedFile = await picker.pickImage(source: ImageSource.gallery);

    if (pickedFile != null) {
      state = File(pickedFile.path);
    }
  }
}
