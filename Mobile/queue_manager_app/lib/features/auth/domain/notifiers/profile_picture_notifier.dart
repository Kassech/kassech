import 'package:flutter_riverpod/flutter_riverpod.dart';

final profilePictureProvider =
    StateNotifierProvider<ProfilePictureNotifier, String?>((ref) {
  return ProfilePictureNotifier();
});

class ProfilePictureNotifier extends StateNotifier<String?> {
  ProfilePictureNotifier() : super(null);

  void setProfilePicture(String? url) {
    state = url;
  }

  void clearProfilePicture() {
    state = null;
  }
}
