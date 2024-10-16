package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.UnitTest
import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import org.junit.jupiter.api.TestInstance

import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.Arguments
import org.junit.jupiter.params.provider.MethodSource
import java.util.stream.Stream

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class MulBlockHelperTest : UnitTest() {

    @ParameterizedTest
    @MethodSource
    fun getBlock(expectedBlock: Long, coordinates: Coordinates) {
        val blockId = MulBlockHelper.getBlockId(coordinates)
        softly.assertThat(blockId).isEqualTo(expectedBlock)
    }

    fun getBlock(): Stream<Arguments> = Stream.of(
        Arguments.of(0L, Coordinates(x = 0, y = 0, mapId = 0)),
        Arguments.of(0L, Coordinates(x = 7, y = 7, mapId = 0)),
        Arguments.of(0L, Coordinates(x = 7, y = 7, mapId = 1)),
        Arguments.of(1L, Coordinates(x = 0, y = 8, mapId = 0)),
        Arguments.of(512L, Coordinates(x = 8, y = 0, mapId = 0)),
        Arguments.of(513L, Coordinates(x = 8, y = 8, mapId = 0)),
    )

}