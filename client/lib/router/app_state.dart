import 'package:flutter/foundation.dart';

class AppState with ChangeNotifier {
  int _bottomNavigatorIndex = 1; // 0: Orders, 1: Home, 2: Account
  int get bottomNavigatorIndex => _bottomNavigatorIndex;
  set bottomNavigatorIndex(int value) {
    _bottomNavigatorIndex = value;
    notifyListeners();
  }

  String? _txnId;
  String? get txnId => _txnId;
  set txnId(String? id) {
    _txnId = id;
    notifyListeners();
  }

  bool? _isShowReceipt;
  bool? get isShowReceipt => _isShowReceipt;
  set isShowReceipt(bool? value) {
    _isShowReceipt = value;
    notifyListeners();
  }
}
