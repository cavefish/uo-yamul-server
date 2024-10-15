package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.IntegrationTest

import org.assertj.core.api.Assertions.assertThat
import kotlin.test.Test

class UopMultimaFileReaderTest : IntegrationTest() {

    @Test
    fun readMap0FirstBlock() {
        UopMultimaFileReader("map0xLegacyMUL").use { reader ->
            val expected = ByteArray(4)
            expected[0] = 0
            expected[1] = 0
            expected[2] = 0
            expected[3] = 0

            val result = reader.getBytes(0, expected.size)
            assertThat(result).isEqualTo(expected)
        }
    }

}