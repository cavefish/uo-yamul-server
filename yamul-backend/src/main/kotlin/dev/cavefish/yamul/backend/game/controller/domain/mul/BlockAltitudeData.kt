package dev.cavefish.yamul.backend.game.controller.domain.mul

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.infra.mul.MulTileDataRepository
import dev.cavefish.yamul.backend.utils.toStringByRecursion

sealed class BlockAltitudeData {
    abstract val origin: Coordinates
    abstract fun getCellAttitude(
        coordinate: Coordinates,
        bodyHeight: Int,
        tileDataRepository: MulTileDataRepository
    ): Int

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

    override fun getCellAttitude(
        coordinate: Coordinates,
        bodyHeight: Int,
        tileDataRepository: MulTileDataRepository
    ): Int {
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
        fun getCorrectedCellAltitude(originalZ: Int, bodyHeight: Int, tileDataRepository: MulTileDataRepository): Int {
            var currentZ = originalZ
            if (originalZ < groundZ) {
                currentZ = groundZ.toInt()
            }
            if (tiles.isEmpty()) return groundZ + bodyHeight
            val topTile = tiles.size - 1
            if (originalZ >= bottomOfTile(topTile))
                return topOfTile(tileDataRepository, topTile) + bodyHeight // Original is too high
            var lowTile = 0
            while (lowTile < topTile) {
                // This tile is over the head, so ignore the rest of tiles
                if (bottomOfTile(lowTile) > currentZ + bodyHeight) break
                // Character is colliding a tile, so it bumps up
                if (topOfTile(tileDataRepository, lowTile) >= currentZ) {
                    currentZ = topOfTile(tileDataRepository, lowTile)
                }
                lowTile++
            }
            return currentZ
        }

        private fun bottomOfTile(idx: Int): Int {
            return tiles[idx].z
        }

        private fun topOfTile(tileDataRepository: MulTileDataRepository, idx: Int): Int {
            val staticCellData = tiles[idx]
            val tileData = tileDataRepository.getStaticTileData(staticCellData.objectId)!!
            return (tileData.height).toInt() + staticCellData.z
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

    override fun getCellAttitude(
        coordinate: Coordinates,
        bodyHeight: Int,
        tileDataRepository: MulTileDataRepository
    ): Int {
        val difference = coordinate.difference(origin)
        val cell = cells[difference.x][difference.y]
        return cell.getCorrectedCellAltitude(coordinate.z, bodyHeight, tileDataRepository)
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
                            .filter { it.x == dx && it.y == dy && it.z >= groundZ }
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
