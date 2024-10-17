package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.IntegrationTest
import dev.cavefish.yamul.backend.infra.localfile.MultimaFileRepository.MulFile.Map0
import org.springframework.beans.factory.annotation.Autowired
import kotlin.test.Test

class UopMultimaFileReaderTest : IntegrationTest() {

    @Autowired
    lateinit var multimaFileRepository: MultimaFileRepository

    @Test
    fun readMap0FirstBlock() {
        (multimaFileRepository.getReaderFor(Map0) as UopMultimaFileReader).use { reader ->
            val expectedFirstBlock = byteArrayOf(0, 0, 0, 0, 0xA8.toByte(), 0, 0xFB.toByte(), 0xA8.toByte())

            // Before Reading
            softly.assertThat(reader.subFiles)
                .describedAs("All tables in the file are pre-loaded")
                .hasSize(113)

            softly.assertThat(reader.subFileHashes)
                .hasSize(113)

            // After reading the first block
            val firstBlock = reader.getBytes(0, expectedFirstBlock.size)
            softly.assertThat(firstBlock).isEqualTo(expectedFirstBlock)
        }
    }

    @Test
    fun readMap0LastBlock() {
        (multimaFileRepository.getReaderFor(Map0) as UopMultimaFileReader).use { reader ->
            val lastBlock0 = 196L * (7168 * 4096) / 64
            val result = reader.getBytes(lastBlock0, 196)
            softly.assertThat(result).isNotNull().hasSize(196)
        }
    }

    @Test
    fun readMap0LastBlockInExcess() {
        (multimaFileRepository.getReaderFor(Map0) as UopMultimaFileReader).use { reader ->
            val lastBlock0 = 196L * (7168 * 4096) / 64
            val result = reader.getBytes(lastBlock0, 196 + 1)
            softly.assertThat(result).isNotNull().hasSize(196 + 1)
            softly.assertThat(result?.get(196) ?: 0).isEqualTo(0)
        }
    }

    @Test
    fun readMap0AfterLastBlock() {
        (multimaFileRepository.getReaderFor(Map0) as UopMultimaFileReader).use { reader ->
            val result = reader.getBytes(89_915_588L + 1L, 1)
            softly.assertThat(result).isNull()
        }
    }

}