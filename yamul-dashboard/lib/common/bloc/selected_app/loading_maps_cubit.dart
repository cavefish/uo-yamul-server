import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:uo_yamul_dashboard/common/bloc/selected_app/loading_maps_state.dart';

class LoadingMapsCubit extends Cubit<LoadingMapsState> {
  LoadingMapsCubit(super.initialState);

  void changeState(LoadingMapsState next) {
    emit(next);
  }
}