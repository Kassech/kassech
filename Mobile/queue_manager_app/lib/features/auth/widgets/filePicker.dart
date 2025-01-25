import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../../../core/theme/app_colors.dart';

class FileSelectorWidget extends StatelessWidget {
  final String label;
  final String? filePath;
  final VoidCallback pickFile;

  const FileSelectorWidget({
    super.key,
    required this.label,
    required this.pickFile, required this.filePath,
  });

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
                label,
                style:
                    const TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
              ),
              const SizedBox(height: 8),
              GestureDetector(
                onTap: () => pickFile(),
                child: Container(
                  padding:
                      const EdgeInsets.symmetric(vertical: 18, horizontal: 22),
                  decoration: BoxDecoration(
                    border: Border.all(
                      color: Theme.of(context).colorScheme.onSecondary,
                    ),
                    borderRadius: BorderRadius.circular(8),
                  ),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Text(
                        filePath == null
                            ? 'Select a file'
                            : filePath!.split('/').last.length > 10
                                ? '${filePath!.split('/').last.substring(0, 10)}...'
                                : filePath!.split('/').last,
                        style:
                            const TextStyle(fontSize: 16,),
                      ),
                      const Icon(Icons.attach_file),
                    ],
                  ),
                ),
              ),
              if (filePath != null)
                Padding(
                  padding: const EdgeInsets.symmetric(vertical: 25),
                  child: Text(
                    'Selected File: ${filePath!.split('/').last}',
                    style: TextStyle(fontSize: 12, color: AppColors.successColor),
                  ),
                ),
            ],
          ),
        );
      },
    );
  }
}
