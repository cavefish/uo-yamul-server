package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.mul.BlockAltitudeData
import dev.cavefish.yamul.backend.game.controller.infra.mul.MulMapBlockRepository
import dev.cavefish.yamul.backend.infra.localfile.MulBlockHelper.getBlockId
import org.springframework.stereotype.Repository
import org.tinylog.kotlin.Logger
import java.util.concurrent.ConcurrentHashMap

private const val BLOCK_WIDTH = 8

private const val BLOCK_DATA_SIZE = 4 + 3 * BLOCK_WIDTH * BLOCK_WIDTH


@Repository
@Suppress("MagicNumber")
class LocalMulMapBlockRepository(
    val multimaFileRepository: MultimaFileRepository
) : MulMapBlockRepository(), AutoCloseable {

//    override fun correctPositionAltitude(cell: Coordinates): Coordinates {
//        val correctedCoordinates = super.correctPositionAltitude(cell)
//        Logger.warn("Bugged behaviour. Returning a z value of 127. Result from mapFile: $correctedCoordinates")
//        return cell.copy(z=127)
//    }

    override fun getBlockAltitudeData(position: Coordinates): BlockAltitudeData {
        val origin = position.toBlockOrigin()
        val blockAltitudeData = BlockAltitudeData(origin, getAltitudeData(origin.mapId, getBlockId(origin)))
        return blockAltitudeData
    }

    // TODO add cache
    private fun getAltitudeData(mapId: Int, offset: Int): Array<Array<Int>> {
        val result = Array(BLOCK_WIDTH) { Array(BLOCK_WIDTH) { 0 } }
        val blockBytes = getBlockBytes(mapId, offset)
        if (blockBytes == null) {
            Logger.error("Block $offset of map $mapId is out-of-bounds")
            return result
        }
        for (x in 0..<BLOCK_WIDTH) {
            for (y in 0..<BLOCK_WIDTH) {
                result[x][y] = blockBytes[4 + 2 + 3 * (x + BLOCK_WIDTH * y)].toInt()
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
                0, 1, 2, 5 -> multimaFileRepository.getReaderFor("map${mapFile}LegacyMUL.uop")
                else -> multimaFileRepository.getReaderFor("map${mapFile}LegacyMUL.uop")
            }
        }

        val rawDataPosition = (offset * BLOCK_DATA_SIZE).toLong()
        val rawData = file.getBytes(rawDataPosition, BLOCK_DATA_SIZE)
        return rawData
    }

    override fun close() {
        Logger.debug("Closing MUL files")
        files.values.forEach { it.close() }
        files.clear()
    }

    companion object {
        private val files = ConcurrentHashMap<Int, MultimaFileReader>()
    }
}