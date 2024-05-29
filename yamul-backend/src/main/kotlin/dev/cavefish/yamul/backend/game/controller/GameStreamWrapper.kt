package dev.cavefish.yamul.backend.game.controller

import dev.cavefish.yamul.backend.game.api.Message
import dev.cavefish.yamul.backend.game.api.MsgType

interface GameStreamWrapper {
    fun send(msgType: MsgType, f: (Message.Builder) -> Unit)
}