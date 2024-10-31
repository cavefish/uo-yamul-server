import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_state_cubit.dart';

import 'common/bloc/auth/auth_state.dart';
import 'presentation/login/pages/home_page.dart';
import 'presentation/login/pages/login_page.dart';
import 'service_locator.dart';

void main() {
  initServiceLocator();
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
        title: 'YAMUL Dashboard',
        theme: ThemeData(
          colorScheme: ColorScheme.fromSeed(seedColor: Colors.green),
          useMaterial3: true,
        ),
        home: MultiBlocProvider(
            providers: [
              BlocProvider<AuthStateCubit>.value(value: sl<AuthStateCubit>()..init())
            ],
            child: BlocBuilder<AuthStateCubit, AuthState>(
                builder: (context, state) {
                  switch(state) {
                    case AuthStateAuthenticated():
                      return HomePage(authState: state);
                    case AuthStateUnauthenticated():
                      return const LoginPage(title: 'Login');
                    default:
                      return const Placeholder(
                        child: Text('Loading'),
                      );
                  }
            })));
  }
}
