import 'package:dartz/dartz.dart';
import 'package:uo_yamul_dashboard/common/bloc/selected_app/loading_maps_cubit.dart';
import 'package:uo_yamul_dashboard/common/bloc/selected_app/loading_maps_state.dart';
import 'package:uo_yamul_dashboard/core/usecase/usecase.dart';
import 'package:uo_yamul_dashboard/domain/entities/game_maps.dart';

import '../../../service_locator.dart';

class ShowMapsUsecase extends UseCase<void, void> {

  @override
  Future<UseCaseResponse<void>> call(void param) async {
    var maps = <GameMap>[];
    maps.add(GameMap(0, 'First'));
    maps.add(GameMap(1, 'Second'));
    maps.add(GameMap(2, 'Third'));
    var gameMaps = GameMaps(maps);
    sl<LoadingMapsCubit>().changeState(LoadingMapsStateLoaded(gameMaps));
    return Right(null);
  }
}