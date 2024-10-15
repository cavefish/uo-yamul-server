package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.IntegrationTest
import dev.cavefish.yamul.backend.createDifferent
import dev.cavefish.yamul.backend.createIntRange
import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.TestInstance
import org.junit.jupiter.params.ParameterizedTest
import org.junit.jupiter.params.provider.Arguments
import org.junit.jupiter.params.provider.MethodSource
import org.mockito.InjectMocks
import org.mockito.Mock
import org.mockito.Mockito
import org.mockito.kotlin.any
import org.mockito.kotlin.never
import org.mockito.kotlin.verify
import org.mockito.kotlin.whenever
import java.util.stream.Stream

private const val BLOCK_SIZE = 196

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class LocalMulMapBlockRepositoryTest : IntegrationTest() {

    @Mock
    lateinit var fileReader: MultimaFileReader

    @Mock
    lateinit var multimaFileRepository: MultimaFileRepository

    @InjectMocks
    lateinit var repository: LocalMulMapBlockRepository

    @BeforeEach
    override fun beforeEach() {
        super.beforeEach()
        Mockito.clearInvocations(fileReader, multimaFileRepository)
    }

    @ParameterizedTest
    @MethodSource
    fun correctPositionAltitude(coordinates: Coordinates, blockFilePosition: Long, bytes: ByteArray?) {
        // Given
        whenever(multimaFileRepository.getReaderFor(any())).thenReturn(fileReader)
        whenever(fileReader.getBytes(blockFilePosition * BLOCK_SIZE, BLOCK_SIZE)).thenReturn(bytes)

        // When
        val result = repository.correctPositionAltitude(coordinates.copy(z = 300))

        // Then
        verify(fileReader).getBytes(blockFilePosition * BLOCK_SIZE, BLOCK_SIZE)
        softly.assertThat(result).isEqualTo(coordinates)

        verify(fileReader, never()).close()
        repository.close()
        verify(fileReader).close()
    }

    @SuppressWarnings("LongMethod")
    fun correctPositionAltitude(): Stream<Arguments> = Stream.of(
        Arguments.of(
            Coordinates(
                x = 0,
                y = 0,
                z = 0,
                mapId = 0
            ),
            0L,
            createBlockArray(6, 0, 0)
        ),
        Arguments.of(
            Coordinates(
                x = 0,
                y = 0,
                z = 0,
                mapId = 0
            ),
            0L,
            null
        ),
        Arguments.of(
            Coordinates(
                x = 0,
                y = 0,
                z = 0,
                mapId = 0
            ),
            0L,
            createBlockArray(6, 0)
        ),
        Arguments.of(
            Coordinates(
                x = 0,
                y = 0,
                z = -1,
                mapId = 0
            ),
            0L,
            createBlockArray(6, -1, 100)
        ),
        Arguments.of(
            Coordinates(
                x = 1,
                y = 0,
                z = 123,
                mapId = 0
            ),
            0L,
            createBlockArray(6 + 3, 123)
        ),
        Arguments.of(
            Coordinates(
                x = 0,
                y = 1,
                z = 123,
                mapId = 0
            ),
            0L,
            createBlockArray(6 + 8 * 3, 123)
        ),
        Arguments.of(
            Coordinates(
                x = 0,
                y = 8,
                z = -1,
                mapId = 0
            ),
            1L,
            createBlockArray(6, -1)
        ),
        createRandomArguments(),
        createRandomArguments(),
        createRandomArguments(),
        createRandomArguments(),
        createRandomArguments(),
    )

    private fun createRandomArguments(mapId: Int? = null): Arguments {
        val mapIdToUse = mapId ?: fixture.createIntRange(0, 5)
        val expectedValueZ = fixture.create(Byte::class.java)
        val subX = fixture.createIntRange(0, 7)
        val subY = fixture.createIntRange(0, 7)
        val blockX = fixture.createIntRange(0, 100)
        val blockY = fixture.createIntRange(0, 100)
        return Arguments.of(
            Coordinates(
                x = subX + blockX*8,
                y = subY + blockY*8,
                z = expectedValueZ.toInt(),
                mapId = mapIdToUse
            ),
            blockX*mapBlockHeights[mapIdToUse]!! + blockY,
            createBlockArray(6 + 3 * (subX + subY * 8), expectedValueZ)
        )
    }

    private fun createBlockArray(position: Int, value: Byte, default: Byte? = null): ByteArray = ByteArray(BLOCK_SIZE) {
        if (it == position) value else default ?: fixture.createDifferent(Byte::class.java, value)
    }

    companion object {
        val mapBlockHeights = mapOf(
            0 to 512,
            1 to 512,
            2 to 200,
            3 to 256,
            4 to 181,
            5 to 512,
        )
    }
}



