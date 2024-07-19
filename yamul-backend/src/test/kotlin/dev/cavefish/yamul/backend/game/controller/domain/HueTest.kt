package dev.cavefish.yamul.backend.game.controller.domain

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.Arguments
import org.junit.jupiter.params.provider.MethodSource
import java.util.stream.Stream


class HueTest {

    @ParameterizedTest
    @MethodSource
    fun toUInt16(expected: Int, value: Hue) {
        assertThat(value.toUInt16()).isEqualTo(expected)
    }

    companion object {
        @JvmStatic
        private fun toUInt16(): Stream<Arguments> = Stream.of(
            Arguments.of(0b1_00000_00000_00000, Hue(red = 0u, green = 0u, blue = 0u)),
            Arguments.of(0b1_11111_00000_00000, Hue(red = 255u, green = 0u, blue = 0u)),
            Arguments.of(0b1_00000_11111_00000, Hue(red = 0u, green = 255u, blue = 0u)),
            Arguments.of(0b1_00000_00000_11111, Hue(red = 0u, green = 0u, blue = 255u)),
            Arguments.of(0b1_11111_11111_11111, Hue(red = 255u, green = 255u, blue = 255u)),
        )
    }
}