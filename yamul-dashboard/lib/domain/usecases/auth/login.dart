import 'dart:async';

import 'package:dartz/dartz.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_cubit.dart';

import '../../../common/bloc/selected_app/loading_maps_cubit.dart';
import '../../../common/bloc/selected_app/loading_maps_state.dart';
import '../../../core/usecase/usecase.dart';
import '../../../data/models/auth_singin_params.dart';
import '../../../service_locator.dart';
import '../../repository/auth.dart';

class AuthLoginUsecase extends UseCase<void, AuthLoginParams> {
  AuthLoginUsecase() : super(timeoutDuration: const Duration(seconds: 3));

  @override
  Future<UseCaseResponse<void>> call(AuthLoginParams param) async {
    return sl<AuthRepository>().login(param).then(notifyChange);
  }

  FutureOr<UseCaseResponse<void>> notifyChange(Either<String, void> value) {
    if (value.isRight()) {
      sl<LoadingMapsCubit>().changeState(LoadingMapsStateLoading());
      sl<AuthCubit>().refreshState();
    }
    return value;
  }
}
