package dev.cavefish.yamul.backend.infra.localfile

import org.springframework.stereotype.Repository

private const val UOP_MAPFILE_PAGE_BLOCK_SIZE = 4096L * 196

@Repository
class MultimaFileRepository {

    fun getReaderFor(properties: MulFileProperties): MultimaFileReader {
        return when (properties) {
            is PlainMulFileProperties -> PlainMultimaFileReader(properties.filename)
            is UopFileProperties -> UopMultimaFileReader(
                properties.filenames,
                UOP_MAPFILE_PAGE_BLOCK_SIZE,
                properties.maxSubFiles,
            ) { properties.subFileTemplate.format(it) }
            is IndexedFileProperties -> TODO()
        }
    }

    sealed class MulFileProperties

    data class UopFileProperties(
        val maxSubFiles: Int,
        // Uop files can auto-patch themselves.
        // You need to override one with the contents of the next
        val filenames: List<String>,
        val subFileTemplate: String
    ) : MulFileProperties()

    data class PlainMulFileProperties(
        val filename: String
    ) : MulFileProperties()

    data class IndexedFileProperties(
        val baseFilename: String,
        val idxFilename: String
    ) : MulFileProperties()

}