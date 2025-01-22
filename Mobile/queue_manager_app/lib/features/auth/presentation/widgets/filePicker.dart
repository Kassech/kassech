import 'package:file_picker/file_picker.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class FileSelectorWidget extends StatefulWidget {
  final String label; // The label to differentiate the file selector
  final void Function(String? filePath, WidgetRef ref)
      onPressed; // The callback function

  const FileSelectorWidget({
    Key? key,
    required this.label,
    required this.onPressed, required void Function(dynamic filePath) onFileSelected,
  }) : super(key: key);

  @override
  _FileSelectorWidgetState createState() => _FileSelectorWidgetState();
}

class _FileSelectorWidgetState extends State<FileSelectorWidget> {
  String? _selectedFilePath;

  // Setter method to update the file path
  void setFilePath(String? filePath) {
    setState(() {
      _selectedFilePath = filePath;
    });
  }

  // Method to open the file picker
  Future<void> _pickFile(WidgetRef ref) async {
    FilePickerResult? result;
    try {
      result = await FilePicker.platform.pickFiles();
    } catch (e) {
      print('Error picking file: $e');
      return;
    }

    if (result != null) {
      setState(() {
        _selectedFilePath = result?.files.single.path;
        
      });
      widget.onPressed(_selectedFilePath, ref); // Call the callback function
    }
  }

  @override
  Widget build(BuildContext context) {
    return Consumer(
      builder: (context, ref, child) {
        return Padding(
          padding: const EdgeInsets.all(18.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                widget.label,
                style:
                    const TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
              ),
              const SizedBox(height: 8),
              GestureDetector(
                onTap: () => _pickFile(ref), // Pass WidgetRef to _pickFile
                child: Container(
                  padding:
                      const EdgeInsets.symmetric(vertical: 18, horizontal: 22),
                  decoration: BoxDecoration(
                    border: Border.all(color: Colors.black87),
                    borderRadius: BorderRadius.circular(8),
                  ),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Text(
                        _selectedFilePath == null
                            ? 'Select a file'
                            : _selectedFilePath!.split('/').last.length > 10
                                ? '${_selectedFilePath!.split('/').last.substring(0, 10)}...'
                                : _selectedFilePath!.split('/').last,
                        style:
                            const TextStyle(fontSize: 16, color: Colors.black),
                      ),
                      const Icon(Icons.attach_file, color: Colors.grey),
                    ],
                  ),
                ),
              ),
              if (_selectedFilePath != null)
                Padding(
                  padding: const EdgeInsets.symmetric(vertical: 25),
                  child: Text(
                    'Selected File: ${_selectedFilePath!.split('/').last}',
                    style: const TextStyle(fontSize: 12, color: Colors.green),
                  ),
                ),
            ],
          ),
        );
      },
    );
  }
}
