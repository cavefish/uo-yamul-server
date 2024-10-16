package dev.cavefish.yamul.backend.infra.localfile

interface MultimaFileReader : AutoCloseable {
    fun getBytes(offset: Long, size: Int): ByteArray?
    fun getSize(): Long
}