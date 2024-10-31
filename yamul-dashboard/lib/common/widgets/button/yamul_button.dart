import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../bloc/button/button_state.dart';
import '../../bloc/button/button_state_cubit.dart';

class YamulButton extends StatelessWidget {
  final VoidCallback onPressed;
  final String title;
  final double? height;
  final double? width;

  const YamulButton(
      {required this.onPressed,
      this.title = '',
      this.height,
      this.width,
      super.key});

  @override
  Widget build(BuildContext context) {
    return BlocBuilder<ButtonStateCubit, ButtonState>(
      builder: (context, state) {
        if (state is ButtonLoadingState) {
          return _loading(context);
        }
        return _initial(context);
      },
    );
  }

  Widget _loading(BuildContext context) {
    return ElevatedButton(
        onPressed: null,
        style: ElevatedButton.styleFrom(
          disabledBackgroundColor: Theme.of(context).colorScheme.secondary,
          minimumSize:
              Size(width ?? MediaQuery.of(context).size.width, height ?? 60),
        ),
        child: const CircularProgressIndicator.adaptive(
          backgroundColor: Colors.white70,
        ));
  }

  Widget _initial(BuildContext context) {
    return Container(
      decoration:
          BoxDecoration(borderRadius: BorderRadius.circular(16), boxShadow: [
        BoxShadow(
          color: Theme.of(context).colorScheme.inversePrimary.withOpacity(0.8),
          offset: const Offset(0, 5),
          blurRadius: 17,
        )
      ]),
      child: ElevatedButton(
          onPressed: onPressed,
          style: ElevatedButton.styleFrom(
            minimumSize:
                Size(width ?? MediaQuery.of(context).size.width, height ?? 60),
            backgroundColor: Theme.of(context).colorScheme.inversePrimary,
          ),
          child: Text(
            title,
            style: TextStyle(
              color: Theme.of(context).colorScheme.primary,
            ),
          )),
    );
  }

  static BlocProvider<ButtonStateCubit> createBloc({
    required StatefulWidget child,
    void Function(dynamic state)? onSuccess,
    void Function(dynamic state)? onFailure,
  }) {
    return BlocProvider(
      create: (context) => ButtonStateCubit(),
      child: BlocListener<ButtonStateCubit, ButtonState>(
        listener: (BuildContext context, state) {
          switch (state) {
            case ButtonSuccessState():
              onSuccess?.call(state);
            case ButtonFailureState():
              onFailure?.call(state);
            case ButtonInitialState():
            // Nothing to do
            case ButtonLoadingState():
            // Nothing to do
          }
        },
        child: child,
      ),
    );
  }
}
