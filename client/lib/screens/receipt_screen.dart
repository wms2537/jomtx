import 'dart:convert';

import 'package:fixnum/fixnum.dart';
import 'package:flutter/material.dart';
import 'package:hex/hex.dart';
import 'package:jomtx/env.dart';
import 'package:jomtx/proto/proto/jomtx.pb.dart';
import 'package:jomtx/router/route_path.dart';
import 'package:http/http.dart' as http;
import 'package:jomtx/widgets/error_dialog.dart';

final String rpcUrl = environment['rpcUrl']!;

class ReceiptPath extends AppRoutePath {
  final String txnId;
  const ReceiptPath(this.txnId);
  @override
  String getRouteInformation() {
    return '/receipt/$txnId';
  }
}

class ReceiptScreen extends StatelessWidget {
  final String txnId;
  const ReceiptScreen({super.key, required this.txnId});

  Future<Txn?> _getData() async {
    try {
      final queryData = QueryGetTxnRequest(id: Int64(int.parse(txnId)));
      final query = {
        "jsonrpc": "2.0",
        "id": 1,
        "method": "abci_query",
        "params": {
          "path": "/jomtx.jomtx.Query/Txn", // Replace with your query path
          "data": HEX.encode(
              queryData.writeToBuffer()), // Replace with query data if needed
        }
      };

      final response = await http.post(
        Uri.parse(rpcUrl),
        headers: {
          'Content-Type': 'application/json',
        },
        body: jsonEncode(query),
      );

      if (response.statusCode == 200) {
        final responseBody = json.decode(response.body);
        final result = responseBody['result']['response']['value'];
        final decodedValue =
            QueryGetTxnResponse.fromBuffer(base64Decode(result));
        return decodedValue.txn;
      }
      return null;
    } catch (e) {
      print(e);
      rethrow;
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: FutureBuilder(
            future: _getData(),
            builder: (context, snapshot) {
              return snapshot.connectionState == ConnectionState.waiting
                  ? const CircularProgressIndicator()
                  : snapshot.hasError
                      ? const Text("Failed to fetch Txn")
                      : Column(
                          children: [
                            const Text("My Receipt"),
                            Text("Id: $txnId"),
                            Text("Invoice No: ${snapshot.data!.invoiceNo}"),
                            Text("Items: ${snapshot.data!.items}"),
                            Text("Total: ${snapshot.data!.total}"),
                          ],
                        );
            }),
      ),
    );
  }
}
