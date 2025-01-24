import 'package:flutter/material.dart';
import 'package:queue_manager_app/config/route/route.dart';

class ErrorPage extends StatelessWidget {
  const ErrorPage({super.key});

  @override
  Widget build(BuildContext context) {
    return  Scaffold(
      appBar: AppBar(
        leading: IconButton(
          icon: const Icon(Icons.arrow_back),
          onPressed: () { AppRouter.router.go('/signin'); },
        ),
       ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(
              'Something went wrong',
              style: TextStyle(fontSize: 40, fontWeight: FontWeight.w700),
            )
          ],
        ),
      ),
    );
  }
}
