import 'package:flutter/material.dart';

class LoadingPage extends StatelessWidget {
  static var routeName = '/loading';

  const LoadingPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Placeholder(
      child: Text('Loading'),
    );
  }
}
