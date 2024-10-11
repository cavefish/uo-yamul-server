package dev.cavefish.yamul.backend.game.controller.infra.mul

import dev.cavefish.yamul.backend.game.controller.domain.mul.LandTileData
import dev.cavefish.yamul.backend.game.controller.domain.mul.StaticTileData

interface MulTileDataRepository {
    fun getLandTileData(id: Int): LandTileData?
    fun getStaticTileData(id: Int): StaticTileData?
}