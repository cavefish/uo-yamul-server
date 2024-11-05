import 'dart:developer';

import 'package:auto_route/auto_route.dart';
import 'package:flutter/cupertino.dart';
import 'package:uo_yamul_dashboard/domain/repository/auth.dart';
import 'package:uo_yamul_dashboard/presentation/home/home_page.dart';
import 'package:uo_yamul_dashboard/presentation/loading_page.dart';
import 'package:uo_yamul_dashboard/presentation/login/login_page.dart';
import 'package:uo_yamul_dashboard/presentation/login/login_page_route.dart';
import 'package:uo_yamul_dashboard/presentation/maps/maps_page.dart';

import 'service_locator.dart';

@AutoRouterConfig()
class AppRouter extends RootStackRouter {
  @override
  RouteType get defaultRouteType =>
      RouteType.material(); //.cupertino, .adaptive ..etc

  @override
  List<AutoRoute> get routes => [
        AutoRoute(page: LoginPageRoute.pageInfo),
        _createRoute(
            HomePage.routeName, initialRoute: true, (data) => const HomePage()),
        _createRoute(MapsPage.routeName, (data) => const MapsPage()),
        _createRoute(LoadingPage.routeName, (data) => const LoadingPage()),
      ];

  @override
  late final List<AutoRouteGuard> guards = [
    AutoRouteGuard.simple(this._loggedInGuard)
    // add more guards here
  ];

  AutoRoute _createRoute(String path, Widget Function(RouteData) builder,
      {bool initialRoute = false}) {
    return AutoRoute(
        page: PageInfo(path, builder: builder), initial: initialRoute);
  }

  Future<void> _loggedInGuard(
      NavigationResolver resolver, StackRouter router) async {
    log('Checkin permissions');
    if (resolver.routeName == LoginPage.routeName) return resolver.next();
    var loginInfo = await sl<AuthRepository>().getLoginInfo();
    if (loginInfo != null) return resolver.next();
    log('Cannot access to page ${resolver.routeName}. Redirecting to ${LoginPage.routeName}');
    //router.pushNamed(resolver.routeName);
    resolver.redirect(LoginPageRoute(redirectTo: resolver.routeName));
  }
}
