package dev.cavefish.yamul.backend.infra.localfile

import org.springframework.stereotype.Repository

private const val UOP_MAPFILE_PAGE_BLOCK_SIZE = 4096L * 196

@Repository
class MultimaFileRepository {

    fun getReaderFor(mulFile: MulFile): MultimaFileReader {
        return when (mulFile.properties) {
            is PlainMulFileProperties -> PlainMultimaFileReader(mulFile.properties.filename)
            is UopFileProperties -> UopMultimaFileReader(
                mulFile.properties.filenames,
                UOP_MAPFILE_PAGE_BLOCK_SIZE,
                mulFile.properties.maxSubFiles,
            ) { mulFile.properties.subFileTemplate.format(it) }
        }
    }

    @SuppressWarnings("MagicNumber")
    enum class MulFile(val properties: MulFileProperties) {
        Map0(
            UopFileProperties(
                0x71,
                listOf("map0LegacyMUL.uop", "map0xLegacyMUL.uop"),
                "build/map0legacymul/%08d.dat"
            )
        ),
        Map1(
            UopFileProperties(
                0x71,
                listOf("map1LegacyMUL.uop", "map1xLegacyMUL.uop"),
                "build/map1legacymul/%08d.dat"
            )
        ),
        Map2(
            UopFileProperties(
                0x71,
                listOf("map2LegacyMUL.uop", "map2xLegacyMUL.uop"),
                "build/map2legacymul/%08d.dat"
            )
        ),
        Map3(UopFileProperties(0x71, listOf("map3LegacyMUL.uop"), "build/map3legacymul/%08d.dat")),
        Map4(UopFileProperties(0x71, listOf("map4LegacyMUL.uop"), "build/map4legacymul/%08d.dat")),
        Map5(
            UopFileProperties(
                0x71,
                listOf("map5LegacyMUL.uop", "map5xLegacyMUL.uop"),
                "build/map5legacymul/%08d.dat"
            )
        ),
    }

    sealed class MulFileProperties

    data class UopFileProperties(
        val maxSubFiles: Int,
        val filenames: List<String>,
        val subFileTemplate: String
    ) : MulFileProperties()

    data class PlainMulFileProperties(
        val filename: String
    ) : MulFileProperties()

}