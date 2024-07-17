package dev.cavefish.yamul.backend.game.controller.domain

import dev.cavefish.yamul.backend.game.controller.domain.Hue.Companion.P99
import dev.cavefish.yamul.backend.game.controller.domain.Hue.Companion.P00

// https://uo.stratics.com/heptazane/fileformats.shtml#1.2
data class Hue(val red: Int, val green: Int, val blue: Int) {

    @SuppressWarnings("MagicNumber")
    fun toInt16():Int {
        val newR = (red shr 3) and 0b11111
        val newG = (green shr 3) and 0b11111
        val newB = (blue shr 3) and 0b11111
        return (newR shl 10) or (newG shl 5) or (newB) or 0b1_00000_00000_00000
    }

    companion object {
        const val P99:Int = 0b11111_000
        const val P00:Int = 0b00000_000
    }
}

@SuppressWarnings("MagicNumber")
enum class Hues(val hue: Hue) {
    Red(Hue(P99, P00, P00)),
    Green(Hue(P00, P99, P00)),
    Blue(Hue(P00, P00, P99)),
    White(Hue(P99, P99, P99)),
    Black(Hue(P00, P00, P00)),
    Character(Hue(P00, P99, 0b01010_000)),
}
