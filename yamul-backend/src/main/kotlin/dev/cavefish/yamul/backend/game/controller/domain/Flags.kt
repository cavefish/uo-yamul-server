package dev.cavefish.yamul.backend.game.controller.domain

@SuppressWarnings("MagicNumber")
enum class Flags(val id: Int) {
    None(0),
    Frozen(0x01),
    Female(0x02),
    Poisoned(0x04),
    Flying(0x04),
    YellowBar(0x08),
    IgnoreMobiles(0x10),
    Movable(0x20),
    WarMode(0x40),
    Hidden(0x80)
}