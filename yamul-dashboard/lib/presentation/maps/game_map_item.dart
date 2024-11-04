import 'package:flutter/material.dart';
import 'package:uo_yamul_dashboard/common/widgets/two_column_data_card.dart';
import 'package:uo_yamul_dashboard/domain/entities/game_maps.dart';

class GameMapItem extends StatefulWidget {
  final GameMap gameMap;

  const GameMapItem(this.gameMap, {super.key});

  @override
  State<GameMapItem> createState() => _GameMapItemState();
}

class _GameMapItemState extends State<GameMapItem> {
  @override
  Widget build(BuildContext context) {
    return TwoColumnDataCard(titles: [
      'ID',
      'Name'
    ], children: [
      Text('${widget.gameMap.idx}'),
      Text(widget.gameMap.name),
    ]);
  }
}
