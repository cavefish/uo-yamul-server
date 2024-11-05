import 'package:dartz/dartz.dart';
import 'package:uo_yamul_dashboard/domain/entities/login_info.dart';

import '../../../domain/repository/auth.dart';
import '../../models/auth_singin_params.dart';

class AuthRepositoryMock extends AuthRepository {
  LoginInfo? loginInfo;

  @override
  Future<Either<String, void>> login(AuthLoginParams params) async {
    await Future.delayed(const Duration(seconds: 2));
    if (params.username.isEmpty || params.password.isEmpty) {
      return const Left('Missing parameters');
    }
    if (params.username == params.password) {
      this.loginInfo = LoginInfo(params.username);
      return const Right(null);
    }
    return const Left('Invalid credentials');
  }

  @override
  Future<LoginInfo?> getLoginInfo() async {
    return loginInfo;
  }

  @override
  Future<void> logout() async {
    loginInfo = null;
  }
}
