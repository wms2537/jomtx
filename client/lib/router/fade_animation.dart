import 'package:flutter/material.dart';

class FadeAnimationPage extends Page {
  final Widget child;

  const FadeAnimationPage({LocalKey? key, required this.child})
      : super(key: key);

  @override
  Route createRoute(BuildContext context) {
    return PageRouteBuilder(
      settings: this,
      pageBuilder: (context, animation, animation2) {
        var curveTween = CurveTween(curve: Curves.easeIn);
        return FadeTransition(
          opacity: animation.drive(curveTween),
          child: child,
        );
      },
    );
  }
}
