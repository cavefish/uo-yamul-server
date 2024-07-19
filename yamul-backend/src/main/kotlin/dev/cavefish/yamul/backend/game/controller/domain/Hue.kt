package dev.cavefish.yamul.backend.game.controller.domain

@SuppressWarnings("MagicNumber")
data class Hue(val group: UInt, val entry: UInt) {

    fun toUInt16():UInt {
        return (group shl 3) or (entry and 0b111u)
    }

    companion object {
        fun fromHex(hex: UInt): Hue = Hue(hex shr 3, hex and 0b111u)
    }
}

@SuppressWarnings("MagicNumber")
enum class Hues(val hue: Hue) {
    Red(Hue.fromHex(0x641u)),
    Green(Hue.fromHex(0x0579u)),
    Blue(Hue.fromHex(0x0515u)),
    White(Hue.fromHex(0x07f6u)),
    Black(Hue.fromHex(0x0497u)),
    Character(Hue.fromHex(0x03eau)),
}
