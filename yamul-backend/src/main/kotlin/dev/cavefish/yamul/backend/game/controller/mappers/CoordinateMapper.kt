package dev.cavefish.yamul.backend.game.controller.mappers

import dev.cavefish.yamul.backend.common.api.Coordinate
import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import org.springframework.stereotype.Service

@Service
class CoordinateMapper {
    fun map(obj: Coordinates) = Coordinate.newBuilder().setXLoc(obj.x).setYLoc(obj.y).setZLoc(obj.z)
}
