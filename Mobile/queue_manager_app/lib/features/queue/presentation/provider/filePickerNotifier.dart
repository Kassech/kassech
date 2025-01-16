import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:file_picker/file_picker.dart';

class FilePickerNotifier extends StateNotifier<Map<String, PlatformFile?>> {
  FilePickerNotifier() : super({});

  void setFile(String fileType, PlatformFile? file) {
    state = {...state, fileType: file};
  }
}

final filePickerProvider =
    StateNotifierProvider<FilePickerNotifier, Map<String, PlatformFile?>>(
        (ref) {
  return FilePickerNotifier();
});
