import 'package:get_it/get_it.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_state.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_cubit.dart';
import 'package:uo_yamul_dashboard/common/bloc/selected_app/loading_maps_cubit.dart';
import 'package:uo_yamul_dashboard/common/bloc/selected_app/loading_maps_state.dart';
import 'package:uo_yamul_dashboard/data/repository/grpc/auth_repository_grpc.dart';
import 'package:uo_yamul_dashboard/domain/usecases/maps/show_maps.dart';

import 'domain/repository/auth.dart';
import 'domain/usecases/auth/login.dart';
import 'domain/usecases/auth/logout.dart';

final sl = GetIt.instance;

void initServiceLocator() {
  // Repositories
  sl.registerSingleton<AuthRepository>(AuthRepositoryGrpc());

  // StateCubits
  sl.registerSingleton(AuthCubit(AuthStateUnknown()));
  sl.registerSingleton(LoadingMapsCubit(LoadingMapsStateLoading()));

  // UseCases
  sl.registerSingleton(AuthLoginUsecase());
  sl.registerSingleton(AuthLogoutUsecase());
  sl.registerSingleton(ShowMapsUsecase());
}
