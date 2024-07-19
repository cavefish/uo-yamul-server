package dev.cavefish.yamul.backend.game.controller.domain

import dev.cavefish.yamul.backend.game.controller.domain.Hue.Companion.P99
import dev.cavefish.yamul.backend.game.controller.domain.Hue.Companion.P00

data class Hue(val red: UInt, val green: UInt, val blue: UInt) {

    @SuppressWarnings("MagicNumber")
    fun toUInt16():UInt { // TODO move this to the HueMulRepository
        val newR = (red shr 3) and 0b11111u
        val newG = (green shr 3) and 0b11111u
        val newB = (blue shr 3) and 0b11111u
        return (newR shl 10) or (newG shl 5) or (newB) or 0b0_00000_00000_00000u
    }

    companion object {
        const val P99:UInt = 0b11111_000u
        const val P00:UInt = 0b00000_000u
    }
}

@SuppressWarnings("MagicNumber")
enum class Hues(val hue: Hue) {
    Red(Hue(P99, P00, P00)),
    Green(Hue(P00, P99, P00)),
    Blue(Hue(P00, P00, P99)),
    White(Hue(P99, P99, P99)),
    Black(Hue(P00, P00, P00)),
    Character(Hue(P00, P99, 0b01010_000u)),
}
