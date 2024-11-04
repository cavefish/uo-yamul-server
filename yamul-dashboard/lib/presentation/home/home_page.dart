import 'package:flutter/material.dart';
import 'package:uo_yamul_dashboard/common/widgets/app_scaffold/yamul_app_scaffold.dart';
import 'package:uo_yamul_dashboard/common/widgets/wip/wip_widget.dart';


class HomePage extends StatelessWidget {
  const HomePage({super.key});

  @override
  Widget build(BuildContext context) {
    return YamulAppScaffold(
        title: 'Home',
        child: YamulWip()
    );
  }
}
