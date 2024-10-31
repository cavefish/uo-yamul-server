import 'package:dartz/dartz.dart';

typedef UseCaseResponse<T> = Either<String, T>;

abstract class UseCase<Type, Param> {
  final Duration timeoutDuration;

  UseCase({this.timeoutDuration = const Duration(seconds: 1)});

  Future<UseCaseResponse<Type>> call(Param param);
}
