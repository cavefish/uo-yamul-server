package dev.cavefish.yamul.backend.game.controller.processors

import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.api.StreamPackage
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.State
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import org.tinylog.kotlin.Logger

abstract class MessageProcessor<T>(
    private val msgType: MsgType,
    private val payloadGetter: (Message) -> T
) {
    protected abstract fun process(
        payload: T,
        state: State,
        wrapper: GameStreamWrapper
    ): State

    fun getType() = msgType

    fun process(
        payload: StreamPackage,
        state: State,
        wrapper: GameStreamWrapper
    ): State {
        val obj = payloadGetter(payload.body)
        Logger.debug(obj)
        return process(obj, state, wrapper)
    }
}
