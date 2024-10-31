import 'package:get_it/get_it.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_state.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_state_cubit.dart';

import 'data/repository/mock/auth_repository_mock.dart';
import 'domain/repository/auth.dart';
import 'domain/usecases/auth/login.dart';
import 'domain/usecases/auth/logout.dart';

final sl = GetIt.instance;

void initServiceLocator() {
  // Repositories
  sl.registerSingleton<AuthRepository>(AuthRepositoryMock());

  // StateCubits
  sl.registerSingleton(AuthStateCubit(AuthStateUnknown()));

  // UseCases
  sl.registerSingleton(AuthLoginUsecase());
  sl.registerSingleton(AuthLogoutUsecase());
}
