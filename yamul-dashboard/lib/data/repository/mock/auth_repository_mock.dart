import 'package:dartz/dartz.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:uo_yamul_dashboard/domain/entities/login_info.dart';

import '../../../domain/repository/auth.dart';
import '../../models/auth_singin_params.dart';

class AuthRepositoryMock extends AuthRepository {
  final _logged_in_preference_key = 'loggedInStr2';

  @override
  Future<Either<String, void>> login(AuthLoginParams params) async {
    await Future.delayed(const Duration(seconds: 2));
    if (params.username.isEmpty || params.password.isEmpty) {
      return const Left('Missing parameters');
    }
    if (params.username == params.password) {
      var preferences = await SharedPreferences.getInstance();
      preferences.setString(_logged_in_preference_key, params.username);
      return const Right(null);
    }
    return const Left('Invalid credentials');
  }

  @override
  Future<LoginInfo> getLoginInfo() async {
    var preferences = await SharedPreferences.getInstance();
    var username = preferences.getString(_logged_in_preference_key) ?? '';
    return LoginInfo(
        username,
        (username == '')
            ? LoginState.unauthenticated
            : LoginState.authenticated);
  }

  @override
  Future<void> logout() async {
    var preferences = await SharedPreferences.getInstance();
    preferences.setString(_logged_in_preference_key, '');
  }
}
