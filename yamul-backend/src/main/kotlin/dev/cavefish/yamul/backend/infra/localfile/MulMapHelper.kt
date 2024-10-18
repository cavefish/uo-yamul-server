package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.infra.localfile.MultimaFileRepository.UopFileProperties

@SuppressWarnings("MagicNumber")
object MulMapHelper {

    fun getBlockId(origin: Coordinates): Int {
        val blockX = origin.x ushr 3
        val blockY = origin.y ushr 3
        return blockX * mapProperties[origin.mapId].width / 8 + blockY
    }

    val mapProperties = arrayOf(
        MapProperties(
            id = 0,
            height = 7168,
            width = 4096,
            mapFile = UopFileProperties(
                maxSubFiles = 0x71,
                filenames = listOf("map0LegacyMUL.uop", "map0xLegacyMUL.uop"),
                subFileTemplate = "build/map0legacymul/%08d.dat"
            )
        ),
        MapProperties(
            id = 1,
            height = 7168,
            width = 4096,
            mapFile = UopFileProperties(
                maxSubFiles = 0x71,
                filenames = listOf("map1LegacyMUL.uop", "map1xLegacyMUL.uop"),
                subFileTemplate = "build/map1legacymul/%08d.dat"
            )
        ),
        MapProperties(
            id = 2,
            height = 2304,
            width = 1600,
            mapFile = UopFileProperties(
                maxSubFiles = 0x71,
                filenames = listOf("map2LegacyMUL.uop", "map2xLegacyMUL.uop"),
                subFileTemplate = "build/map2legacymul/%08d.dat"
            )
        ),
        MapProperties(
            id = 3,
            height = 2560,
            width = 2048,
            mapFile = UopFileProperties(
                maxSubFiles = 0x71,
                filenames = listOf("map3LegacyMUL.uop"),
                subFileTemplate = "build/map3legacymul/%08d.dat"
            )
        ),
        MapProperties(
            id = 4,
            height = 1448,
            width = 1448,
            mapFile = UopFileProperties(
                maxSubFiles = 0x71,
                filenames = listOf("map4LegacyMUL.uop"),
                subFileTemplate = "build/map4legacymul/%08d.dat"
            )
        ),
        MapProperties(
            id = 5,
            height = 1280,
            width = 4096,
            mapFile = UopFileProperties(
                maxSubFiles = 0x71,
                filenames = listOf("map5LegacyMUL.uop", "map5xLegacyMUL.uop"),
                subFileTemplate = "build/map5legacymul/%08d.dat"
            )
        ),
    )

    data class MapProperties(
        val id: Int,
        val height: Int,
        val width: Int,
        val mapFile: MultimaFileRepository.MulFileProperties,
    )
}