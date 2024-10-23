package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.mul.BlockAltitudeData
import dev.cavefish.yamul.backend.game.controller.domain.mul.StaticCellData
import dev.cavefish.yamul.backend.game.controller.infra.mul.MulMapBlockRepository
import dev.cavefish.yamul.backend.game.controller.infra.mul.MulTileDataRepository
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
    override val mulTileDataRepository: MulTileDataRepository,
) : MulMapBlockRepository(), AutoCloseable {


    override fun getBlockAltitudeData(position: Coordinates): BlockAltitudeData {
        val origin = position.toBlockOrigin()
        val blockId = getBlockId(origin)
        val blockAltitudeData = BlockAltitudeData.create(
            origin = origin,
            mapValues = getAltitudeData(origin.mapId, blockId),
            staticCells = getStaticCells(origin.mapId, blockId),
        )
        return blockAltitudeData
    }

    // TODO add cache
    private fun getStaticCells(mapId: Int, blockId: Int): List<StaticCellData> {
        val staticFile = staticsFiles.computeIfAbsent(mapId) {
            multimaFileRepository.getReaderFor(mapProperties[mapId].staticsFile)
        }
        val buffer = staticFile.getBuffer(blockId.toLong())?.order(ByteOrder.LITTLE_ENDIAN)
            ?: return emptyList()

        val result = ArrayList<StaticCellData>()
        while (buffer.hasRemaining()) {
            result.add(
                StaticCellData(
                    objectId = buffer.getShort().toUInt().toInt(),
                    x = buffer.get().toUInt().toInt(),
                    y = buffer.get().toUInt().toInt(),
                    z = buffer.get().toUInt().toInt()
                )
            )
            buffer.getShort()
        }
        return result
    }

    // TODO add cache
    private fun getAltitudeData(mapId: Int, offset: Int): Array<Array<Pair<Short, Byte>>> {
        val zeroShort: Short = 0
        val result: Array<Array<Pair<Short, Byte>>> = Array(BLOCK_WIDTH) { Array(BLOCK_WIDTH) { zeroShort to 0 } }
        val blockBytes = getMapBlockBuffer(mapId, offset)?.order(ByteOrder.LITTLE_ENDIAN)
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

    private fun getMapBlockBuffer(mapFileId: Int, offset: Int): ByteBuffer? {
        val mapFile = mapFiles.computeIfAbsent(mapFileId) {
            return@computeIfAbsent multimaFileRepository.getReaderFor(mapProperties[mapFileId].mapFile)
        }

        val rawDataPosition = (offset * BLOCK_DATA_SIZE)
        return mapFile.getBuffer(rawDataPosition, BLOCK_DATA_SIZE)
    }

    override fun close() {
        Logger.debug("Closing MUL files")
        mapFiles.values.forEach { it.close() }
        mapFiles.clear()
        staticsFiles.values.forEach { it.close() }
        staticsFiles.clear()
    }

    companion object {
        private val mapFiles = ConcurrentHashMap<Int, MultimaFileReader>()
        private val staticsFiles = ConcurrentHashMap<Int, MultimaFileReader>()
    }
}