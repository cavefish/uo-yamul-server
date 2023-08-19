package dev.cavefish.yamul.backend.utils


object StringUtils {
    fun trimZeros(str: String): String {
        val pos = str.indexOf(Char(0))
        return if (pos < 0) {
            str
        } else {
            str.substring(0, pos)
        }
    }
}
