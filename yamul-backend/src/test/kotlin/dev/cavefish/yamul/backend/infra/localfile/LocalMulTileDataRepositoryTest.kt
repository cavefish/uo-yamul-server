package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.IntegrationTest
import dev.cavefish.yamul.backend.game.controller.domain.mul.LandTileData
import dev.cavefish.yamul.backend.game.controller.domain.mul.StaticTileData
import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.TestInstance
import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.Arguments
import org.junit.jupiter.params.provider.MethodSource
import org.springframework.beans.factory.annotation.Autowired
import java.util.stream.Stream

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
@SuppressWarnings("LongMethod")
class LocalMulTileDataRepositoryTest : IntegrationTest() {

    @Autowired
    lateinit var repository: LocalMulTileDataRepository

    @ParameterizedTest
    @MethodSource
    fun readLandTileData(id: Int, expected: LandTileData) {
        val result = repository.getLandTileData(id)

        assertThat(result).isEqualTo(
            expected
        )
    }

    fun readLandTileData(): Stream<Arguments> = Stream.of(
        Arguments.of(
            1, LandTileData(
                name = "VOID!!!!!!",
                textureId = 1,
                flags = 0x50
            )
        ),
        Arguments.of(
            2, LandTileData(
                name = "NODRAW",
                textureId = 2,
                flags = 0x0
            )
        ),
        Arguments.of(
            3, LandTileData(
                name = "grass",
                textureId = 3,
                flags = 0x0
            )
        ),
        Arguments.of(
            32,
            LandTileData(
                name = "sand",
                textureId = 32,
                flags = 0x40
            ),
        ),
        Arguments.of(
            16337, // Last "valid" land tile
            LandTileData(
                name = "NoName",
                textureId = 16337,
                flags = 0x0
            ),
        ),
        Arguments.of(
            16338, // First empty tile data
            LandTileData(
                name = "",
                textureId = 0,
                flags = 0x0
            ),
        ),
    )

    @ParameterizedTest
    @MethodSource
    fun getStaticTileData(id: Int, expected: StaticTileData) {
        val result = repository.getStaticTileData(id)
        assertThat(result).isEqualTo(expected)
    }

    fun getStaticTileData() : Stream<Arguments> = Stream.of(
        Arguments.of(
            1,
            StaticTileData(
                id = 1,
                flags = 0x02_00_00,
                weight = 0u,
                layer = 0u,
                count = 0,
                animId = 0,
                hue = 0,
                lightIndex = 0,
                height = 0u,
                name = "nodraw",
            ),
        ),
        Arguments.of(
            2,
            StaticTileData(
                id = 2,
                flags = 0x04_00_A0_40,
                weight = 0xFFu,
                layer = 0u,
                count = 0,
                animId = 0,
                hue = 0,
                lightIndex = 0,
                height = 12u,
                name = "ankh",
            ),
        ),
        Arguments.of(
            31,
            StaticTileData(
                id = 31,
                flags = 0x60_50,
                weight = 0xFFu,
                layer = 0u,
                count = 0,
                animId = 0,
                hue = 0,
                lightIndex = 0,
                height = 20u,
                name = "stone wall",
            ),
        ),
        Arguments.of(
            32,
            StaticTileData(
                id = 32,
                flags = 0x04_00_60_50,
                weight = 0xFFu,
                layer = 0u,
                count = 0,
                animId = 0,
                hue = 0,
                lightIndex = 0,
                height = 20u,
                name = "stone wall",
            ),
        ),
        Arguments.of(
            45090,
            StaticTileData(
                id = 45090,
                flags = 0x40,
                weight = 0xFFu,
                layer = 0u,
                count = 0,
                animId = 0,
                hue = 0,
                lightIndex = 0,
                height = 10u,
                name = "Zombie Banner East A",
            ),
        ),
        Arguments.of(
            45091,
            StaticTileData(
                id = 45091,
                flags = 0,
                weight = 0u,
                layer = 0u,
                count = 0,
                animId = 0,
                hue = 0,
                lightIndex = 0,
                height = 0u,
                name = "",
            ),
        ),
        Arguments.of(
            0x204F,
            StaticTileData(
                id = 0x204F,
                flags = 0x0C_44_40_00,
                weight = 0u,
                layer = 22u,
                count = 0,
                animId = 987,
                hue = 0,
                lightIndex = 0,
                height = 1u,
                name = "GM Robe",
            ),
        ),
    )
}