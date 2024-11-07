import 'package:grpc/grpc_web.dart';
import 'package:uo_yamul_dashboard/generated/grpc/yamul-dashboard-login.pbgrpc.dart';

class LoginGrpcService {
  static Future<DashboardLoginServiceClient> createClient() async {

    final channel = GrpcWebClientChannel.xhr(
      Uri.parse('http://localhost:8091'),
    );
    return DashboardLoginServiceClient(channel);
  }
}
