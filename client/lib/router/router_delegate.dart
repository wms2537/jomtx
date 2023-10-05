import 'package:jomtx/router/app_state.dart';
import 'package:jomtx/router/fade_animation.dart';
import 'package:jomtx/router/route_path.dart';
import 'package:flutter/material.dart';
import 'package:jomtx/screens/claim_txn_screen.dart.dart';
import 'package:jomtx/screens/page_not_found_screen.dart';
import 'package:jomtx/screens/receipt_screen.dart';

class AppRouterDelegate extends RouterDelegate<AppRoutePath>
    with ChangeNotifier, PopNavigatorRouterDelegateMixin<AppRoutePath> {
  @override
  final GlobalKey<NavigatorState> navigatorKey;
  bool? _showPageNotFound;
  // String? _emailVerificationToken;

  final AppState appState = AppState();

  AppRouterDelegate() : navigatorKey = GlobalKey<NavigatorState>() {
    appState.addListener(notifyListeners);
  }

  @override
  AppRoutePath? get currentConfiguration {
    if (appState.txnId != null) {
      if (appState.isShowReceipt == true) {
        return ReceiptPath(appState.txnId!);
      }
      return ClaimTxnPath(appState.txnId!);
    }
    if (_showPageNotFound == true) {
      return PageNotFoundPath();
    }
    return null;
  }

  @override
  Widget build(BuildContext context) {
    List<Page> stack = [];
    if (appState.txnId != null) {
      if (appState.isShowReceipt == true) {
        stack = [
          FadeAnimationPage(
            child: ReceiptScreen(
              txnId: appState.txnId!,
            ),
            key: const ValueKey('ReceiptPage'),
          ),
        ];
      } else {
        stack = [
          FadeAnimationPage(
            child: ClaimTxnScreen(
              txnId: appState.txnId!,
              goToReceipt: (txnId) {
                appState.isShowReceipt = true;
              },
            ),
            key: const ValueKey('ClaimTxnPage'),
          ),
        ];
      }
    } else {
      stack = [
        const FadeAnimationPage(
          key: ValueKey('PageNotFoundPage'),
          child: PageNotFoundScreen(),
        )
      ];
    }
    return Navigator(
      key: navigatorKey,
      pages: stack,
      onPopPage: (route, result) {
        if (!route.didPop(result)) return false;
        return true;
      },
    );
  }

  @override
  Future<void> setNewRoutePath(AppRoutePath configuration) async {
    if (configuration is ClaimTxnPath) {
      appState.isShowReceipt = false;
      appState.txnId = configuration.txnId;
    } else if (configuration is ReceiptPath) {
      appState.isShowReceipt = true;
      appState.txnId = configuration.txnId;
    } else {
      appState.txnId = null;
    }
    if (configuration is PageNotFoundPath) {
      _showPageNotFound = true;
      return;
    } else {
      _showPageNotFound = false;
    }
  }
}
