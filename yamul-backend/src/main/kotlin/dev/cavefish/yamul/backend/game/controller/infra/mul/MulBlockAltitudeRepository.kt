package dev.cavefish.yamul.backend.game.controller.infra.mul

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.mul.BlockAltitudeData

abstract class MulBlockAltitudeRepository {
    abstract fun getBlockAltitudeData(position: Coordinates): BlockAltitudeData

    open fun correctPositionAltitude(cell: Coordinates): Coordinates {
        val blockAltitudeData = getBlockAltitudeData(cell)
        val altitude = blockAltitudeData.getCellAttitude(cell) + 1
        return if (altitude == cell.z) cell else cell.copy(z = altitude)
    }
}