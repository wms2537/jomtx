//
//  Generated code. Do not modify.
//  source: proto/jomtx.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

class QueryGetTxnRequest extends $pb.GeneratedMessage {
  factory QueryGetTxnRequest({
    $fixnum.Int64? id,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    return $result;
  }
  QueryGetTxnRequest._() : super();
  factory QueryGetTxnRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetTxnRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'QueryGetTxnRequest', createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, _omitFieldNames ? '' : 'id', $pb.PbFieldType.QU6, defaultOrMaker: $fixnum.Int64.ZERO)
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetTxnRequest clone() => QueryGetTxnRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetTxnRequest copyWith(void Function(QueryGetTxnRequest) updates) => super.copyWith((message) => updates(message as QueryGetTxnRequest)) as QueryGetTxnRequest;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static QueryGetTxnRequest create() => QueryGetTxnRequest._();
  QueryGetTxnRequest createEmptyInstance() => create();
  static $pb.PbList<QueryGetTxnRequest> createRepeated() => $pb.PbList<QueryGetTxnRequest>();
  @$core.pragma('dart2js:noInline')
  static QueryGetTxnRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetTxnRequest>(create);
  static QueryGetTxnRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class QueryGetTxnResponse extends $pb.GeneratedMessage {
  factory QueryGetTxnResponse({
    Txn? txn,
  }) {
    final $result = create();
    if (txn != null) {
      $result.txn = txn;
    }
    return $result;
  }
  QueryGetTxnResponse._() : super();
  factory QueryGetTxnResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory QueryGetTxnResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'QueryGetTxnResponse', createEmptyInstance: create)
    ..aQM<Txn>(1, _omitFieldNames ? '' : 'Txn', protoName: 'Txn', subBuilder: Txn.create)
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  QueryGetTxnResponse clone() => QueryGetTxnResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  QueryGetTxnResponse copyWith(void Function(QueryGetTxnResponse) updates) => super.copyWith((message) => updates(message as QueryGetTxnResponse)) as QueryGetTxnResponse;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static QueryGetTxnResponse create() => QueryGetTxnResponse._();
  QueryGetTxnResponse createEmptyInstance() => create();
  static $pb.PbList<QueryGetTxnResponse> createRepeated() => $pb.PbList<QueryGetTxnResponse>();
  @$core.pragma('dart2js:noInline')
  static QueryGetTxnResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<QueryGetTxnResponse>(create);
  static QueryGetTxnResponse? _defaultInstance;

  @$pb.TagNumber(1)
  Txn get txn => $_getN(0);
  @$pb.TagNumber(1)
  set txn(Txn v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasTxn() => $_has(0);
  @$pb.TagNumber(1)
  void clearTxn() => clearField(1);
  @$pb.TagNumber(1)
  Txn ensureTxn() => $_ensure(0);
}

class Txn extends $pb.GeneratedMessage {
  factory Txn({
    $fixnum.Int64? id,
    $core.String? invoiceNo,
    $core.Iterable<$core.String>? proofs,
    $core.String? items,
    $core.String? remarks,
    $core.Iterable<$core.String>? files,
    $core.String? creator,
    $core.String? customer,
    $fixnum.Int64? total,
    $core.String? currency,
    $core.int? decimals,
  }) {
    final $result = create();
    if (id != null) {
      $result.id = id;
    }
    if (invoiceNo != null) {
      $result.invoiceNo = invoiceNo;
    }
    if (proofs != null) {
      $result.proofs.addAll(proofs);
    }
    if (items != null) {
      $result.items = items;
    }
    if (remarks != null) {
      $result.remarks = remarks;
    }
    if (files != null) {
      $result.files.addAll(files);
    }
    if (creator != null) {
      $result.creator = creator;
    }
    if (customer != null) {
      $result.customer = customer;
    }
    if (total != null) {
      $result.total = total;
    }
    if (currency != null) {
      $result.currency = currency;
    }
    if (decimals != null) {
      $result.decimals = decimals;
    }
    return $result;
  }
  Txn._() : super();
  factory Txn.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Txn.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);

