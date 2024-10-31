import 'package:uo_yamul_dashboard/domain/entities/login_info.dart';

sealed class AuthState {}

class AuthStateAuthenticated extends AuthState {
  LoginInfo loginInfo;

  AuthStateAuthenticated(this.loginInfo);
}

class AuthStateUnauthenticated extends AuthState {}

class AuthStateUnknown extends AuthState {}