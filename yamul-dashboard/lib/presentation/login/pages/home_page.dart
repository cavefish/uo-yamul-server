import 'package:flutter/material.dart';
import 'package:uo_yamul_dashboard/domain/usecases/auth/logout.dart';

import '../../../common/bloc/auth/auth_state.dart';
import '../../../core/shimmer_options.dart';
import '../../../service_locator.dart';

class HomePage extends StatefulWidget {
  const HomePage(
      {ShimmerOptions shimmerOptions = ShimmerOptions.none,
      super.key,
      required this.authState});

  final AuthStateAuthenticated authState;

  @override
  State<StatefulWidget> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: const Text('YAMUL Dashboard'),
        centerTitle: true,
      ),
      drawer: Drawer(
        child: ListView(
          padding: EdgeInsets.zero,
          children: [
            _buildDrawerUserInfo(context),
            ListTile(
              title: const Text('Logout'),
              onTap: () => sl<AuthLogoutUsecase>().call(null),
            )
          ],
        ),
      ),
    );
  }

  DrawerHeader _buildDrawerUserInfo(BuildContext context) {
    return DrawerHeader(
      decoration: BoxDecoration(
        color: Theme.of(context).colorScheme.primary,
      ),
      child: Text('Username: ${widget.authState.loginInfo.username}'),
    );
  }
}
