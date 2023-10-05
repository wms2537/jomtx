import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:intl_phone_number_input/intl_phone_number_input.dart';
import 'package:jomtx/env.dart';
import 'package:jomtx/router/route_path.dart';
import 'package:http/http.dart' as http;
import 'package:jomtx/widgets/error_dialog.dart';

final String baseUrl = environment['apiUrl']!;

class ClaimTxnPath extends AppRoutePath {
  final String txnId;
  const ClaimTxnPath(this.txnId);
  @override
  String getRouteInformation() {
    return '/claim-txn/$txnId';
  }
}

class ClaimTxnScreen extends StatefulWidget {
  final String txnId;
  final Function(String txnId) goToReceipt;
  const ClaimTxnScreen(
      {super.key, required this.txnId, required this.goToReceipt});

  @override
  State<ClaimTxnScreen> createState() => _ClaimTxnScreenState();
}

class _ClaimTxnScreenState extends State<ClaimTxnScreen> {
  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();

  final TextEditingController controller = TextEditingController();
  String initialCountry = 'MY';
  PhoneNumber number = PhoneNumber(isoCode: 'MY');
  String? _currentNumber;
  bool _isLoading = false, _isSuccess = false;

  @override
  Widget build(BuildContext context) {
    final deviceSize = MediaQuery.of(context).size;
    return Scaffold(
      body: Center(
        child: Card(
          elevation: 4.0,
          child: Container(
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(5.0),
            ),
            width: deviceSize.aspectRatio > 1
                ? deviceSize.width * 0.6
                : deviceSize.width * 0.9,
            margin: const EdgeInsets.all(16),
            child: Form(
              key: _formKey,
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  const Padding(
                    padding: EdgeInsets.all(32.0),
                    child: Text(
                      "Claim Txn",
                      style: TextStyle(
                        fontSize: 20,
                        fontWeight: FontWeight.w600,
                      ),
                    ),
                  ),
                  Padding(
                    padding: const EdgeInsets.all(16.0),
                    child: _isSuccess
                        ? const Text(
                            "You have successfully claimed your Receipt",
                            style: TextStyle(fontSize: 14.5),
                            textAlign: TextAlign.center,
                          )
                        : InternationalPhoneNumberInput(
                            onInputChanged: (PhoneNumber number) {
                              _currentNumber = number.phoneNumber;
                            },
                            selectorConfig: const SelectorConfig(
                              selectorType: PhoneInputSelectorType.DIALOG,
                            ),
                            ignoreBlank: false,
                            autoValidateMode:
                                AutovalidateMode.onUserInteraction,
                            selectorTextStyle:
                                const TextStyle(color: Colors.black),
                            initialValue: number,
                            textFieldController: controller,
                            formatInput: true,
                            keyboardType: const TextInputType.numberWithOptions(
                                signed: true, decimal: true),
                            inputBorder: const OutlineInputBorder(),
                            onSaved: (PhoneNumber number) {
                              _currentNumber = number.phoneNumber;
                            },
                          ),
                  ),
                  Padding(
                    padding: const EdgeInsets.all(16.0),
                    child: _isLoading
                        ? const CircularProgressIndicator()
                        : _isSuccess
                            ? ElevatedButton(
                                onPressed: () {
                                  widget.goToReceipt(widget.txnId);
                                },
                                child: const Text('View Receipt'))
                            : ElevatedButton(
                                onPressed: () async {
                                  if (_formKey.currentState?.validate() ==
                                      true) {
                                    setState(() {
                                      _isLoading = true;
                                    });
                                    try {
                                      final Map<String, dynamic> data = {
                                        'txnId': widget.txnId,
                                        'phoneNumber': _currentNumber,
                                        'captcha': "",
                                      };
                                      await http.post(
                                        Uri.parse("$baseUrl/claimTxn"),
                                        headers: {
                                          'Content-Type':
                                              'application/json', // Specify the content type
                                        },
                                        body: json.encode(data),
                                      );
                                      setState(() {
                                        _isSuccess = true;
                                      });
                                    } catch (e) {
                                      if (!mounted) return;
                                      showErrorDialog(e.toString(), context);
                                    }

                                    setState(() {
                                      _isLoading = false;
                                    });
                                  }
                                },
                                child: const Text('Claim'),
                              ),
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
