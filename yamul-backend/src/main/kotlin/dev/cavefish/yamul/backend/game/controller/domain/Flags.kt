package dev.cavefish.yamul.backend.game.controller.domain

@SuppressWarnings("MagicNumber")
enum class Flags(val id: FlagsValue) {
    None(0),
    Frozen(0x01),
    Female(0x02),
    Poisoned(0x04),
    Flying(0x04),
    YellowBar(0x08),
    IgnoreMobiles(0x10),
    Movable(0x20),
    WarMode(0x40),
    Hidden(0x80);

    companion object {
        fun and(vararg values: Flags): FlagsValue {
            var result = 0
            for (value in values) {result = value.id or result}
            return result
        }
    }
}

typealias FlagsValue = Int
