import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';

void showWarning(BuildContext context, String msg) {
  ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text(msg)));
}
