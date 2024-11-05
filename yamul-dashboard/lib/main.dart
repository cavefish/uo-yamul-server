import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:uo_yamul_dashboard/app_router.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_cubit.dart';
import 'package:uo_yamul_dashboard/common/bloc/selected_app/loading_maps_cubit.dart';

import 'service_locator.dart';

void main() {
  initServiceLocator();
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  final appRouter = AppRouter();

  MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
        providers: [
          BlocProvider<AuthCubit>.value(value: sl<AuthCubit>()..init()),
          BlocProvider<LoadingMapsCubit>.value(value: sl<LoadingMapsCubit>()),
        ],
        child: MaterialApp.router(
          routerConfig: appRouter.config(reevaluateListenable: sl<AuthCubit>().getListenable()),
          title: 'YAMUL Dashboard',
          theme: ThemeData(
            useMaterial3: true,
            colorScheme: ColorScheme.fromSeed(
                brightness: Brightness.light, seedColor: Colors.green),
          ),
        ));
  }
}
