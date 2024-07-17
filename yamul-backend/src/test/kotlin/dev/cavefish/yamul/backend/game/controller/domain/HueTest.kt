package dev.cavefish.yamul.backend.game.controller.domain

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.Arguments
import org.junit.jupiter.params.provider.MethodSource
import java.util.stream.Stream


class HueTest {

    @ParameterizedTest
    @MethodSource
    fun toInt16(expected: Int, value: Hue) {
        assertThat(value.toInt16()).isEqualTo(expected)
    }

    companion object {
        @JvmStatic
        private fun toInt16(): Stream<Arguments> = Stream.of(
            Arguments.of(0b1_00000_00000_00000, Hue(red = 0, green = 0, blue = 0)),
            Arguments.of(0b1_11111_00000_00000, Hue(red = 255, green = 0, blue = 0)),
            Arguments.of(0b1_00000_11111_00000, Hue(red = 0, green = 255, blue = 0)),
            Arguments.of(0b1_00000_00000_11111, Hue(red = 0, green = 0, blue = 255)),
            Arguments.of(0b1_11111_11111_11111, Hue(red = 255, green = 255, blue = 255)),
        )
    }
}