  static final $pb.BuilderInfo _i = $pb.BuilderInfo(_omitMessageNames ? '' : 'Txn', createEmptyInstance: create)
    ..a<$fixnum.Int64>(1, _omitFieldNames ? '' : 'id', $pb.PbFieldType.QU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aQS(2, _omitFieldNames ? '' : 'invoiceNo', protoName: 'invoiceNo')
    ..pPS(3, _omitFieldNames ? '' : 'proofs')
    ..aQS(4, _omitFieldNames ? '' : 'items')
    ..aQS(5, _omitFieldNames ? '' : 'remarks')
    ..pPS(6, _omitFieldNames ? '' : 'files')
    ..aQS(7, _omitFieldNames ? '' : 'creator')
    ..aOS(8, _omitFieldNames ? '' : 'customer')
    ..a<$fixnum.Int64>(9, _omitFieldNames ? '' : 'total', $pb.PbFieldType.QU6, defaultOrMaker: $fixnum.Int64.ZERO)
    ..aQS(10, _omitFieldNames ? '' : 'currency')
    ..a<$core.int>(11, _omitFieldNames ? '' : 'decimals', $pb.PbFieldType.QU3)
  ;

  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  Txn clone() => Txn()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  Txn copyWith(void Function(Txn) updates) => super.copyWith((message) => updates(message as Txn)) as Txn;

  $pb.BuilderInfo get info_ => _i;

  @$core.pragma('dart2js:noInline')
  static Txn create() => Txn._();
  Txn createEmptyInstance() => create();
  static $pb.PbList<Txn> createRepeated() => $pb.PbList<Txn>();
  @$core.pragma('dart2js:noInline')
  static Txn getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Txn>(create);
  static Txn? _defaultInstance;

  @$pb.TagNumber(1)
  $fixnum.Int64 get id => $_getI64(0);
  @$pb.TagNumber(1)
  set id($fixnum.Int64 v) { $_setInt64(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get invoiceNo => $_getSZ(1);
  @$pb.TagNumber(2)
  set invoiceNo($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasInvoiceNo() => $_has(1);
  @$pb.TagNumber(2)
  void clearInvoiceNo() => clearField(2);

  @$pb.TagNumber(3)
  $core.List<$core.String> get proofs => $_getList(2);

  @$pb.TagNumber(4)
  $core.String get items => $_getSZ(3);
  @$pb.TagNumber(4)
  set items($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasItems() => $_has(3);
  @$pb.TagNumber(4)
  void clearItems() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get remarks => $_getSZ(4);
  @$pb.TagNumber(5)
  set remarks($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasRemarks() => $_has(4);
  @$pb.TagNumber(5)
  void clearRemarks() => clearField(5);

  @$pb.TagNumber(6)
  $core.List<$core.String> get files => $_getList(5);

  @$pb.TagNumber(7)
  $core.String get creator => $_getSZ(6);
  @$pb.TagNumber(7)
  set creator($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasCreator() => $_has(6);
  @$pb.TagNumber(7)
  void clearCreator() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get customer => $_getSZ(7);
  @$pb.TagNumber(8)
  set customer($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasCustomer() => $_has(7);
  @$pb.TagNumber(8)
  void clearCustomer() => clearField(8);

  @$pb.TagNumber(9)
  $fixnum.Int64 get total => $_getI64(8);
  @$pb.TagNumber(9)
  set total($fixnum.Int64 v) { $_setInt64(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasTotal() => $_has(8);
  @$pb.TagNumber(9)
  void clearTotal() => clearField(9);

  @$pb.TagNumber(10)
  $core.String get currency => $_getSZ(9);
  @$pb.TagNumber(10)
  set currency($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasCurrency() => $_has(9);
  @$pb.TagNumber(10)
  void clearCurrency() => clearField(10);

  @$pb.TagNumber(11)
  $core.int get decimals => $_getIZ(10);
  @$pb.TagNumber(11)
  set decimals($core.int v) { $_setUnsignedInt32(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasDecimals() => $_has(10);
  @$pb.TagNumber(11)
  void clearDecimals() => clearField(11);
}


const _omitFieldNames = $core.bool.fromEnvironment('protobuf.omit_field_names');
const _omitMessageNames = $core.bool.fromEnvironment('protobuf.omit_message_names');
