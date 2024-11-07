import 'package:dartz/dartz.dart';
import 'package:uo_yamul_dashboard/core/grpc/login_grpc_service.dart';
import 'package:uo_yamul_dashboard/data/models/auth_singin_params.dart';
import 'package:uo_yamul_dashboard/domain/entities/login_info.dart';
import 'package:uo_yamul_dashboard/domain/repository/auth.dart';
import 'package:uo_yamul_dashboard/generated/grpc/yamul-dashboard-login.pb.dart';

class AuthRepositoryGrpc extends AuthRepository {
  LoginInfo? loginInfo;

  @override
  Future<LoginInfo?> getLoginInfo() async {
    return loginInfo;
  }

  @override
  Future<Either<String, void>> login(AuthLoginParams params) async {
    var client = await LoginGrpcService.createClient();
    var response = await client.validateLogin(LoginRequest(username: params.username, password: params.password));
    if (response.value == LoginResponse_LoginResponseValue.valid) {
      this.loginInfo = LoginInfo(params.username);
      return const Right(null);
    }
    return const Left('Invalid credentials');
  }

  @override
  Future<void> logout() async {
    loginInfo = null;
  }

}