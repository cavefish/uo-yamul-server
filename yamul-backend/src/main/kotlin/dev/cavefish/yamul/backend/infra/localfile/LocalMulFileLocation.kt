package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.backend.Constants
import java.io.File

object LocalMulFileLocation {
    private fun getBasePath(): String {
        val basePath = System.getenv()[Constants.MULTIMA_PATH]
        assert(!basePath.isNullOrBlank()) { "${Constants.MULTIMA_PATH} must be initialized" }
        return basePath!!
    }

    fun getFileLocation(name: String): String? {
        assert(name.endsWith(".mul") || name.endsWith(".uop")) {"Invalid filename: $name"}
        assert(name.contains(Regex("\\w+"))) {"Filename must contain valid characters"}
        val basePath = getBasePath()
        val baseDirectory = File(basePath)
        if (!baseDirectory.exists() && !baseDirectory.isDirectory) return null
        val file = baseDirectory.resolve(name)
        if (!file.exists() && !file.isFile) return null
        return file.toString()
    }
}