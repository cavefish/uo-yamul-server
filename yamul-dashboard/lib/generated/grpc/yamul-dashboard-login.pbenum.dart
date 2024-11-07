//
//  Generated code. Do not modify.
//  source: yamul-dashboard-login.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class LoginResponse_LoginResponseValue extends $pb.ProtobufEnum {
  static const LoginResponse_LoginResponseValue undefined = LoginResponse_LoginResponseValue._(0, _omitEnumNames ? '' : 'undefined');
  static const LoginResponse_LoginResponseValue valid = LoginResponse_LoginResponseValue._(1, _omitEnumNames ? '' : 'valid');
  static const LoginResponse_LoginResponseValue invalid = LoginResponse_LoginResponseValue._(2, _omitEnumNames ? '' : 'invalid');

  static const $core.List<LoginResponse_LoginResponseValue> values = <LoginResponse_LoginResponseValue> [
    undefined,
    valid,
    invalid,
  ];

  static final $core.Map<$core.int, LoginResponse_LoginResponseValue> _byValue = $pb.ProtobufEnum.initByValue(values);
  static LoginResponse_LoginResponseValue? valueOf($core.int value) => _byValue[value];

  const LoginResponse_LoginResponseValue._($core.int v, $core.String n) : super(v, n);
}


const _omitEnumNames = $core.bool.fromEnvironment('protobuf.omit_enum_names');
