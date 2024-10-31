import 'package:dartz/dartz.dart';
import 'package:uo_yamul_dashboard/domain/entities/login_info.dart';

import '../../data/models/auth_singin_params.dart';

abstract class AuthRepository {
  Future<Either<String, void>> login(AuthLoginParams params);
  Future<LoginInfo> getLoginInfo();
  Future<void> logout();
}
