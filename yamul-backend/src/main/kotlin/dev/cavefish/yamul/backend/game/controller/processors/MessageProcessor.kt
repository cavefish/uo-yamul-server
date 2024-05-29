package dev.cavefish.yamul.backend.game.controller.processors

import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.api.StreamPackage
import dev.cavefish.yamul.backend.game.controller.domain.GameState
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import org.tinylog.Logger

abstract class MessageProcessor<T> {
    abstract fun getType(): MsgType
    protected abstract fun getPayload(payload: StreamPackage): T
    protected abstract fun process(payload: T, currentState: GameState, wrapper: GameStreamWrapper): GameState

    fun process(payload: StreamPackage, state: GameState, wrapper: GameStreamWrapper): GameState {
        val obj = getPayload(payload)
        Logger.debug(obj)
        return process(obj, state, wrapper)
    }
}