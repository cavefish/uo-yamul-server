sealed class ButtonState {}

class ButtonInitialState extends ButtonState {}

class ButtonLoadingState extends ButtonState {}

class ButtonSuccessState extends ButtonState {
  final dynamic body;

  ButtonSuccessState(this.body);
}

class ButtonFailureState extends ButtonState {
  final String errorMessage;

  ButtonFailureState(this.errorMessage);
}
