import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../core/usecase/usecase.dart';
import 'button_state.dart';

class ButtonStateCubit extends Cubit<ButtonState> {
  ButtonStateCubit() : super(ButtonInitialState());

  void execute({dynamic params, required UseCase usecase}) async {
    emit(ButtonLoadingState());
    try {
      var result = await usecase.call(params).timeout(usecase.timeoutDuration);

      result.fold((error) {
        emit(ButtonFailureState(error));
      }, (data) {
        emit(ButtonSuccessState(data));
      });
    } catch (e) {
      emit(ButtonFailureState('Unexpected exception: $e'));
    }
  }
}
