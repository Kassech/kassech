import 'package:file_picker/file_picker.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../models/user_params.dart';

final userDataProvider = NotifierProvider<UserDataNotifier, UserParams?>(() {
  return UserDataNotifier();
});

class UserDataNotifier extends Notifier<UserParams?> {
  @override
  UserParams? build() {
    return UserParams();
  }

  void updateUserData(UserParams user) {
    state = user;
  }

  void updateUserFiles({
    PlatformFile? kebeleId,
    PlatformFile? profilePicture,
    PlatformFile? drivingLicenseFile,
    PlatformFile? insuranceDocumentFile,
  }) {
    state = state?.copyWith(
      kebeleId: kebeleId,
      profilePicture: profilePicture,
      drivingLicenseFile: drivingLicenseFile,
      insuranceDocumentFile: insuranceDocumentFile,
    );
  }

  void clearUserData() {
    state = null;
  }
}
