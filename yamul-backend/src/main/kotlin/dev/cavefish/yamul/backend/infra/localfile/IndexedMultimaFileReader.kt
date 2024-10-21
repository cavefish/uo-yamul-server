package dev.cavefish.yamul.backend.infra.localfile

import org.tinylog.kotlin.Logger
import java.nio.ByteBuffer
import java.nio.ByteOrder

private const val IDX_BLOCK_SIZE = 12L

class IndexedMultimaFileReader(
    idxFilename: String,
    baseFilename: String
) : MultimaFileReader {

    private val idxReader = PlainMultimaFileReader(idxFilename)
    private val baseReader = PlainMultimaFileReader(baseFilename)

    override fun getBytes(offset: Long, size: Int): ByteArray? {
        val indexItem = getIndexItem(offset) ?: return null
        return baseReader.getBytes(indexItem.first, indexItem.second.toInt())
    }

    override fun getBuffer(offset: Long, size: Long?): ByteBuffer? {
        val indexItem = getIndexItem(offset) ?: return null
        return baseReader.getBuffer(indexItem.first, indexItem.second)
    }

    private fun getIndexItem(offset: Long): Pair<Long, Long>? {
        val buffer =
            idxReader.getBuffer(offset * IDX_BLOCK_SIZE, IDX_BLOCK_SIZE)
                ?.order(ByteOrder.LITTLE_ENDIAN) ?: return null
        val start = buffer.getInt().toLong()
        val length = buffer.getInt().toLong()
        if (start < 0) return null
        if (length <= 0) return null
        return start to length
    }

    override fun close() {
        idxReader.close()
        baseReader.close()
    }

}
