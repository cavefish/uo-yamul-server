package dev.cavefish.yamul.backend.game.controller.domain.mul

import dev.cavefish.yamul.backend.Constants
import dev.cavefish.yamul.backend.game.controller.domain.Coordinates


data class BlockAltitudeData(
    val origin: Coordinates,
    private val mapValues: Array<Array<Pair<Short, Byte>>>,
    private var staticCells: List<StaticCellData>
) {

    init {
        staticCells = staticCells.sortedBy(StaticCellData::z)
    }

    fun getCellAttitude(coordinate: Coordinates): Int {
        val difference = coordinate.difference(origin)
        var lowestZ = mapValues[difference.x][difference.y].second.toInt()
        var currentResult = coordinate.z
        var higherZ = lowestZ
        if (lowestZ > currentResult) currentResult = lowestZ // Under the map, correct to map position
        for (cell in staticCells) {
            if (cell.z > currentResult) break
            if (cell.x != difference.x || cell.y != difference.y) continue
            if (higherZ < cell.z) higherZ = cell.z
            if (cell.z in (lowestZ + 1)..currentResult) { lowestZ = cell.z } // The tile is the new floor
        }
        // TODO this must accept situations like been between too valid objects. I.e. inside a castle
        if (currentResult > higherZ) return higherZ // Floating on objects, return higher valid

        return lowestZ
    }

    override fun hashCode(): Int {
        return origin.hashCode()
    }

    override fun toString(): String {
        val valuesAsString =
            mapValues.joinToString("\n\t") {
                it.joinToString(", ") { it2 ->
                    "${Constants.toHexFormat(it2.first)}:${
                        Constants.toHexFormat(it2.second)
                    }"
                }
            }
        return "BlockAltitudeData(origin=$origin, mapValues=[\n\t$valuesAsString\n], cellValues=[$staticCells])"
    }

    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false

        other as BlockAltitudeData

        if (origin != other.origin) return false
        if (!mapValues.contentDeepEquals(other.mapValues)) return false
        if (staticCells != other.staticCells) return false

        return true
    }


}
