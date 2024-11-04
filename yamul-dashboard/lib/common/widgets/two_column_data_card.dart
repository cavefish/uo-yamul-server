import 'package:flutter/material.dart';

class TwoColumnDataCard extends StatefulWidget {
  const TwoColumnDataCard(
      {required this.titles, required this.children, super.key});

  final List<String> titles;
  final List<Widget> children;

  @override
  State<TwoColumnDataCard> createState() => _TwoColumnDataCardState();
}

class _TwoColumnDataCardState extends State<TwoColumnDataCard> {
  @override
  Widget build(BuildContext context) {
    assert(widget.titles.length <= widget.children.length);

    var tableChildren = <TableRow>[];
    var padding = const EdgeInsets.all(2.0);
    for (int i = 0; i < widget.titles.length; i++) {
      tableChildren.add(TableRow(children: [
        Container(
            padding: padding,
            alignment: Alignment.topRight,
            child: Text(widget.titles[i])),
        Container(
            padding: padding,
            alignment: Alignment.topLeft,
            child: widget.children[i])
      ]));
    }
    for (int i = widget.titles.length; i < widget.children.length; i++) {
      tableChildren.add(TableRow(children: [Container(), widget.children[i]]));
    }

    return Card(
      child: SizedBox(
        //height: 100,
        child: Padding(
          padding: const EdgeInsets.all(2.0),
          child: Table(
            columnWidths: {0: IntrinsicColumnWidth(), 1: FlexColumnWidth()},
            children: tableChildren,
          ),
        ),
      ),
    );
  }
}
