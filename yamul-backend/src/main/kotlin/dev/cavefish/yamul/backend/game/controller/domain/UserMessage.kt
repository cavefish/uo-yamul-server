package dev.cavefish.yamul.backend.game.controller.domain

import dev.cavefish.yamul.backend.common.api.MessageType

data class UserMessage(val originObject: GameObject, val text: String, val type: UserMessageType)

enum class UserMessageType(val apiValue: MessageType) {
    Normal(MessageType.MessageType_normal),
    Broadcast(MessageType.MessageType_broadcast),
    Emote(MessageType.MessageType_emote),
    System(MessageType.MessageType_system),
    Label(MessageType.MessageType_label),
    Focus(MessageType.MessageType_focus),
    Whisper(MessageType.MessageType_whisper),
    Yell(MessageType.MessageType_yell),
    Spell(MessageType.MessageType_spell),
    Guild(MessageType.MessageType_guild),
    Alliance(MessageType.MessageType_alliance),
    Command(MessageType.MessageType_command),
    Encoded(MessageType.MessageType_encoded),;

    companion object {
        private val byMessageType = entries.associateBy(UserMessageType::apiValue)
        fun mapFromApi(mode: MessageType?): UserMessageType = byMessageType.getOrDefault(mode, Normal)
    }
}