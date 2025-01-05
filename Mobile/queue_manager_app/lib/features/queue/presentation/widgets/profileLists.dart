import 'package:flutter/material.dart';

class ProfileField extends StatelessWidget {
  final String label;
  final String value;

  const ProfileField({
    super.key,
    required this.label,
    required this.value,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 8.0),
      child: Container(
        width: double.infinity,
        decoration: BoxDecoration(
          color: Colors.white,
          borderRadius: BorderRadius.circular(8),
          border: Border.all(color: Colors.black),
        ),
        padding: const EdgeInsets.symmetric(horizontal: 36, vertical: 8.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              label,
              style: const TextStyle(
                fontWeight: FontWeight.bold,
                fontSize: 16,
              ),
            ),
            const SizedBox(height: 4),
            Text(
              value.isNotEmpty ? value : 'N/A',
              style: TextStyle(
                fontSize: 16,
                color: value.isEmpty ? Colors.grey : Colors.black,
              ),
            ),
          ],
        ),
      ),
    );
  }
}
