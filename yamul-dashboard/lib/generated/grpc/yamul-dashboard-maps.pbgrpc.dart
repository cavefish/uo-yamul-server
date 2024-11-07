//
//  Generated code. Do not modify.
//  source: yamul-dashboard-maps.proto
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

import 'yamul-dashboard-common.pb.dart' as $1;
import 'yamul-dashboard-maps.pb.dart' as $2;

export 'yamul-dashboard-maps.pb.dart';

@$pb.GrpcServiceName('dashboard.maps.DashboardMapsService')
class DashboardMapsServiceClient extends $grpc.Client {
  static final _$getMaps = $grpc.ClientMethod<$1.Empty, $2.GetMapsResponseItem>(
      '/dashboard.maps.DashboardMapsService/getMaps',
      ($1.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.GetMapsResponseItem.fromBuffer(value));

  DashboardMapsServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options,
        interceptors: interceptors);

  $grpc.ResponseFuture<$2.GetMapsResponseItem> getMaps($1.Empty request, {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getMaps, request, options: options);
  }
}

@$pb.GrpcServiceName('dashboard.maps.DashboardMapsService')
abstract class DashboardMapsServiceBase extends $grpc.Service {
  $core.String get $name => 'dashboard.maps.DashboardMapsService';

  DashboardMapsServiceBase() {
    $addMethod($grpc.ServiceMethod<$1.Empty, $2.GetMapsResponseItem>(
        'getMaps',
        getMaps_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.Empty.fromBuffer(value),
        ($2.GetMapsResponseItem value) => value.writeToBuffer()));
  }

  $async.Future<$2.GetMapsResponseItem> getMaps_Pre($grpc.ServiceCall call, $async.Future<$1.Empty> request) async {
    return getMaps(call, await request);
  }

  $async.Future<$2.GetMapsResponseItem> getMaps($grpc.ServiceCall call, $1.Empty request);
}
