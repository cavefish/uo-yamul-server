package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.mul.BlockAltitudeData
import dev.cavefish.yamul.backend.game.controller.infra.mul.MulBlockAltitudeRepository
import org.springframework.stereotype.Repository

private const val BLOCK_SIZE = 8

@Repository
class LocalMulBlockAltitudeRepository: MulBlockAltitudeRepository {
    override fun getBlockAltitudeData(position: Coordinates): BlockAltitudeData {
        return BlockAltitudeData(position.toBlockOrigin(), Array(BLOCK_SIZE) {Array(BLOCK_SIZE){0} })
    }
}