package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.mul.BlockAltitudeData
import dev.cavefish.yamul.backend.game.controller.infra.mul.MulBlockAltitudeRepository
import org.springframework.stereotype.Repository
import org.tinylog.kotlin.Logger
import java.util.concurrent.ConcurrentHashMap
import javax.annotation.PreDestroy

private const val BLOCK_WIDTH = 8

private const val BLOCK_DATA_SIZE = 4 + 3 * BLOCK_WIDTH * BLOCK_WIDTH


@Repository
@Suppress("MagicNumber")
class LocalMulBlockAltitudeRepository : MulBlockAltitudeRepository() {

    override fun correctPositionAltitude(cell: Coordinates): Coordinates {
        val correctedCoordinates = super.correctPositionAltitude(cell)
        Logger.warn("Bugged behaviour. Returning a z value of 127. Result from mapFile: $correctedCoordinates")
        return cell.copy(z=127)
    }

    override fun getBlockAltitudeData(position: Coordinates): BlockAltitudeData {
        val origin = position.toBlockOrigin()
        val blockAltitudeData = BlockAltitudeData(origin, getAltitudeData(origin.mapId, getBlockFileOffset(origin)))
        return blockAltitudeData
    }

    private fun getBlockFileOffset(origin: Coordinates): Int {
        val blockX = origin.x ushr 3
        val blockY = origin.y ushr 3
        return blockX * mapWidths[origin.mapId]!! + blockY
    }

    // TODO add cache
    private fun getAltitudeData(mapId: Int, offset: Int): Array<Array<Int>> {
        val result = Array(BLOCK_WIDTH) { Array(BLOCK_WIDTH) { 0 } }
        val blockBytes = getBlockBytes(mapId, offset)
        if (blockBytes == null) {
            Logger.debug("Block $offset of map $mapId is out-of-bounds")
            return result
        }
        for (i in 0..<BLOCK_WIDTH) {
            for (j in 0..<BLOCK_WIDTH) {
                result[i][j] = blockBytes.get(4 + 2 + 3 * (i * BLOCK_WIDTH + j)).toInt()
            }
        }
        Logger.debug(
            "mapId=$mapId, offset=$offset, result:\n ${
                result.joinToString("\n") { it.contentToString() }
            }]"
        )
        return result
    }

    private fun getBlockBytes(mapFile: Int, offset: Int): ByteArray? {
        val file = files.computeIfAbsent(mapFile) {
            return@computeIfAbsent when (mapFile) {
                0, 1, 2, 5 -> UopFileReader("map${mapFile}xLegacyMUL")
                else -> UopFileReader("map${mapFile}LegacyMUL")
            }
        }

        val rawDataPosition = (offset * BLOCK_DATA_SIZE).toLong()
        val rawData = file.getBytes(rawDataPosition, BLOCK_DATA_SIZE)
        return rawData
    }

    companion object {

        @PreDestroy
        fun onDestroy() {
            Logger.debug("Closing MUL files")
            files.values.forEach { it.close() }
        }

        private val files = ConcurrentHashMap<Int, UopFileReader>()
        private val mapWidths = mapOf(
            1 to 7168 / 8,
            2 to 7168 / 8,
            3 to 2304 / 8,
            4 to 2560 / 8,
            5 to 1448 / 8,
            6 to 1280 / 8,
        )
    }
}