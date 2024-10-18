package dev.cavefish.yamul.backend.infra.localfile

import com.google.common.annotations.VisibleForTesting
import dev.cavefish.yamul.backend.Constants
import java.io.RandomAccessFile
import java.nio.ByteBuffer
import java.nio.ByteOrder
import java.nio.channels.FileChannel
import java.util.concurrent.locks.ReentrantLock
import kotlin.collections.HashMap
import kotlin.concurrent.withLock

private const val FILE_HEADER_SIZE = 20L

private const val TABLE_HEADER_SIZE = 12L

private const val ENTRY_HEADER_SIZE = 34L

@SuppressWarnings("TooManyFunctions")
class UopMultimaFileReader(
    private val filenames: List<String>,
    private val bytesPerSubFile: Long,
    private val maxSubFiles: Int,
    private val subFilenameGenerator: (Int) -> String
) :
    MultimaFileReader {

    private val channels = ArrayList<FileChannel>()
    private val files = ArrayList<RandomAccessFile>()

    @VisibleForTesting
    val subFiles = HashMap<ULong, UopFileEntry>()

    @VisibleForTesting
    val subFileHashes = HashMap<Int, ULong>()

    private val mutex = ReentrantLock()

    init {
        filenames.forEach {
            val fileLocation = LocalMulFileLocation.getFileLocation(it)
            val file = RandomAccessFile(fileLocation, "r")
            files.add(file)
            val channel = file.channel
            channels.add(channel)
            readAllTables(channel)
        }
        generateLookupTable(maxSubFiles)
    }

    private fun generateLookupTable(maxID: Int) {
        for (i in 0..<maxID) {
            val subFileName = this.subFilenameGenerator(i)
            val hash = calculateHash(subFileName)
            assert(this.subFiles.containsKey(hash)) { "Missing subFile: $subFileName" }
            this.subFileHashes[i] = hash
        }
    }

    private fun readAllTables(channel: FileChannel) {
        val header = readUopHeader(channel)
        var nextTablePosition = header.firstTablePosition
        while (nextTablePosition > 0) {
            val table = loadTableAtOffset(channel, nextTablePosition) ?: return
            nextTablePosition = table.nextTablePosition
        }
    }

    private fun readUopHeader(channel: FileChannel): UopFileHeader {
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
            subFiles.clear()
            subFileHashes.clear()
            channels.forEach { it.close() }
            channels.clear()
            files.forEach { it.close() }
            files.clear()
        }
    }

    override fun toString(): String {
        return "UopFileReader($filenames)"
    }

    override fun getBytes(offset: Long, size: Int): ByteArray? {
        val result = ByteArray(size)
        var idx = 0
        while (idx < size) {
            val table = getEntryFor(offset + idx) ?: return if (idx == 0) null else result
            val addedBytes = fillBuffer(result, idx, offset + idx, size - idx, table)
            if (addedBytes==0) break
            idx += addedBytes
        }
        return result
    }

    override fun getBuffer(offset: Long, size: Long): ByteBuffer? {
        val bytes = getBytes(offset, size.toInt()) ?: return null
        return ByteBuffer.wrap(bytes)
    }

    private fun getEntryFor(dataOffset: Long): UopFileEntry? {
        val subFileId = (dataOffset / this.bytesPerSubFile).toInt()
        val hash = this.subFileHashes[subFileId]
        val uopFileEntry = this.subFiles[hash] ?: return null
        if (dataOffset % this.bytesPerSubFile > uopFileEntry.size) return null
        return uopFileEntry
    }


    private fun fillBuffer(target: ByteArray, targetOffset: Int, offset: Long, size: Int, entry: UopFileEntry): Int {
        var idx = 0
        val byteBuffer = entry.getMappedByteBuffer()
        val bufferInnerOffset = (offset % this.bytesPerSubFile).toInt()
        byteBuffer.position(bufferInnerOffset)
        while (idx < size) {
            if (bufferInnerOffset + idx >= byteBuffer.limit()) break
            target[idx + targetOffset] = byteBuffer.get()
            idx++
        }
        return idx
    }

    private fun loadTableAtOffset(channel: FileChannel, tablePosition: Long): UopFileTable? {
        if (tablePosition <= 0L) return null
        val headerByteBuffer = channel.map(FileChannel.MapMode.READ_ONLY, tablePosition, TABLE_HEADER_SIZE)
            .order(ByteOrder.LITTLE_ENDIAN)
        val numberOfEntries = headerByteBuffer.getInt()
        val nextTablePosition = headerByteBuffer.getLong()

        for (i in 0..<numberOfEntries) {
            val entry = loadEntry(channel, tablePosition + TABLE_HEADER_SIZE + i * ENTRY_HEADER_SIZE)
            if (entry.size == 0) continue
            subFiles[entry.filenameHash] = entry
        }

        val table = UopFileTable(
            nextTablePosition = nextTablePosition,
            numberOfEntries = numberOfEntries,
        )
        return table
    }

    private fun loadEntry(channel: FileChannel, entryOffset: Long): UopFileEntry {
        val byteBuffer = channel.map(FileChannel.MapMode.READ_ONLY, entryOffset, ENTRY_HEADER_SIZE)
            .order(ByteOrder.LITTLE_ENDIAN)

        val result = UopFileEntry(
            fileChannel = channel,
            compressedDataOffset = byteBuffer.long,
            headerLength = byteBuffer.int,
            size = byteBuffer.int,
            decompressedSize = byteBuffer.int,
            filenameHash = byteBuffer.long.toULong(),
            hash = byteBuffer.int,
            compressionMethod = byteBuffer.short,
        )
        assert(result.compressionMethod == 0.toShort())
        return result
    }

    companion object {

        @SuppressWarnings("MagicNumber", "CyclomaticComplexMethod", "LongMethod")
        @VisibleForTesting
        fun calculateHash(s: String): ULong {
            var eax = 0u
            val ecx: UInt
            var edx: UInt
            var ebx: UInt
            var esi: UInt
            var edi: UInt
            ebx = (s.length + 0xDEADBEEF).toUInt()
            edi = ebx
            esi = ebx
            var i = 0

            while (i + 12 < s.length) {
                edi = ((s[i + 7].code.toUInt() shl 24) or
                        (s[i + 6].code.toUInt() shl 16) or
                        (s[i + 5].code.toUInt() shl 8) or
                        s[i + 4].code.toUInt()) + edi
                esi = ((s[i + 11].code.toUInt() shl 24) or
                        (s[i + 10].code.toUInt() shl 16) or
                        (s[i + 9].code.toUInt() shl 8) or
                        s[i + 8].code.toUInt()) + esi
                edx = ((s[i + 3].code.toUInt() shl 24) or
                        (s[i + 2].code.toUInt() shl 16) or
                        (s[i + 1].code.toUInt() shl 8) or
                        s[i].code.toUInt()) - esi
                edx = (edx + ebx) xor (esi shr 28) xor (esi shl 4)
                esi += edi
                edi = (edi - edx) xor (edx shr 26) xor (edx shl 6)
                edx += esi
                esi = (esi - edi) xor (edi shr 24) xor (edi shl 8)
                edi += edx
                ebx = (edx - esi) xor (esi shr 16) xor (esi shl 16)
                esi += edi
                edi = (edi - ebx) xor (ebx shr 13) xor (ebx shl 19)
                ebx += esi
                esi = (esi - edi) xor (edi shr 28) xor (edi shl 4)
                edi += ebx
                i += 12
            }

            val rest = s.length - i
            assert(rest in 0..12)
            if (rest > 0) {
                if (rest >= 12) esi += (s[i + 11].code.toUInt() shl 24)
                if (rest >= 11) esi += (s[i + 10].code.toUInt() shl 16)
                if (rest >= 10) esi += (s[i + 9].code.toUInt() shl 8)
                if (rest >= 9) esi += s[i + 8].code.toUInt()
                if (rest >= 8) edi += (s[i + 7].code.toUInt() shl 24)
                if (rest >= 7) edi += (s[i + 6].code.toUInt() shl 16)
                if (rest >= 6) edi += (s[i + 5].code.toUInt() shl 8)
                if (rest >= 5) edi += s[i + 4].code.toUInt()
                if (rest >= 4) ebx += (s[i + 3].code.toUInt() shl 24)
                if (rest >= 3) ebx += (s[i + 2].code.toUInt() shl 16)
                if (rest >= 2) ebx += (s[i + 1].code.toUInt() shl 8)
                ebx += s[i].code.toUInt()

                esi = (esi xor edi) - ((edi shr 18) xor (edi shl 14))
                ecx = (esi xor ebx) - ((esi shr 21) xor (esi shl 11))
                edi = (edi xor ecx) - ((ecx shr 7) xor (ecx shl 25))
                esi = (esi xor edi) - ((edi shr 16) xor (edi shl 16))
                edx = (esi xor ecx) - ((esi shr 28) xor (esi shl 4))
                edi = (edi xor edx) - ((edx shr 18) xor (edx shl 14))
                eax = (esi xor edi) - ((edi shr 8) xor (edi shl 24))

                return (edi.toULong() shl 32) or eax.toULong()
            }

            return (esi.toULong() shl 32) or eax.toULong()
        }
    }

    data class UopFileHeader(
        val header: Int,
        val version: Int,
        val timestamp: Int,
        val firstTablePosition: Long
    ) {
        init {
            assert(header >= 0)
            assert(version >= 0)
            assert(firstTablePosition > 0)
        }
    }

    data class UopFileTable(
        val numberOfEntries: Int,
        val nextTablePosition: Long,
    ) {

        init {
            assert(numberOfEntries >= 0)
            assert(nextTablePosition >= 0)
        }
    }

    data class UopFileEntry(
        val fileChannel: FileChannel,
        val compressedDataOffset: Long,
        val headerLength: Int,
        val size: Int,
        val decompressedSize: Int,
        val filenameHash: ULong,
        val hash: Int,
        val compressionMethod: Short
    ) {


        fun getMappedByteBuffer(): ByteBuffer {
            val realDataPosition = compressedDataOffset + headerLength
            return fileChannel.map(FileChannel.MapMode.READ_ONLY, realDataPosition, size.toLong())
                .order(ByteOrder.LITTLE_ENDIAN)
        }

        override fun toString(): String {
            return "UopFileEntry(" +
                    "compressedDataOffset=$compressedDataOffset," +
                    " headerLength=$headerLength," +
                    " size=$size," +
                    " decompressedSize=$decompressedSize," +
                    " filenameHash=${Constants.toHexFormat(filenameHash)}," +
                    " hash=${Constants.toHexFormat(hash)}," +
                    " compressionMethod=$compressionMethod" +
                    ")"
        }
    }
}