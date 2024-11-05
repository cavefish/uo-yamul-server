import 'package:auto_route/auto_route.dart';

import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_cubit.dart';
import 'package:uo_yamul_dashboard/common/bloc/auth/auth_state.dart';
import 'package:uo_yamul_dashboard/presentation/login/login_page.dart';
import 'package:uo_yamul_dashboard/presentation/maps/maps_page.dart';

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
        route: MapsPage.routeName,
        name: 'Maps',
        icon: Icons.map_outlined,
        usecase: ShowMapsUsecase,
        params: null),
    YamulDrawerAction(
        route: LoginPage.routeName,
        name: 'Logout',
        icon: Icons.logout,
        usecase: AuthLogoutUsecase,
        params: null)
  ];

  YamulAppScaffold(
      {super.key,
      required this.title,
      required this.child,
      this.showDrawer = true});

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
    List<Widget> children = [];
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

  Widget _buildDrawerAction(BuildContext context, YamulDrawerAction action) {
    return YamulDrawerListTile(selected: context.router.current.name == action.route, action: action);
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
  final Type? usecase;
  final dynamic params;
  final String? route;

  YamulDrawerAction(
      {required this.route,
      required this.name,
      required this.icon,
      required this.usecase,
      required this.params});
}

class YamulDrawerListTile extends StatelessWidget {
  final bool selected;
  final YamulDrawerAction action;

  const YamulDrawerListTile({super.key, required this.selected, required this.action});

  @override
  Widget build(BuildContext context) {
    var themeData = Theme.of(context);
    return ListTile(
      iconColor: themeData.primaryColor,
      selectedTileColor: themeData.focusColor,
      selected: selected,
      leading: Icon(action.icon),
      title: Text(action.name),
      onTap: _createOnTap(context),
    );
  }

  void Function() _createOnTap(BuildContext context) {
    if (action.usecase == null) return () => _navigateTo(action.route, context);
    var usecase = sl<UseCase<dynamic, dynamic>>(type: action.usecase);
    return () async {
      var result = await usecase.call(action.params);
      result.fold((err) => showWarning(context, err),
              (_) => _navigateTo(action.route, context));
    };
  }

  void _navigateTo(String? route, BuildContext context) {
    Navigator.pop(context);
    if (route != null) {
      context.router.replaceNamed(route);
    }
  }

}