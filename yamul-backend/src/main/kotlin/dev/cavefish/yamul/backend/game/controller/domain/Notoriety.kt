package dev.cavefish.yamul.backend.game.controller.domain

@SuppressWarnings("MagicNumber")
enum class Notoriety(val id: NotorietyValue) {
    Unknown(0x00),
    Innocent(0x01),
    Ally(0x02),
    Gray(0x03),
    Criminal(0x04),
    Enemy(0x05),
    Murderer(0x06),
    Invulnerable(0x07);

    companion object {
        fun and(vararg values: Notoriety): NotorietyValue {
            var result = 0
            for (value in values) {result = value.id or result}
            return result
        }
    }
}

typealias NotorietyValue = Int
