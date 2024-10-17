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
    val multimaFileRepository: MultimaFileRepository,
) : MulMapBlockRepository(), AutoCloseable {

    override fun getBlockAltitudeData(position: Coordinates): BlockAltitudeData {
        val origin = position.toBlockOrigin()
        val blockAltitudeData = BlockAltitudeData(origin, getAltitudeData(origin.mapId, getBlockId(origin)))
        Logger.debug(blockAltitudeData.toString())
        return blockAltitudeData
    }

    // TODO add cache
    private fun getAltitudeData(mapId: Int, offset: Int): Array<Array<Pair<Int, Int>>> {
        val result = Array(BLOCK_WIDTH) { Array(BLOCK_WIDTH) { 0 to 0 } }
        val blockBytes = getBlockBytes(mapId, offset)
        if (blockBytes == null) {
            Logger.error("Block $offset of map $mapId is out-of-bounds")
            return result
        }
        for (x in 0..<BLOCK_WIDTH) {
            for (y in 0..<BLOCK_WIDTH) {
                val inBlockOffset = 3 * (x + BLOCK_WIDTH * y)
                val tileId =
                    ((blockBytes[inBlockOffset].toInt() shl 8) or blockBytes[1 + inBlockOffset].toInt()) and 0xFFFF
                val z = blockBytes[2 + inBlockOffset].toInt()
                result[x][y] = tileId to z
            }
        }
        return result
    }

    @SuppressWarnings("TooGenericExceptionThrown")
    private fun getBlockBytes(mapFile: Int, offset: Int): ByteArray? {
        val file = files.computeIfAbsent(mapFile) {
            return@computeIfAbsent multimaFileRepository.getReaderFor(when (mapFile) {
                0 -> MultimaFileRepository.MulFile.Map0
                1 -> MultimaFileRepository.MulFile.Map1
                2 -> MultimaFileRepository.MulFile.Map2
                3 -> MultimaFileRepository.MulFile.Map3
                4 -> MultimaFileRepository.MulFile.Map4
                5 -> MultimaFileRepository.MulFile.Map5
                else -> throw Exception("Unknown mapFile=$mapFile")
            })
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