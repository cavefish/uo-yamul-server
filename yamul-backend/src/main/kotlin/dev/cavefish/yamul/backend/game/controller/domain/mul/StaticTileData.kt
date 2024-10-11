package dev.cavefish.yamul.backend.game.controller.domain.mul

data class StaticTileData(
    val name: String,
    val id: Int,
    val flags: Long,
    val weight: UByte,
    val layer: UByte,
    val count: Int,
    val animId: Short,
    val hue: Short,
    val lightIndex: Short,
    val height: UByte
)
