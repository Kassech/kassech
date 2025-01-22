import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:file_picker/file_picker.dart';

final filePickerProvider =
    StateNotifierProvider<FilePickerNotifier, Map<String, PlatformFile?>>(
        (ref) {
  return FilePickerNotifier();
});

class FilePickerNotifier extends StateNotifier<Map<String, PlatformFile?>> {
  FilePickerNotifier() : super({});

  void setFile(String fileType, PlatformFile file) {
    state = {...state, fileType: file};
  }

  void removeFile(String fileType) {
    state = {...state}..remove(fileType);
  }
}
