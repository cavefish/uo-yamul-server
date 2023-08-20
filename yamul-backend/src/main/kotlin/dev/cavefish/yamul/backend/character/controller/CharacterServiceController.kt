package dev.cavefish.yamul.backend.character.controller

import dev.cavefish.yamul.backend.character.api.CharacterListResponse
import dev.cavefish.yamul.backend.character.api.CharacterServiceGrpc.CharacterServiceImplBase
import dev.cavefish.yamul.backend.common.api.Empty
import io.grpc.stub.StreamObserver

class CharacterServiceController:CharacterServiceImplBase() {
    override fun getCharacterList(request: Empty?, responseObserver: StreamObserver<CharacterListResponse>?) {
        responseObserver?.onNext(CharacterListResponse.getDefaultInstance())
        responseObserver?.onCompleted()
    }
}