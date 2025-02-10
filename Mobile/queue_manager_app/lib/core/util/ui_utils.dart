import 'package:flutter/material.dart';

import '../theme/app_colors.dart';

class UiUtils {
  static GlobalKey<ScaffoldMessengerState> scaffoldMessengerKey =
      GlobalKey<ScaffoldMessengerState>();

  static Future<dynamic> showBottomSheet({
    required Widget child,
    required BuildContext context,
    bool? enableDrag,
  }) async {
    final result = await showModalBottomSheet(
      enableDrag: enableDrag ?? false,
      backgroundColor: Colors.transparent,
      isScrollControlled: true,
      shape: const RoundedRectangleBorder(
        borderRadius: BorderRadius.only(
          topLeft: Radius.circular(10),
          topRight: Radius.circular(10),
        ),
      ),
      context: context,
      builder: (_) => child,
    );

    return result;
  }

  static void showSnackBar({
    required String message,
    String? label,
    VoidCallback? onPressed,
    bool isError = false,
    Color? color,
  }) {
    final snackBar = SnackBar(
      margin: EdgeInsets.symmetric(horizontal: 20, vertical: 20),
      dismissDirection: DismissDirection.horizontal,
      backgroundColor:
          color ?? (isError ? AppColors.errorColor : AppColors.successColor),
      showCloseIcon: label != null && onPressed != null ? true : false,
      action: label != null && onPressed != null
          ? SnackBarAction(
              label: label,
              textColor: Colors.white,
              onPressed: onPressed,
            )
          : null,
      behavior: SnackBarBehavior.floating,
      content: Text(message, style: const TextStyle(color: Colors.white)),
    );

    scaffoldMessengerKey.currentState!
      ..removeCurrentSnackBar()
      ..showSnackBar(snackBar);
  }

  static void showAlertDialog(
      BuildContext context, String title, String message,
      {String? positiveLabel,
      String? negativeLabel,
      VoidCallback? onPositivePressed,
      VoidCallback? onNegativePressed}) {
    showDialog(
      context: context,
      builder: (context) => AlertDialog(
        title: Text(title),
        content: Text(message),
        actions: <Widget>[
          if (negativeLabel != null)
            TextButton(
              onPressed: onNegativePressed,
              child: Text(negativeLabel),
            ),
          if (positiveLabel != null)
            TextButton(
              onPressed: onPositivePressed,
              child: Text(positiveLabel),
            ),
        ],
      ),
    );
  }

  static void showOverlay(BuildContext context, String message, Color color) {
    final overlay = Overlay.of(context);
    final size = MediaQuery.sizeOf(context);
    OverlayEntry? overlayEntry;

    void dismissOverlay() {
      if (overlayEntry != null) {
        overlayEntry!.remove();
        overlayEntry = null;
      }
    }

    overlayEntry = OverlayEntry(
      builder: (context) => Positioned(
        bottom: size.height * 0.05,
        left: size.width * 0.1,
        right: size.width * 0.1,
        child: Material(
          color: Colors.transparent,
          child: GestureDetector(
            onHorizontalDragEnd: (details) {
              // Dismiss when swipe ends
              dismissOverlay();
            },
            child: Container(
              padding: const EdgeInsets.all(16.0),
              decoration: BoxDecoration(
                color: color,
                borderRadius: BorderRadius.circular(8.0),
                boxShadow: const [
                  BoxShadow(
                    color: Colors.black26,
                    blurRadius: 10.0,
                    offset: Offset(0, 10),
                  ),
                ],
              ),
              child: Text(
                message,
                style: const TextStyle(color: Colors.white),
              ),
            ),
          ),
        ),
      ),
    );

    overlay.insert(overlayEntry!);

    Future.delayed(const Duration(seconds: 3), () {
      if (overlayEntry != null) {
        overlayEntry!.remove();
      }
    });
  }
}
