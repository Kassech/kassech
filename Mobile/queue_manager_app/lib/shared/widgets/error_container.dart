import 'package:flutter/material.dart';
import 'package:flutter_animate/flutter_animate.dart';
import 'package:flutter_svg/flutter_svg.dart';

import '../../core/util/helper.dart';

class ErrorContainer extends StatelessWidget {
  final String? errorMessageCode;
  final String? errorMessageText;
  final String? buttonText;
  final bool? showRetryButton;
  final bool? showErrorImage;
  final Color? errorMessageColor;
  final double? errorMessageFontSize;
  final Function? onTapRetry;
  final Color? retryButtonBackgroundColor;
  final Color? retryButtonTextColor;

  const ErrorContainer({
    super.key,
    this.errorMessageCode,
    this.errorMessageText,
    this.errorMessageColor,
    this.errorMessageFontSize,
    this.onTapRetry,
    this.showErrorImage,
    this.retryButtonBackgroundColor,
    this.retryButtonTextColor,
    this.showRetryButton,
    this.buttonText,
  });

  @override
  Widget build(BuildContext context) {
    final size = MediaQuery.sizeOf(context);
    return Animate(
      effects: const [
        ScaleEffect(
          duration: Duration(
            milliseconds: 200,
          ),
          curve: Curves.bounceOut,
        ),
      ],
      child: SingleChildScrollView(
        padding: const EdgeInsets.all(20),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            SizedBox(
              child: SvgPicture.asset(
                !errorMessageText.toString().toLowerCase().contains('no internet')
                    ? Helper.getImagePath('no_internet.svg')
                    : Helper.getImagePath('something_went_wrong.svg'),
                height: size.height * 0.5,
                fit: BoxFit.fitWidth,
              ),
            ),
            const SizedBox(
              height: 15,
            ),
            Padding(
              padding: const EdgeInsets.symmetric(horizontal: 8),
              child: Text(
                errorMessageText ?? 'Something went wrong, Please try again!',
                textAlign: TextAlign.center,
                style: TextStyle(
                  color: errorMessageColor
                ),
              ),
            ),
            const SizedBox(
              height: 15,
            ),
            (showRetryButton ?? true)
                ? ElevatedButton(
                  onPressed: () => onTapRetry!(),
                  child: Text(
                    buttonText ?? 'Retry',
                    style: TextStyle(
                      color: retryButtonTextColor,
                    ),
                  ),
                )
                : const SizedBox()
          ],
        ),
      ),
    );
  }
}
