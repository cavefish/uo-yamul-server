import 'package:dartz/dartz.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_cubit.dart';

import '../../../core/usecase/usecase.dart';
import '../../../service_locator.dart';
import '../../repository/auth.dart';

class AuthLogoutUsecase extends UseCase<void, void> {
  @override
  Future<UseCaseResponse<void>> call(void param) async {
    await sl<AuthRepository>().logout();
    sl<AuthCubit>().refreshState();
    return Right(null);
  }
}
