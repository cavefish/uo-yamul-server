package dev.cavefish.yamul.backend.game.controller.domain

@SuppressWarnings("MagicNumber")
enum class Flags(val id: Int) {
    Normal(0),
    CanAlterPaperDoll(0x02),
    Poisoned(0x04),
    GoldenHealth(0x08),
    WarMode(0x40),
    Hidden(0x80);
}
