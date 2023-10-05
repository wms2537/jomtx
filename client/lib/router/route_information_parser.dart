import 'package:jomtx/router/route_path.dart';
import 'package:jomtx/screens/claim_txn_screen.dart.dart';
import 'package:jomtx/screens/page_not_found_screen.dart';
import 'package:flutter/material.dart';
import 'package:jomtx/screens/receipt_screen.dart';

class AppRouteInformationParser extends RouteInformationParser<AppRoutePath> {
  @override
  Future<AppRoutePath> parseRouteInformation(
      RouteInformation routeInformation) async {
    final uri = routeInformation.uri;
    if (uri.pathSegments.length == 2) {
      final first = uri.pathSegments[0].toLowerCase();
      // final second = uri.pathSegments[1].toLowerCase();
      if (first == 'claim-txn') {
        return ClaimTxnPath(uri.pathSegments[1]);
      } else if (first == 'receipt') {
        return ReceiptPath(uri.pathSegments[1]);
      }
    }
    return PageNotFoundPath();
  }

  @override
  RouteInformation restoreRouteInformation(AppRoutePath configuration) {
    return RouteInformation(
        uri: Uri.parse(configuration.getRouteInformation()));
  }
}
