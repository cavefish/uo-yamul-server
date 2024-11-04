import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:uo_yamul_dashboard/common/bloc/selected_app/loading_maps_cubit.dart';
import 'package:uo_yamul_dashboard/common/bloc/selected_app/selected_app_state.dart';
import 'package:uo_yamul_dashboard/common/widgets/app_scaffold/yamul_app_scaffold.dart';
import 'package:uo_yamul_dashboard/common/widgets/loading/loading_widget.dart';
import 'package:uo_yamul_dashboard/presentation/maps/game_map_item.dart';

class MapsPage extends StatefulWidget {
  const MapsPage({super.key});

  @override
  State<MapsPage> createState() => _MapsPageState();
}

class _MapsPageState extends State<MapsPage> {
  @override
  Widget build(BuildContext context) {
    return YamulAppScaffold(
      title: 'Maps',
      child: BlocBuilder<LoadingMapsCubit, LoadingMapsState>(
          builder: (BuildContext context, LoadingMapsState state) {
        switch (state) {
          case LoadingMapsStateLoaded():
            return _buildBody(state);
          case LoadingMapsStateLoading():
          default:
            return YamulLoading();
        }
      }),
    );
  }

  ListView _buildBody(LoadingMapsStateLoaded state) {
    return ListView.builder(
        itemCount: state.gameMaps.maps.length,
        itemBuilder: (BuildContext context, int index) {
          return Padding(
            padding: const EdgeInsets.all(8.0),
            child: GameMapItem(state.gameMaps.maps[index]),
          );
        });
  }
}
