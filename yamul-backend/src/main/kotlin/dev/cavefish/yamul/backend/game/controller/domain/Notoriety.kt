package dev.cavefish.yamul.backend.game.controller.domain

@SuppressWarnings("MagicNumber")
enum class Notoriety(val id: Int) {
    Unknown(0x00),
    Innocent(0x01),
    Ally(0x02),
    Gray(0x03),
    Criminal(0x04),
    Enemy(0x05),
    Murderer(0x06),
    Invulnerable(0x07)
}