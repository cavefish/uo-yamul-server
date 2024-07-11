package dev.cavefish.yamul.backend.game.controller.domain

data class Hue(val red: Int, val green: Int, val blue: Int, val alpha: Int = 0) {

    @SuppressWarnings("MagicNumber")
    fun toInt16():Int {
        val c = alpha shl 24 or red shl 16 or green shl 8 or blue
        return ((((c and 0xFF) shl 5) shr 8)
                or
                (((((c shr 16) and 0xFF) shl 5) shr 8) shl 10)
                or
                (((((c shr 8) and 0xFF) shl 5) shr 8) shl 5))
    }
}

@SuppressWarnings("MagicNumber")
enum class Hues(val hue: Hue) {
    Red(Hue(255, 0, 0)),
    Green(Hue(0, 255, 0)),
    Blue(Hue(0, 0, 255)),
    White(Hue(255, 255, 255)),
    Black(Hue(0, 0, 0))
}
