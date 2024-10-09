package dev.cavefish.yamul.backend.game.controller.infra.mul

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.mul.BlockAltitudeData

interface MulBlockAltitudeRepository {
    fun getBlockAltitudeData(position: Coordinates): BlockAltitudeData
}