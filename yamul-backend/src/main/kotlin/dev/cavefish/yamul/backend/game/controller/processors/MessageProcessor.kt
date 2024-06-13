package dev.cavefish.yamul.backend.game.controller.processors

import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.api.StreamPackage
import dev.cavefish.yamul.backend.game.controller.domain.GameState
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import org.tinylog.Logger
import kotlin.reflect.KFunction

abstract class MessageProcessor<T> (
    private val msgType: MsgType,
    private val payloadGetter: (Message) -> T
){
    protected abstract fun process(payload: T, currentState: GameState, wrapper: GameStreamWrapper): GameState

    fun getType() = msgType

    fun process(payload: StreamPackage, state: GameState, wrapper: GameStreamWrapper): GameState {
        val obj = payloadGetter(payload.body)
        Logger.debug(obj)
        return process(obj, state, wrapper)
    }
}