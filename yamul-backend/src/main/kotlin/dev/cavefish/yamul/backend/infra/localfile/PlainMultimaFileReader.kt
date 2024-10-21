package dev.cavefish.yamul.backend.infra.localfile

import org.tinylog.kotlin.Logger
import java.io.RandomAccessFile
import java.nio.ByteBuffer
import java.nio.channels.FileChannel

class PlainMultimaFileReader(private val filename: String) : MultimaFileReader {

    private val channel: FileChannel
    private val file: RandomAccessFile

    init {
        val fileLocation = LocalMulFileLocation.getFileLocation(filename)
        assert(fileLocation != null) {"Cannot access $filename"}
        file = RandomAccessFile(fileLocation, "r")
        channel = file.channel
    }

    override fun getBytes(offset: Long, size: Int): ByteArray? {
        val map = channel.map(FileChannel.MapMode.READ_ONLY, offset, size.toLong())
        val result = ByteArray(size)
        map.get(result)
        return result
    }

    override fun getBuffer(offset: Long, size: Long?): ByteBuffer? {
        assert(size != null && size >=0)
        return channel.map(FileChannel.MapMode.READ_ONLY, offset, size!!)
    }

    override fun close() {
        channel.close()
        file.close()
    }
}