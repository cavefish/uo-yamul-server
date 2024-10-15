package dev.cavefish.yamul.backend.game.controller.infra.mul

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.mul.BlockAltitudeData
import org.tinylog.kotlin.Logger

abstract class MulMapBlockRepository {
    abstract fun getBlockAltitudeData(position: Coordinates): BlockAltitudeData

    open fun correctPositionAltitude(cell: Coordinates): Coordinates {
        val blockAltitudeData = getBlockAltitudeData(cell)
        Logger.debug(blockAltitudeData.toString())
        val altitude = blockAltitudeData.getCellAttitude(cell)
        return if (altitude == cell.z) cell else cell.copy(z = altitude)
    }
}