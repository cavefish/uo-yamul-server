package dev.cavefish.yamul.backend.game.controller;

import dev.cavefish.yamul.backend.game.api.GameServiceGrpc
import dev.cavefish.yamul.backend.game.api.StreamPackage
import io.grpc.stub.StreamObserver
import org.springframework.stereotype.Component

@Component
class GameServiceController : GameServiceGrpc.GameServiceImplBase() {

    override fun openGameStream(responseObserver: StreamObserver<StreamPackage>?): StreamObserver<StreamPackage>? {
        return GameStreamObserver(responseObserver!!)
    }

}
