package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.backend.game.controller.domain.mul.LandTileData
import dev.cavefish.yamul.backend.game.controller.domain.mul.StaticTileData
import dev.cavefish.yamul.backend.game.controller.infra.mul.MulTileDataRepository
import org.springframework.beans.factory.DisposableBean
import org.springframework.stereotype.Repository
import java.io.RandomAccessFile
import java.nio.ByteOrder
import java.nio.channels.FileChannel

private const val LAND_TILE_DATA_SIZE = 30L

private const val LAND_TILE_GROUP_HEADER_SIZE = 4L

private const val LAND_TILE_GROUP_DATA_SIZE = LAND_TILE_GROUP_HEADER_SIZE + 32 * LAND_TILE_DATA_SIZE

private const val STATIC_TILE_GROUP_OFFSET = 512 * LAND_TILE_GROUP_DATA_SIZE

private const val STATIC_TILE_DATA_SIZE = 41L
private const val STATIC_TILE_GROUP_HEADER = 4L
private const val STATIC_TILE_GROUP_SIZE = STATIC_TILE_GROUP_HEADER + 32 * STATIC_TILE_DATA_SIZE

@Repository
@SuppressWarnings("MagicNumber")
class LocalMulTileDataRepository : MulTileDataRepository, DisposableBean {

    override fun getLandTileData(id: Int): LandTileData? {
        assert(id < 0x4000)
        val tileGroup = id ushr 5
        val subGroupId = id and 0x1F

        val position =
            LAND_TILE_GROUP_HEADER_SIZE + tileGroup * LAND_TILE_GROUP_DATA_SIZE + subGroupId * LAND_TILE_DATA_SIZE
        val byteBuffer =
            channel.map(FileChannel.MapMode.READ_ONLY, position, LAND_TILE_DATA_SIZE).order(ByteOrder.LITTLE_ENDIAN)

        return LandTileData(
            flags = byteBuffer.long,
            textureId = byteBuffer.short.toInt(),
            name = MappedByteBufferHelper.readString(byteBuffer, 20),
        )
    }

    override fun getStaticTileData(id: Int): StaticTileData? {
        val tileGroup = id ushr 5
        val subGroupId = id and 0x1F

        val position =
            STATIC_TILE_GROUP_OFFSET +
                    STATIC_TILE_GROUP_HEADER +
                    tileGroup * STATIC_TILE_GROUP_SIZE +
                    subGroupId * STATIC_TILE_DATA_SIZE
        val byteBuffer =
            channel.map(FileChannel.MapMode.READ_ONLY, position, STATIC_TILE_DATA_SIZE).order(ByteOrder.LITTLE_ENDIAN)

        return StaticTileData(
            id = id,
            flags = byteBuffer.long,
            weight = byteBuffer.get().toUByte(),
            layer = byteBuffer.get().toUByte(),
            count = byteBuffer.int,
            animId = byteBuffer.short,
            hue = byteBuffer.short,
            lightIndex = byteBuffer.short,
            height = byteBuffer.get().toUByte(),
            name = MappedByteBufferHelper.readString(byteBuffer, 20),
        )
    }

    override fun destroy() {
        channel.close()
        file.close()
    }

    companion object {
        private val channel: FileChannel
        private val file: RandomAccessFile

        init {
            val fileLocation = LocalMulFileLocation.getFileLocation("tiledata.mul")!!
            file = RandomAccessFile(fileLocation, "r")
            channel = file.channel!!
        }
    }
}