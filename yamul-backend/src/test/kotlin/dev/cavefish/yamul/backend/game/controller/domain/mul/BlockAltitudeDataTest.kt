package dev.cavefish.yamul.backend.game.controller.domain.mul

import com.google.common.collect.Table.Cell
import dev.cavefish.yamul.UnitTest
import dev.cavefish.yamul.backend.createIntRange
import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.randomize

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.TestInstance
import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.Arguments
import org.junit.jupiter.params.provider.MethodSource
import java.util.stream.Stream

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class BlockAltitudeDataTest : UnitTest() {

    @ParameterizedTest
    @MethodSource
    fun getCellAttitude(dx: Int, dy: Int, initZ: Int, altitudeData: BlockAltitudeData, expectedZ: Int) {
        val origin = altitudeData.origin
        val cellAttitude = altitudeData.getCellAttitude(origin.copy(x = origin.x + dx, y = origin.y + dy, z = initZ))
        assertThat(cellAttitude).isEqualTo(expectedZ)
    }

    fun getCellAttitude(): Stream<Arguments> = Stream.of(
        Arguments.of(
            0,
            0,
            10,
            createBlockAltitudeData(0, 0, -1),
            -1
        ),
        Arguments.of(
            1,
            1,
            0,
            createBlockAltitudeData(1, 1, 2),
            2
        ),
        createRandomSimpleCase(),
        createRandomSimpleCase(),
        createRandomSimpleCase(),
        createRandomSimpleCase(),
        Arguments.of(
            0,
            0,
            10,
            createBlockAltitudeData(0, 0, -1, listOf(2, 4)),
            4
        ),
        createComplexCase(-1, 10, 10, 2, 4),
        createComplexCase(1, 0, 0, 2, 4),
        createComplexCase(3, 2, 0, 2, 4),
        createComplexCase(4, 4, 0, 2, 4),
        createComplexCase(5, 4, 0, 2, 4),
    )

    private fun createComplexCase(currentZ: Int, expectedZ: Byte, mapZ: Byte, vararg values: Int): Arguments? {
        val dx = fixture.createIntRange(0, 7)
        val dy = fixture.createIntRange(0, 7)
        return Arguments.of(
            dx,
            dy,
            currentZ,
            createBlockAltitudeData(dx, dy, mapZ, values.toList()),
            expectedZ
        )
    }

    private fun createRandomSimpleCase(): Arguments {
        val dx = fixture.createIntRange(0, 7)
        val dy = fixture.createIntRange(0, 7)
        val expectedZ = fixture.create(Byte::class.java)
        return Arguments.of(
            dx,
            dy,
            fixture.create(Byte::class.java),
            createBlockAltitudeData(dx, dy, expectedZ),
            expectedZ
        )
    }


    private fun createBlockAltitudeData(
        dx: Int,
        dy: Int,
        zValueOnMap: Byte
    ): BlockAltitudeData {
        val mapValues: Array<Array<Pair<Short, Byte>>> =
            Array(8) { Array(8) { fixture.create(Short::class.java) to fixture.create(Byte::class.java) } }
        mapValues[dx][dy] = fixture.create(Short::class.java) to zValueOnMap
        val altitudeData = BlockAltitudeData(
            origin = Coordinates(
                x = 8 * fixture.createIntRange(0, 100),
                y = 8 * fixture.createIntRange(0, 100),
                z = 0,
                mapId = fixture.createIntRange(0, 5)
            ),
            mapValues = mapValues,
            staticCells = emptyList()
        )
        return altitudeData
    }


    private fun createBlockAltitudeData(
        dx: Int,
        dy: Int,
        zValueOnMap: Byte,
        cellAltitudes: List<Int>
    ): BlockAltitudeData {
        val mapValues: Array<Array<Pair<Short, Byte>>> =
            Array(8) { Array(8) { fixture.create(Short::class.java) to fixture.create(Byte::class.java) } }
        mapValues[dx][dy] = fixture.create(Short::class.java) to zValueOnMap
        val staticCells = ArrayList<StaticCellData>()
        staticCells.addAll(cellAltitudes.map {
            StaticCellData(
                objectId = fixture.create(Int::class.java),
                x = dx,
                y = dy,
                z = it
            )
        })
        val altitudeData = BlockAltitudeData(
            origin = Coordinates(
                x = 8 * fixture.createIntRange(0, 100),
                y = 8 * fixture.createIntRange(0, 100),
                z = 0,
                mapId = fixture.createIntRange(0, 5)
            ),
            mapValues = mapValues,
            staticCells = fixture.randomize(staticCells)
        )
        return altitudeData
    }
}
