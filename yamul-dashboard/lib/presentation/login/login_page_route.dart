import 'package:auto_route/auto_route.dart';
import 'package:uo_yamul_dashboard/presentation/home/home_page.dart';

import 'login_page.dart';

class LoginPageRoute extends PageRouteInfo {
  final String redirectTo;

  const LoginPageRoute({required this.redirectTo}) : super(LoginPage.routeName);

  static var pageInfo =
      PageInfo(LoginPage.routeName, builder: (data) => LoginPage(redirectTo: HomePage.routeName));
}
