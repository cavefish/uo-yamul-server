package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.mul.BlockAltitudeData
import dev.cavefish.yamul.backend.game.controller.infra.mul.MulMapBlockRepository
import dev.cavefish.yamul.backend.infra.localfile.MulMapHelper.getBlockId
import dev.cavefish.yamul.backend.infra.localfile.MulMapHelper.mapProperties
import org.springframework.stereotype.Repository
import org.tinylog.kotlin.Logger
import java.nio.ByteBuffer
import java.nio.ByteOrder
import java.util.concurrent.ConcurrentHashMap

private const val BLOCK_WIDTH = 8

private const val BLOCK_DATA_SIZE: Long = 4L + 3 * BLOCK_WIDTH * BLOCK_WIDTH


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
    private fun getAltitudeData(mapId: Int, offset: Int): Array<Array<Pair<Short, Byte>>> {
        val zeroShort: Short = 0
        val result: Array<Array<Pair<Short, Byte>>> = Array(BLOCK_WIDTH) { Array(BLOCK_WIDTH) { zeroShort to 0 } }
        val blockBytes = getBlockBuffer(mapId, offset)?.order(ByteOrder.LITTLE_ENDIAN)
        if (blockBytes == null) {
            Logger.error("Block $offset of map $mapId is out-of-bounds")
            return result
        }
        for (y in 0..<BLOCK_WIDTH) {
            for (x in 0..<BLOCK_WIDTH) {
                val tileId = blockBytes.getShort()
                val z = blockBytes.get()
                result[x][y] = tileId to z
            }
        }
        // val header = blockBytes.getInt()
        return result
    }

    private fun getBlockBuffer(mapFile: Int, offset: Int): ByteBuffer? {
        val file = files.computeIfAbsent(mapFile) {
            return@computeIfAbsent multimaFileRepository.getReaderFor(mapProperties[mapFile].mapFile)
        }

        val rawDataPosition = (offset * BLOCK_DATA_SIZE)
        return file.getBuffer(rawDataPosition, BLOCK_DATA_SIZE)
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