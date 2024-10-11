package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.backend.utils.StringUtils.trimZeros
import java.nio.ByteBuffer

object MappedByteBufferHelper {
    fun readString(byteBuffer: ByteBuffer, size: Int): String {
        val bytes = ByteArray(size)
        byteBuffer.get(bytes)
        return trimZeros(String(bytes, Charsets.UTF_8))
    }

}