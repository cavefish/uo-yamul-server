package dev.cavefish.yamul.backend.infra.localfile

import com.google.common.annotations.VisibleForTesting
import java.io.RandomAccessFile
import java.nio.ByteOrder
import java.nio.MappedByteBuffer
import java.nio.channels.FileChannel
import java.util.*
import java.util.concurrent.locks.ReentrantLock
import kotlin.concurrent.withLock

private const val FILE_HEADER_SIZE = 20L

private const val TABLE_HEADER_SIZE = 12L

private const val ENTRY_HEADER_SIZE = 34L

@SuppressWarnings("TooManyFunctions")
class UopMultimaFileReader(private val filename: String) : MultimaFileReader {

    private val channel: FileChannel
    private val file: RandomAccessFile
    @VisibleForTesting
    val header: UopFileHeader
    @VisibleForTesting
    val tables: TreeMap<Long, UopFileTable> = TreeMap()
    private val mutex = ReentrantLock()

    init {
        val fileLocation = LocalMulFileLocation.getFileLocation(filename)
        file = RandomAccessFile(fileLocation, "r")
        channel = file.channel
        header = readUopHeader()
        readAllTables()
    }

    private fun readAllTables() {
        var nextTablePosition = header.firstTablePosition
        var dataOffset = 0L
        while (nextTablePosition>0) {
            val table = loadTableAtOffset(dataOffset, nextTablePosition) ?: return
            val lastEntry = table.entries.lastEntry().value ?: return
            dataOffset = lastEntry.decompressedDataOffset + lastEntry.decompressedSize
            nextTablePosition = table.nextTablePosition
        }
    }

    private fun readUopHeader(): UopFileHeader {
        val buffer = channel.map(FileChannel.MapMode.READ_ONLY, 0, FILE_HEADER_SIZE)
            .order(ByteOrder.LITTLE_ENDIAN)
        val uopFileHeader = UopFileHeader(
            header = buffer.int,
            version = buffer.int,
            timestamp = buffer.int,
            firstTablePosition = buffer.long
        )
        return uopFileHeader
    }

    override fun close() {
        mutex.withLock {
            tables.clear()
            channel.close()
            file.close()
        }
    }

    override fun toString(): String {
        return "UopFileReader($filename)"
    }

    override fun getBytes(offset: Long, size: Int): ByteArray? {
        val result = ByteArray(size)
        var idx = 0
        while (idx < size) {
            val table = getTableFor(offset + idx) ?: return if (idx == 0) null else result
            idx += fillBuffer(result, idx, offset + idx, size - idx, table)
        }
        return result
    }

    override fun getSize(): Long {
        return tables.lastEntry().value.entries.lastEntry().value.decompressedDataOffset
    }

    private fun fillBuffer(target: ByteArray, targetOffset: Int, offset: Long, size: Int, table: UopFileTable): Int {
        var idx = 0
        while (idx < size) {
            val entry = table.getEntryFor(offset) ?: return idx
            idx += fillBuffer(target, targetOffset, offset, size, entry)
        }
        return idx
    }

    private fun fillBuffer(target: ByteArray, targetOffset: Int, offset: Long, size: Int, entry: UopFileEntry): Int {
        var idx = 0
        val byteBuffer = entry.getMappedByteBuffer(channel)
        val bufferInnerOffset = (offset - entry.decompressedDataOffset).toInt()
        byteBuffer.position(bufferInnerOffset)
        while (idx < size) {
            if (bufferInnerOffset + idx >= byteBuffer.limit()) break
            target[idx + targetOffset] = byteBuffer.get()
            idx++
        }
        assert(idx > 0)
        return idx
    }

    private fun getTableFor(offset: Long): UopFileTable? = mutex.withLock {
        var floorTable = tables.floorEntry(offset)?.value ?: return null
        if (!floorTable.containsDecompressedOffset(offset)) {
             return null
        }
        return@withLock floorTable
    }

    private fun loadTableAtOffset(dataOffset: Long, tablePosition: Long): UopFileTable? {
        if (tablePosition <= 0L) return null
        val headerByteBuffer = channel.map(FileChannel.MapMode.READ_ONLY, tablePosition, TABLE_HEADER_SIZE)
            .order(ByteOrder.LITTLE_ENDIAN)
        val numberOfEntries = headerByteBuffer.getInt()
        val nextTablePosition = headerByteBuffer.getLong()
        val entries = TreeMap<Long, UopFileEntry>()

        var nextEntryOffset = dataOffset
        for (i in 0..<numberOfEntries) {
            val entry = loadEntry(nextEntryOffset, tablePosition + TABLE_HEADER_SIZE + i * ENTRY_HEADER_SIZE)
            entries[nextEntryOffset] = entry
            nextEntryOffset += entry.decompressedSize
        }

        val table = UopFileTable(
            nextTablePosition = nextTablePosition,
            numberOfEntries = numberOfEntries,
            entries = entries
        )
        tables[dataOffset] = table
        return table
    }

    private fun loadEntry(dataOffset: Long, entryOffset: Long): UopFileEntry {
        val byteBuffer = channel.map(FileChannel.MapMode.READ_ONLY, entryOffset, ENTRY_HEADER_SIZE)
            .order(ByteOrder.LITTLE_ENDIAN)

        val result = UopFileEntry(
            decompressedDataOffset = dataOffset,
            compressedDataOffset = byteBuffer.long,
            headerLength = byteBuffer.int,
            size = byteBuffer.int,
            decompressedSize = byteBuffer.int,
            filenameHash = byteBuffer.long,
            hash = byteBuffer.int,
            compressionMethod = byteBuffer.short,
        )
        assert(result.compressionMethod == 0.toShort())
        return result
    }

    data class UopFileHeader(
        val header: Int,
        val version: Int,
        val timestamp: Int,
        val firstTablePosition: Long
    ) {
        init {
            assert(header>=0)
            assert(version>=0)
            assert(firstTablePosition>0)
        }
    }

    data class UopFileTable(
        val numberOfEntries: Int,
        val nextTablePosition: Long,
        val entries: TreeMap<Long, UopFileEntry>,
    ) {

        init {
            assert(numberOfEntries>=0)
            assert(nextTablePosition>=0)
        }

        fun getEntryFor(offset: Long): UopFileEntry? {
            val entry = entries.floorEntry(offset).value
            if (!entry.containsDecompressedPosition(offset)) return null
            return entry
        }

        fun containsDecompressedOffset(offset: Long): Boolean {
            if (entries.firstEntry().value.decompressedDataOffset > offset) return false
            val last = entries.lastEntry().value
            return last.decompressedDataOffset + last.size >= offset
        }
    }

    data class UopFileEntry(
        val decompressedDataOffset: Long,
        val compressedDataOffset: Long,
        val headerLength: Int,
        val size: Int,
        val decompressedSize: Int,
        val filenameHash: Long,
        val hash: Int,
        val compressionMethod: Short
    ) {
        fun containsDecompressedPosition(offset: Long): Boolean {
            if (decompressedDataOffset > offset) return false
            if (decompressedDataOffset + decompressedSize <= offset) return false
            return true
        }

        fun getMappedByteBuffer(channel: FileChannel): MappedByteBuffer {
            val realDataPosition = compressedDataOffset + headerLength
            return channel.map(FileChannel.MapMode.READ_ONLY, realDataPosition, size.toLong())
        }
    }
}