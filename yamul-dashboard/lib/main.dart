import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_cubit.dart';
import 'package:uo_yamul_dashboard/common/bloc/selected_app/loading_maps_cubit.dart';
import 'package:uo_yamul_dashboard/presentation/maps/maps_page.dart';

import 'common/bloc/auth/auth_state.dart';
import 'presentation/home/home_page.dart';
import 'presentation/login/login_page.dart';
import 'service_locator.dart';

void main() {
  initServiceLocator();
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
        providers: [
          BlocProvider<AuthCubit>.value(value: sl<AuthCubit>()..init()),
          BlocProvider<LoadingMapsCubit>.value(value: sl<LoadingMapsCubit>()),
        ],
        child: MaterialApp(
            title: 'YAMUL Dashboard',
            theme: ThemeData(
              useMaterial3: true,
              colorScheme: ColorScheme.fromSeed(
                  brightness: Brightness.light, seedColor: Colors.green),
            ),
            initialRoute: '/login',
            routes: {
              '/': (context) => const HomePage(),
              '/maps': (context) => const MapsPage(),
              '/login': (context) => LoginPage(),
              '/loading': (context) => const Placeholder(
                    child: Text('Loading'),
                  ),
            }));
  }
}
