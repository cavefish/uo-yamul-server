import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_state.dart';
import 'package:uo_yamul_dashboard/domain/entities/login_info.dart';

import '../../../domain/repository/auth.dart';
import '../../../service_locator.dart';

class AuthCubit extends Cubit<AuthState> {
  AuthCubit(super.initialState);

  void init() async {
    await Future.delayed(const Duration(seconds: 1));
    refreshState();
  }

  void refreshState() async {
    await sl<AuthRepository>().getLoginInfo().then((loginState) {
      if (loginState.state == LoginState.authenticated) {
        emit(AuthStateAuthenticated(loginState));
      } else {
        emit(AuthStateUnauthenticated());
      }
    });
  }
}
