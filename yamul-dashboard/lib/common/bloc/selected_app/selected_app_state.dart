import 'package:uo_yamul_dashboard/domain/entities/game_maps.dart';

sealed class LoadingMapsState{}

class LoadingMapsStateLoading extends LoadingMapsState {}

class LoadingMapsStateLoaded extends LoadingMapsState {
  GameMaps gameMaps;

  LoadingMapsStateLoaded(this.gameMaps);
}