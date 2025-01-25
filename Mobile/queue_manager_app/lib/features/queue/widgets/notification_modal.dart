
import 'package:flutter/material.dart';

class NotificationModal extends StatelessWidget {
  final String title;
  final String body;

  const NotificationModal({
    Key? key,
    required this.title,
    required this.body,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: Text(title),
      content: Text(body),
      actions: [
        TextButton(
          onPressed: () {
            Navigator.of(context).pop();
          },
          child: Text('Arrived'),
        ),
      ],
    );
  }
}
