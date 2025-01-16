import 'package:file_picker/file_picker.dart';
import 'package:flutter/material.dart';

class FileSelectorWidget extends StatefulWidget {
  final String label; // The label to differentiate the file selector

  const FileSelectorWidget({super.key, required this.label, required Future<void> Function() onPressed});

  @override
  _FileSelectorWidgetState createState() => _FileSelectorWidgetState();
}

class _FileSelectorWidgetState extends State<FileSelectorWidget> {
  String? _selectedFilePath;

  // Method to open the file picker
  Future<void> _pickFile() async {
    // Open the file picker
    FilePickerResult? result = await FilePicker.platform.pickFiles();

    // Check if a file was selected
    if (result != null) {
      setState(() {
        _selectedFilePath = result.files.single.path;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(18.0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            widget.label, // Display the label for the file selector
            style:const TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
          ),
          const SizedBox(height: 8),
          GestureDetector(
            onTap: _pickFile, // Open file picker when tapped
            child: Container(
              padding: const EdgeInsets.symmetric(vertical: 18, horizontal: 22),
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
                        : _selectedFilePath != null
                            ? _selectedFilePath!
                                .split('/')
                                .last
                                .length > 10
                                ? '${_selectedFilePath!.split('/').last.substring(0, 10)}...'
                                : _selectedFilePath!
                                    .split('/')
                                    .last // Show file name if selected
                            : 'Select a file',
                    style: const TextStyle(fontSize: 16, color: Colors.black),
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
                'Selected File: ${_selectedFilePath!.split('/').last}', // Display selected file name
                style: const TextStyle(fontSize: 12, color: Colors.green),
              ),
            ),
        ],
      ),
    );
  }
}
