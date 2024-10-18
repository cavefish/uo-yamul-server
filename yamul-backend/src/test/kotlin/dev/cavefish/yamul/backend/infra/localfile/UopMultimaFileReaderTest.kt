package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.IntegrationTest
import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.ValueSource
import org.springframework.beans.factory.annotation.Autowired
import kotlin.test.Test

class UopMultimaFileReaderTest : IntegrationTest() {

    @Autowired
    lateinit var multimaFileRepository: MultimaFileRepository

    private fun createReader(mapId: Int) =
        (multimaFileRepository.getReaderFor(MulMapHelper.mapProperties[mapId].mapFile) as UopMultimaFileReader)

    @Test
    fun readMap0FirstBlock() {
        createReader(0).use { reader ->
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

    @ParameterizedTest
    @ValueSource(ints = [0])
    fun readMap0LastBlock(mapId: Int) {
        val map = MulMapHelper.mapProperties[mapId]
        softly.assertThat(map.id).isEqualTo(mapId)
        createReader(map.id).use { reader ->
            val lastBlock0 = 196L * (map.height * map.width) / 64
            val result = reader.getBytes(lastBlock0, 196)
            softly.assertThat(result).isNotNull().hasSize(196)
        }
    }

    @Test
    fun readMap0LastBlockInExcess() {
        createReader(0).use { reader ->
            val lastBlock0 = 196L * (7168 * 4096) / 64
            val result = reader.getBytes(lastBlock0, 196 + 1)
            softly.assertThat(result).isNotNull().hasSize(196 + 1)
            softly.assertThat(result?.get(196) ?: 0).isEqualTo(0)
        }
    }

    @Test
    fun readMap0AfterLastBlock() {
        createReader(0).use { reader ->
            val result = reader.getBytes(89_915_588L + 1L, 1)
            softly.assertThat(result).isNull()
        }
    }

}