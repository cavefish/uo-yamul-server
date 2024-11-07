import 'dart:developer';

import 'package:auto_route/auto_route.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_cubit.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_state.dart';
import 'package:uo_yamul_dashboard/common/widgets/app_scaffold/yamul_app_scaffold.dart';
import 'package:uo_yamul_dashboard/presentation/warning_snackbar.dart';

import '../../common/bloc/button/button_state_cubit.dart';
import '../../common/widgets/button/yamul_button.dart';
import '../../data/models/auth_singin_params.dart';
import '../../domain/usecases/auth/login.dart';
import '../../service_locator.dart';

class LoginPage extends StatelessWidget {
  final String redirectTo;
  LoginPage({super.key, required this.redirectTo});

  final GlobalKey<FormState> _formKey = GlobalKey<FormState>();

  static const routeName = '/login';


  @override
  Widget build(BuildContext context) {
    return YamulAppScaffold(
      showDrawer: false,
      title: 'Login',
      child: YamulButton.createBloc(
          child: _buildForm(context),
          onSuccess: (state) => {Navigator.pop(context)},
          onFailure: (state) => {showWarning(context, state.errorMessage)}),
    );
  }

  Form _buildForm(BuildContext context) {
    var usernameController = TextEditingController();
    var passwordController = TextEditingController();
    return Form(
      key: _formKey,
      child: Center(
        child: Container(
          constraints: BoxConstraints.expand(
            height:
                Theme.of(context).textTheme.headlineMedium!.fontSize! * 1.1 +
                    200.0,
            width: Theme.of(context).textTheme.headlineMedium!.fontSize! * 1.1 +
                300.0,
          ),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            mainAxisSize: MainAxisSize.min,
            children: <Widget>[
              const Text(
                'Username',
              ),
              _buildUsernameField(usernameController),
              const Text(
                'Password',
              ),
              _buildPasswordField(passwordController),
              Builder(builder: (context) {
                return BlocListener<AuthCubit, AuthState>(
                  listener: (BuildContext context, AuthState state) {
                    if (state is AuthStateAuthenticated) {
                      log('Redirecting to $redirectTo');
                      context.router.replaceNamed(this.redirectTo);
                    }
                  },
                  child: YamulButton(
                    onPressed: () {
                      var currentState = _formKey.currentState;
                      if (currentState == null) return;
                      if (!currentState.validate()) return;
                      context.read<ButtonStateCubit>().execute(
                          usecase: sl<AuthLoginUsecase>(),
                          params: AuthLoginParams(
                              usernameController.text, passwordController.text));
                    },
                    title: 'Submit',
                  ),
                );
              }),
            ],
          ),
        ),
      ),
    );
  }

  TextFormField _buildUsernameField(TextEditingController usernameController) {
    return TextFormField(
      controller: usernameController,
      decoration: const InputDecoration(hintText: 'Type your username'),
      validator: _validateValueNotEmpty,
    );
  }

  TextFormField _buildPasswordField(TextEditingController passwordController) {
    return TextFormField(
      controller: passwordController,
      obscureText: true,
      enableSuggestions: false,
      autocorrect: false,
      decoration: const InputDecoration(hintText: 'Type your password'),
      validator: _validateValueNotEmpty,
    );
  }

  String? _validateValueNotEmpty(String? value) {
    if (value == null || value.isEmpty) {
      return 'Please enter some text';
    }
    return null;
  }
}
