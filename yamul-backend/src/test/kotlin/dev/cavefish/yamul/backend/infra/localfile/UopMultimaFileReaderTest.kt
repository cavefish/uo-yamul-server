package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.IntegrationTest
import kotlin.test.Test

class UopMultimaFileReaderTest : IntegrationTest() {

    @Test
    fun readMap0FirstBlock() {
        UopMultimaFileReader("map0xLegacyMUL.uop").use { reader ->
            val expectedFirstBlock = byteArrayOf(0, 0, 0, 0, 0xA8.toByte(), 0, 0xFB.toByte(), 0xA8.toByte())

            // Before Reading
            softly.assertThat(reader.header).isEqualTo(UopMultimaFileReader.UopFileHeader(
                header = 0x0050594D,
                version = 0x0004,
                timestamp = 0xFD23EC43.toInt(),
                firstTablePosition = 0x28L
            ))
            softly.assertThat(reader.tables).describedAs("All tables in the file are pre-loaded").hasSize(2)
            softly.assertThat(reader.tables.lastEntry().value).extracting { it.nextTablePosition }.isEqualTo(0L)

            // Asking for decompressed file size
            softly.assertThat(reader.getSize()).isEqualTo(89_915_392L)

            // After reading the first block
            val firstBlock = reader.getBytes(0, expectedFirstBlock.size)
            softly.assertThat(firstBlock).isEqualTo(expectedFirstBlock)
        }
    }

    @Test
    fun readMap0LastBlock() {
        UopMultimaFileReader("map0xLegacyMUL.uop").use { reader ->
            val lastBlock0 = 196L*(7168*4096)/64
            val result = reader.getBytes(lastBlock0, 196)
            softly.assertThat(result).isNotNull().hasSize(196)
        }
    }

    @Test
    fun readMap0AfterLastBlock() {
        UopMultimaFileReader("map0xLegacyMUL.uop").use { reader ->
            val result = reader.getBytes(89_915_588L + 1L, 1)
            softly.assertThat(result).isNull()
        }
    }

}