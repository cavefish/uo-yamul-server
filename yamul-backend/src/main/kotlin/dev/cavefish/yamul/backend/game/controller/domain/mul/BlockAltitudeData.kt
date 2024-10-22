package dev.cavefish.yamul.backend.game.controller.domain.mul

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.utils.toStringByRecursion

sealed class BlockAltitudeData {
    abstract val origin: Coordinates
    abstract fun getCellAttitude(coordinate: Coordinates): Int

    companion object {
        fun create(
            origin: Coordinates,
            mapValues: Array<Array<Pair<Short, Byte>>>,
            staticCells: List<StaticCellData>?
        ): BlockAltitudeData {
            if (staticCells.isNullOrEmpty()) {
                return BlockAltitudeDataNoStatics(
                    origin = origin,
                    mapValues = mapValues,
                )
            }
            return BlockAltitudeDataWithStatics.create(
                origin = origin,
                mapValues = mapValues,
                staticCells = staticCells
            )
        }
    }
}

private data class BlockAltitudeDataNoStatics(
    override val origin: Coordinates,
    private val mapValues: Array<Array<Pair<Short, Byte>>>,
) : BlockAltitudeData() {

    override fun getCellAttitude(coordinate: Coordinates): Int {
        val difference = coordinate.difference(origin)
        return mapValues[difference.x][difference.y].second.toInt()
    }

    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false

        other as BlockAltitudeDataNoStatics

        if (origin != other.origin) return false
        if (!mapValues.contentDeepEquals(other.mapValues)) return false

        return true
    }

    override fun hashCode(): Int {
        var result = origin.hashCode()
        result = 31 * result + mapValues.contentDeepHashCode()
        return result
    }

}

@SuppressWarnings("MagicNumber")
private data class BlockAltitudeDataWithStatics(
    override val origin: Coordinates,
    private val cells: Array<Array<BlockAltitudeCells>>
) : BlockAltitudeData() {

    data class BlockAltitudeCells(
        val groundTile: Short,
        val groundZ: Byte,
        val tiles: Array<StaticCellData>,
    ) {
        fun getCorrectedCellAltitude(originalZ: Int): Int {
            var currentZ = originalZ
            if (originalZ < groundZ) {
                currentZ = groundZ + 1
            }
            if (tiles.isEmpty()) return currentZ
            val topTile = tiles.size - 1
            if (originalZ >= tiles[topTile].z) return tiles[topTile].z + 1 // Original is too high
            var lowTile = 0
            while (lowTile < topTile) {
                if (tiles[lowTile].z > currentZ) break // This tile is over the head, so ignore the rest of tiles
                if (tiles[lowTile].z == originalZ) {
                    // Character is colliding a tile, so it bumps up
                    currentZ++
                }
                if (tiles[lowTile + 1].z > originalZ) {
                    // Character is under next tile, so it goes down
                    currentZ = tiles[lowTile].z + 1
                }
                lowTile++
            }
            return currentZ
        }

        override fun equals(other: Any?): Boolean {
            if (this === other) return true
            if (javaClass != other?.javaClass) return false

            other as BlockAltitudeCells

            if (groundTile != other.groundTile) return false
            if (groundZ != other.groundZ) return false
            if (!tiles.contentEquals(other.tiles)) return false

            return true
        }

        override fun hashCode(): Int {
            var result = groundTile.toInt()
            result = 31 * result + groundZ
            result = 31 * result + tiles.contentHashCode()
            return result
        }

        override fun toString() = toStringByRecursion(
            "groundTile" to groundTile,
            "groundZ" to groundZ,
            "tiles" to tiles
        )
    }

    override fun getCellAttitude(coordinate: Coordinates): Int {
        val difference = coordinate.difference(origin)
        val cell = cells[difference.x][difference.y]
        return cell.getCorrectedCellAltitude(coordinate.z)
    }

    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false

        other as BlockAltitudeDataWithStatics

        if (origin != other.origin) return false
        if (!cells.contentDeepEquals(other.cells)) return false

        return true
    }

    override fun hashCode(): Int {
        var result = origin.hashCode()
        result = 31 * result + cells.contentDeepHashCode()
        return result
    }

    override fun toString() = this.toStringByRecursion(
        "origin" to origin,
        "cells" to cells
    )

    companion object {
        fun create(
            origin: Coordinates,
            mapValues: Array<Array<Pair<Short, Byte>>>,
            staticCells: List<StaticCellData>
        ): BlockAltitudeDataWithStatics {
            val cells = Array(8) { dx ->
                Array(8) { dy ->
                    val groundZ = mapValues[dx][dy].second
                    BlockAltitudeCells(
                        groundTile = mapValues[dx][dy].first,
                        groundZ = groundZ,
                        tiles = staticCells
                            .filter { it.x == dx && it.y == dy && it.z >= groundZ}
                            .sortedBy { it.z }
                            .toTypedArray()
                    )


                }
            }
            return BlockAltitudeDataWithStatics(
                origin = origin,
                cells = cells
            )
        }
    }

}
