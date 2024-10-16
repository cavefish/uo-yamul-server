package dev.cavefish.yamul.backend.game.controller.domain.mul

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates


@OptIn(ExperimentalStdlibApi::class)
data class BlockAltitudeData(val origin: Coordinates, val values: Array<Array<Pair<Int, Int>>>) {
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
        return values[difference.x][difference.y].second
    }

    override fun toString(): String {
        val valuesAsString =
            values.joinToString {
                it.joinToString { it2 ->
                    "${it2.first.toShort().toHexString(numberFormat)}:${it2.second.toByte().toHexString(numberFormat)}"
                }
            }
        return "BlockAltitudeData(origin=$origin, values=[$valuesAsString])"
    }

    companion object {
        val numberFormat = HexFormat {
            upperCase = true
            number {
                removeLeadingZeros = false
                prefix = "0x"
            }
        }
    }


}
