//
//  Generated code. Do not modify.
//  source: yamul-dashboard-login.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:convert' as $convert;
import 'dart:core' as $core;
import 'dart:typed_data' as $typed_data;

@$core.Deprecated('Use loginRequestDescriptor instead')
const LoginRequest$json = {
  '1': 'LoginRequest',
  '2': [
    {'1': 'username', '3': 1, '4': 1, '5': 9, '10': 'username'},
    {'1': 'password', '3': 2, '4': 1, '5': 9, '10': 'password'},
  ],
};

/// Descriptor for `LoginRequest`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List loginRequestDescriptor = $convert.base64Decode(
    'CgxMb2dpblJlcXVlc3QSGgoIdXNlcm5hbWUYASABKAlSCHVzZXJuYW1lEhoKCHBhc3N3b3JkGA'
    'IgASgJUghwYXNzd29yZA==');

@$core.Deprecated('Use loginResponseDescriptor instead')
const LoginResponse$json = {
  '1': 'LoginResponse',
  '2': [
    {'1': 'value', '3': 1, '4': 1, '5': 14, '6': '.dashboard.login.LoginResponse.LoginResponseValue', '10': 'value'},
  ],
  '4': [LoginResponse_LoginResponseValue$json],
};

@$core.Deprecated('Use loginResponseDescriptor instead')
const LoginResponse_LoginResponseValue$json = {
  '1': 'LoginResponseValue',
  '2': [
    {'1': 'undefined', '2': 0},
    {'1': 'valid', '2': 1},
    {'1': 'invalid', '2': 2},
  ],
};

/// Descriptor for `LoginResponse`. Decode as a `google.protobuf.DescriptorProto`.
final $typed_data.Uint8List loginResponseDescriptor = $convert.base64Decode(
    'Cg1Mb2dpblJlc3BvbnNlEkcKBXZhbHVlGAEgASgOMjEuZGFzaGJvYXJkLmxvZ2luLkxvZ2luUm'
    'VzcG9uc2UuTG9naW5SZXNwb25zZVZhbHVlUgV2YWx1ZSI7ChJMb2dpblJlc3BvbnNlVmFsdWUS'
    'DQoJdW5kZWZpbmVkEAASCQoFdmFsaWQQARILCgdpbnZhbGlkEAI=');

