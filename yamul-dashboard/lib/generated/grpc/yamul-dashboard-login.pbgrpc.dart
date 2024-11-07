//
//  Generated code. Do not modify.
//  source: yamul-dashboard-login.proto
//
// @dart = 2.12

// ignore_for_file: annotate_overrides, camel_case_types, comment_references
// ignore_for_file: constant_identifier_names, library_prefixes
// ignore_for_file: non_constant_identifier_names, prefer_final_fields
// ignore_for_file: unnecessary_import, unnecessary_this, unused_import

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'package:protobuf/protobuf.dart' as $pb;

import 'yamul-dashboard-login.pb.dart' as $0;

export 'yamul-dashboard-login.pb.dart';

@$pb.GrpcServiceName('dashboard.login.DashboardLoginService')
class DashboardLoginServiceClient extends $grpc.Client {
  static final _$validateLogin = $grpc.ClientMethod<$0.LoginRequest, $0.LoginResponse>(
      '/dashboard.login.DashboardLoginService/validateLogin',
      ($0.LoginRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.LoginResponse.fromBuffer(value));

  DashboardLoginServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$0.LoginResponse> validateLogin($0.LoginRequest request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$validateLogin, request, options: options);
  }
}

@$pb.GrpcServiceName('dashboard.login.DashboardLoginService')
abstract class DashboardLoginServiceBase extends $grpc.Service {
  $core.String get $name => 'dashboard.login.DashboardLoginService';

  DashboardLoginServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.LoginRequest, $0.LoginResponse>(
        'validateLogin',
        validateLogin_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.LoginRequest.fromBuffer(value),
        ($0.LoginResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.LoginResponse> validateLogin_Pre($grpc.ServiceCall call, $async.Future<$0.LoginRequest> request) async {
    return validateLogin(call, await request);
  }

  $async.Future<$0.LoginResponse> validateLogin($grpc.ServiceCall call, $0.LoginRequest request);
}
