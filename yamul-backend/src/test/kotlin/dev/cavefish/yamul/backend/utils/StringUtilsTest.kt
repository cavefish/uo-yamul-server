package dev.cavefish.yamul.backend.utils

import org.assertj.core.api.AssertionsForClassTypes.assertThat
import org.junit.jupiter.api.TestInstance
import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.Arguments
import org.junit.jupiter.params.provider.MethodSource
import java.util.stream.Stream

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class StringUtilsTest {
    @ParameterizedTest
    @MethodSource("argumentsTrimZeros")
    fun testTrimZeros(input: String, expected: String) {
        val result = StringUtils.trimZeros(input)
        assertThat(result).isEqualTo(expected)
    }

    private fun argumentsTrimZeros(): Stream<Arguments> = Stream.of(
        Arguments.of("", ""),
        Arguments.of(String(byteArrayOf(0x0)), ""),
        Arguments.of(String(byteArrayOf(0x20, 0x0)), " "),
        Arguments.of(String(byteArrayOf(0x0, 0x20)), String(byteArrayOf(0x20))),
        Arguments.of(String(byteArrayOf(0x0, 0x20, 0x0)), String(byteArrayOf(0x20))),
        Arguments.of(
            String(byteArrayOf(0x0, 0x20, 'a'.code.toByte(), 0x20, 0x0)),
            String(byteArrayOf(0x20, 'a'.code.toByte(), 0x20))
        ),
    )
}
