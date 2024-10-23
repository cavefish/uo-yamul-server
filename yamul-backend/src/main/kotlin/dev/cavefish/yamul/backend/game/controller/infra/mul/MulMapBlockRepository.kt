package dev.cavefish.yamul.backend.game.controller.infra.mul

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.mul.BlockAltitudeData

abstract class MulMapBlockRepository {
    abstract val mulTileDataRepository: MulTileDataRepository
    abstract fun getBlockAltitudeData(position: Coordinates): BlockAltitudeData

    open fun correctPositionAltitude(cell: Coordinates, bodyHeight: Int): Coordinates {
        val blockAltitudeData = getBlockAltitudeData(cell)
        val altitude = blockAltitudeData.getCellAttitude(cell, bodyHeight, mulTileDataRepository)
        return if (altitude == cell.z) cell else cell.copy(z = altitude)
    }
}