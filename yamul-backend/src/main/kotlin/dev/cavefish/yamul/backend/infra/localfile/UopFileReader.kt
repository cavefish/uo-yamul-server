package dev.cavefish.yamul.backend.infra.localfile

import java.io.RandomAccessFile
import java.nio.MappedByteBuffer
import java.nio.channels.FileChannel
import java.util.*
import java.util.concurrent.locks.ReentrantLock
import kotlin.concurrent.withLock

private const val FILE_HEADER_SIZE = 20L

private const val TABLE_HEADER_SIZE = 12L

private const val ENTRY_HEADER_SIZE = 34L

@SuppressWarnings("TooManyFunctions")
class UopFileReader(private val filename: String) : AutoCloseable {

    private val channel: FileChannel
    private val file: RandomAccessFile
    private val header: UopFileHeader
    private val tables: TreeMap<Long, UopFileTable> = TreeMap()
    private val mutex = ReentrantLock()

    init {
        val fileLocation = LocalMulFileLocation.getFileLocation("$filename.uop")
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
        val header = readInt(buffer)
        val version = readInt(buffer)
        val timestamp = readInt(buffer)
        val firstTablePosition = readLong(buffer)
        return UopFileHeader(
            header = header,
            version = version,
            timestamp = timestamp,
            firstTablePosition = firstTablePosition
        )
    }

    override fun close() {
        channel.close()
        file.close()
    }

    override fun toString(): String {
        return "UopFileReader($filename)"
    }

    fun getBytes(offset: Long, size: Int): ByteArray? {
        val result = ByteArray(size)
        var idx = 0
        while (idx < size) {
            val table = getTableFor(offset + idx) ?: return if (idx == 0) null else result
            idx += fillBuffer(result, idx, offset + idx, size - idx, table)
        }
        return result
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
        val numberOfEntries = readInt(headerByteBuffer)
        val nextTablePosition = readLong(headerByteBuffer)
        val entries = TreeMap<Long, UopFileEntry>()

        var nextEntryOffset = dataOffset
        for (i in 0..<numberOfEntries) {
            val entry = loadEntry(nextEntryOffset, tablePosition + TABLE_HEADER_SIZE + i * ENTRY_HEADER_SIZE)
            entries.put(nextEntryOffset, entry)
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

        val compressedDataOffset = readLong(byteBuffer)
        val headerLength = readInt(byteBuffer)
        val size = readInt(byteBuffer)
        val decompressedSize = readInt(byteBuffer)
        val filenameHash = readLong(byteBuffer)
        val hash = readInt(byteBuffer)
        val compressionMethod = readShort(byteBuffer)

        assert(compressionMethod.toInt() == 0)

        return UopFileEntry(
            decompressedDataOffset = dataOffset,
            compressedDataOffset = compressedDataOffset,
            headerLength = headerLength,
            size = size,
            decompressedSize = decompressedSize,
            filenameHash = filenameHash,
            hash = hash,
            compressionMethod = compressionMethod,
        )
    }

    @SuppressWarnings("MagicNumber")
    private fun readShort(buffer: MappedByteBuffer): Short {
        val byte1 = buffer.get().toInt()
        val byte0 = buffer.get().toInt()
        return (byte0 shl 8 or byte1).toShort()
    }

    @SuppressWarnings("MagicNumber")
    private fun readInt(buffer: MappedByteBuffer): Int {
        val byte3 = buffer.get().toInt() and 0xFF
        val byte2 = buffer.get().toInt() and 0xFF
        val byte1 = buffer.get().toInt() and 0xFF
        val byte0 = buffer.get().toInt() and 0xFF
        return (byte0 shl 24) or (byte1 shl 16) or (byte2 shl 8) or (byte3)
    }

    @SuppressWarnings("MagicNumber")
    private fun readLong(byteBuffer: MappedByteBuffer): Long {
        val byte47 = readInt(byteBuffer).toLong()
        val byte03 = readInt(byteBuffer).toLong()
        return byte03 shl 32 or byte47
    }

    data class UopFileHeader(
        val header: Int,
        val version: Int,
        val timestamp: Int,
        val firstTablePosition: Long
    )

    data class UopFileTable(
        val numberOfEntries: Int,
        val nextTablePosition: Long,
        val entries: TreeMap<Long, UopFileEntry>,
    ) {
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