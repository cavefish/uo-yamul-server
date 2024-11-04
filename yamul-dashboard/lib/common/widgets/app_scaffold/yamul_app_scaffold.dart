import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_cubit.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_state.dart';

import '../../../core/usecase/usecase.dart';
import '../../../domain/usecases/auth/logout.dart';
import '../../../domain/usecases/maps/show_maps.dart';
import '../../../presentation/warning_snackbar.dart';
import '../../../service_locator.dart';

class YamulAppScaffold extends StatelessWidget {
  final String title;
  final Widget child;
  final bool showDrawer;

  // TODO convert into a service
  final List<YamulDrawerAction> actions = [
    YamulDrawerAction(
        route: '/maps',
        name: 'Maps',
        icon: Icons.map_outlined,
        type: ShowMapsUsecase,
        params: null),
    YamulDrawerAction(
        route: '/login',
        name: 'Logout',
        icon: Icons.logout,
        type: AuthLogoutUsecase,
        params: null)
  ];

  YamulAppScaffold({super.key, required this.title, required this.child, this.showDrawer = true});

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        appBar: AppBar(
          automaticallyImplyLeading: showDrawer,
          backgroundColor: Theme.of(context).colorScheme.inversePrimary,
          title: Text(title),
          centerTitle: true,
        ),
        drawer: _buildDrawer(context),
        body: child,
      ),
    );
  }

  Drawer? _buildDrawer(BuildContext context) {
    if (!this.showDrawer) return null;
    List<StatelessWidget> children = [];
    children.add(_buildDrawerUserInfo(context));
    for (var action in actions) {
      children.add(_buildDrawerAction(context, action));
    }
    return Drawer(
        child: ListView(
          padding: EdgeInsets.zero,
          children: children,
        ),
      );
  }

  ListTile _buildDrawerAction(BuildContext context, YamulDrawerAction action) {
    var usecase = sl<UseCase<dynamic, dynamic>>(type: action.type);
    return ListTile(
      leading: Icon(action.icon),
      title: Text(action.name),
      onTap: () async {
        var result = await usecase.call(action.params);
        result.fold((err) {
          showWarning(context, err);
        }, (_) {
          var route = action.route;
          if (route != null) {
            Navigator.pushReplacementNamed(context, route);
          }
        });
      },
    );
  }

  DrawerHeader _buildDrawerUserInfo(BuildContext context) {
    return DrawerHeader(
        decoration: BoxDecoration(
          color: Theme.of(context).colorScheme.inversePrimary,
        ),
        child: BlocBuilder<AuthCubit, AuthState>(
            builder: (BuildContext context, AuthState state) {
          switch (state) {
            case AuthStateAuthenticated():
              return Text('Username: ${state.loginInfo.username}');
            default:
              return Text('Not logged in');
          }
        }));
  }
}

class YamulDrawerAction {
  final String name;
  final IconData icon;
  final Type type;
  final dynamic params;
  final String? route;

  YamulDrawerAction(
      {required this.route,
      required this.name,
      required this.icon,
      required this.type,
      required this.params});
}
