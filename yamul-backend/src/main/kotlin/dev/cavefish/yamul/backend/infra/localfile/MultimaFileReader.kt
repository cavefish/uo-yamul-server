package dev.cavefish.yamul.backend.infra.localfile

import java.nio.ByteBuffer

interface MultimaFileReader : AutoCloseable {
    fun getBytes(offset: Long, size: Int): ByteArray?
    fun getBuffer(offset: Long, size: Long? = null): ByteBuffer?
}