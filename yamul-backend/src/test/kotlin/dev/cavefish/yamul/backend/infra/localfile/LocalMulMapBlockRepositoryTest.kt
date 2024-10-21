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
import org.mockito.kotlin.never
import org.mockito.kotlin.verify
import org.mockito.kotlin.whenever
import java.nio.ByteBuffer
import java.util.stream.Stream

private const val BLOCK_SIZE = 196L

@TestInstance(TestInstance.Lifecycle.PER_CLASS)
class LocalMulMapBlockRepositoryTest : IntegrationTest() {

    @Mock
    lateinit var mapFileReader: MultimaFileReader

    @Mock
    lateinit var staticsFileReader: MultimaFileReader

    @Mock
    lateinit var multimaFileRepository: MultimaFileRepository

    @InjectMocks
    lateinit var repository: LocalMulMapBlockRepository

    @BeforeEach
    override fun beforeEach() {
        super.beforeEach()
        Mockito.clearInvocations(mapFileReader, staticsFileReader, multimaFileRepository)
    }

    @ParameterizedTest
    @MethodSource
    fun correctPositionAltitude(
        coordinates: Coordinates,
        blockFilePosition: Long,
        mapBytes: ByteArray?,
        staticsBytes: ByteArray?
    ) {
        val mapByteBuffer = if (mapBytes != null) ByteBuffer.wrap(mapBytes) else null
        val staticsBuffer = if (staticsBytes != null) ByteBuffer.wrap(staticsBytes) else null

        // Given
        whenever(multimaFileRepository.getReaderFor(MulMapHelper.mapProperties[coordinates.mapId].mapFile))
            .thenReturn(mapFileReader)
        whenever(multimaFileRepository.getReaderFor(MulMapHelper.mapProperties[coordinates.mapId].staticsFile))
            .thenReturn(staticsFileReader)
        whenever(mapFileReader.getBuffer(blockFilePosition * BLOCK_SIZE, BLOCK_SIZE))
            .thenReturn(mapByteBuffer)
        whenever(staticsFileReader.getBuffer(blockFilePosition))
            .thenReturn(staticsBuffer)

        // When
        val result = repository.correctPositionAltitude(coordinates.copy(z = 300))

        // Then
        softly.assertThat(result).isEqualTo(coordinates)
        verify(multimaFileRepository).getReaderFor(MulMapHelper.mapProperties[coordinates.mapId].mapFile)
        verify(multimaFileRepository).getReaderFor(MulMapHelper.mapProperties[coordinates.mapId].staticsFile)
        verify(mapFileReader).getBuffer(blockFilePosition * BLOCK_SIZE, BLOCK_SIZE)
        if (mapByteBuffer != null) softly.assertThat(mapByteBuffer.remaining())
            .describedAs("Only the footer remains").isLessThanOrEqualTo(4)
        if (staticsBuffer != null) softly.assertThat(staticsBuffer.remaining())
            .describedAs("The buffer is fully consumed").isEqualTo(0)
        verify(staticsFileReader).getBuffer(blockFilePosition, null)
        verify(mapFileReader, never()).close()
        verify(staticsFileReader, never()).close()

        // and When Close
        repository.close()
        // Then
        verify(mapFileReader).close()
        verify(staticsFileReader).close()
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
            createBlockArray(2, 0, 0),
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
            null,
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
            createBlockArray(2, 0),
            null
        ),
        Arguments.of(
            Coordinates(
                x = 0,
                y = 0,
                z = 1,
                mapId = 0
            ),
            0L,
            createBlockArray(2, 0),
            byteArrayOf(
                0xFF.toByte(), 0xFF.toByte(), 0, 0, 1, 0xFF.toByte(), 0xFF.toByte(),
            )
        ),
        Arguments.of(
            Coordinates(
                x = 0,
                y = 0,
                z = 0,
                mapId = 0
            ),
            0L,
            createBlockArray(2, 0),
            byteArrayOf(
                0xFF.toByte(), 0xFF.toByte(), 1, 1, 10, 0xFF.toByte(), 0xFF.toByte(),
            )
        ),
        Arguments.of(
            Coordinates(
                x = 0,
                y = 0,
                z = 12,
                mapId = 0
            ),
            0L,
            createBlockArray(2, 0),
            byteArrayOf(
                0xFF.toByte(), 0xFF.toByte(), 0, 0, 1, 0xFF.toByte(), 0xFF.toByte(),
                0xFF.toByte(), 0xFF.toByte(), 0, 0, 12, 0xFF.toByte(), 0xFF.toByte(),
            )
        ),
        Arguments.of(
            Coordinates(
                x = 0,
                y = 0,
                z = 0,
                mapId = 0
            ),
            0L,
            createBlockArray(2, 0),
            null
        ),
        Arguments.of(
            Coordinates(
                x = 0,
                y = 0,
                z = -1,
                mapId = 0
            ),
            0L,
            createBlockArray(2, -1, 100),
            null
        ),
        Arguments.of(
            Coordinates(
                x = 1,
                y = 0,
                z = 123,
                mapId = 0
            ),
            0L,
            createBlockArray(2 + 3, 123),
            null
        ),
        Arguments.of(
            Coordinates(
                x = 0,
                y = 1,
                z = 123,
                mapId = 0
            ),
            0L,
            createBlockArray(2 + 8 * 3, 123),
            null
        ),
        Arguments.of(
            Coordinates(
                x = 0,
                y = 8,
                z = -1,
                mapId = 0
            ),
            1L,
            createBlockArray(2, -1),
            null
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
                x = subX + blockX * 8,
                y = subY + blockY * 8,
                z = expectedValueZ.toInt(),
                mapId = mapIdToUse
            ),
            blockX * mapBlockHeights[mapIdToUse]!! + blockY,
            createBlockArray(2 + 3 * (subX + subY * 8), expectedValueZ),
            null // TODO create random statics
        )
    }

    private fun createBlockArray(position: Int, value: Byte, default: Byte? = null): ByteArray =
        ByteArray(BLOCK_SIZE.toInt()) {
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



