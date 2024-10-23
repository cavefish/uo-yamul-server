package dev.cavefish.yamul.backend.game.controller.processors

import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.api.StreamPackage
import dev.cavefish.yamul.backend.game.controller.domain.gamestate.State
import dev.cavefish.yamul.backend.game.controller.GameStreamWrapper
import kotlinx.coroutines.runBlocking
import org.tinylog.kotlin.Logger

abstract class MessageProcessor<T>(
    private val msgType: MsgType,
    private val payloadGetter: (Message) -> T
) {
    protected abstract suspend fun process(
        payload: T,
        state: State,
        wrapper: GameStreamWrapper
    ): State

    fun getType() = msgType

    fun processStreamPackage(
        payload: StreamPackage,
        state: State,
        wrapper: GameStreamWrapper
    ): State = runBlocking {
        val obj = payloadGetter(payload.body)
        return@runBlocking process(obj, state, wrapper)
    }
}
