package dev.cavefish.yamul.backend.game.controller;

import dev.cavefish.yamul.backend.game.api.GameServiceGrpc
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.api.StreamPackage
import dev.cavefish.yamul.backend.game.controller.domain.gameevents.GameStreamEventCoordinator
import dev.cavefish.yamul.backend.game.controller.processors.MessageProcessor
import io.grpc.stub.StreamObserver
import org.springframework.stereotype.Component

@Component
class GameServiceController(
    val processors: Map<MsgType, MessageProcessor<*>>,
    val gameStreamEventCoordinator: GameStreamEventCoordinator,
) : GameServiceGrpc.GameServiceImplBase(

) {

    override fun openGameStream(responseObserver: StreamObserver<StreamPackage>?): StreamObserver<StreamPackage>? {
        return GameStreamObserver(responseObserver!!, processors, gameStreamEventCoordinator)
    }

}
