package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates

@SuppressWarnings("MagicNumber")
object MulBlockHelper {

    fun getBlockId(origin: Coordinates): Int {
        val blockX = origin.x ushr 3
        val blockY = origin.y ushr 3
        return blockX * mapWidths[origin.mapId]!! + blockY
    }

    private val mapWidths = mapOf(
        0 to 4096 / 8,
        1 to 4096 / 8,
        2 to 1600 / 8,
        3 to 2048 / 8,
        4 to 1448 / 8,
        5 to 4096 / 8,
    )
}