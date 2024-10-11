package dev.cavefish.yamul.backend.game.controller.domain.mul

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates


data class BlockAltitudeData(val origin: Coordinates, val values: Array<Array<Int>>) {
    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false

        other as BlockAltitudeData

        if (origin != other.origin) return false
        if (!values.contentDeepEquals(other.values)) return false

        return true
    }

    override fun hashCode(): Int {
        return origin.hashCode()
    }

    fun getCellAttitude(coordinate: Coordinates): Int {
        val difference = coordinate.difference(origin)
        return values[difference.x][difference.y]
    }

    override fun toString(): String {
        return "BlockAltitudeData(origin=$origin, values=[${values.joinToString { it.contentToString() }}])"
    }


}
