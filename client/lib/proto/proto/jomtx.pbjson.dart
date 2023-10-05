//
//  Generated code. Do not modify.
//  source: proto/jomtx.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use queryGetTxnRequestDescriptor instead')
const QueryGetTxnRequest$json = {
  '1': 'QueryGetTxnRequest',
  '2': [
    {'1': 'id', '3': 1, '4': 2, '5': 4, '10': 'id'},
  ],
};

/// Descriptor for `QueryGetTxnRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List queryGetTxnRequestDescriptor = $convert.base64Decode(
    'ChJRdWVyeUdldFR4blJlcXVlc3QSDgoCaWQYASACKARSAmlk');

@$core.Deprecated('Use queryGetTxnResponseDescriptor instead')
const QueryGetTxnResponse$json = {
  '1': 'QueryGetTxnResponse',
  '2': [
    {'1': 'Txn', '3': 1, '4': 2, '5': 11, '6': '.Txn', '10': 'Txn'},
  ],
};

/// Descriptor for `QueryGetTxnResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List queryGetTxnResponseDescriptor = $convert.base64Decode(
    'ChNRdWVyeUdldFR4blJlc3BvbnNlEhYKA1R4bhgBIAIoCzIELlR4blIDVHhu');

@$core.Deprecated('Use txnDescriptor instead')
const Txn$json = {
  '1': 'Txn',
  '2': [
    {'1': 'id', '3': 1, '4': 2, '5': 4, '10': 'id'},
    {'1': 'invoiceNo', '3': 2, '4': 2, '5': 9, '10': 'invoiceNo'},
    {'1': 'proofs', '3': 3, '4': 3, '5': 9, '10': 'proofs'},
    {'1': 'items', '3': 4, '4': 2, '5': 9, '10': 'items'},
    {'1': 'remarks', '3': 5, '4': 2, '5': 9, '10': 'remarks'},
    {'1': 'files', '3': 6, '4': 3, '5': 9, '10': 'files'},
    {'1': 'creator', '3': 7, '4': 2, '5': 9, '10': 'creator'},
    {'1': 'customer', '3': 8, '4': 1, '5': 9, '10': 'customer'},
    {'1': 'total', '3': 9, '4': 2, '5': 4, '10': 'total'},
    {'1': 'currency', '3': 10, '4': 2, '5': 9, '10': 'currency'},
    {'1': 'decimals', '3': 11, '4': 2, '5': 13, '10': 'decimals'},
  ],
};

/// Descriptor for `Txn`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List txnDescriptor = $convert.base64Decode(
    'CgNUeG4SDgoCaWQYASACKARSAmlkEhwKCWludm9pY2VObxgCIAIoCVIJaW52b2ljZU5vEhYKBn'
    'Byb29mcxgDIAMoCVIGcHJvb2ZzEhQKBWl0ZW1zGAQgAigJUgVpdGVtcxIYCgdyZW1hcmtzGAUg'
    'AigJUgdyZW1hcmtzEhQKBWZpbGVzGAYgAygJUgVmaWxlcxIYCgdjcmVhdG9yGAcgAigJUgdjcm'
    'VhdG9yEhoKCGN1c3RvbWVyGAggASgJUghjdXN0b21lchIUCgV0b3RhbBgJIAIoBFIFdG90YWwS'
    'GgoIY3VycmVuY3kYCiACKAlSCGN1cnJlbmN5EhoKCGRlY2ltYWxzGAsgAigNUghkZWNpbWFscw'
    '==');

