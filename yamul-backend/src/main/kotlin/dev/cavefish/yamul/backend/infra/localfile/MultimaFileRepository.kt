package dev.cavefish.yamul.backend.infra.localfile

import org.springframework.stereotype.Repository

@Repository
class MultimaFileRepository {

    fun getReaderFor(filename: String): MultimaFileReader {
        return if (filename.endsWith(".uop")) UopMultimaFileReader(filename)
        else PlainMultimaFileReader(filename)
    }

}