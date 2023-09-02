package dev.cavefish.yamul.backend.utils


object StringUtils {
    fun trimZeros(str: String): String {
        var head = 0
        var tail = str.length - 1
        while (head<=tail && str[head] == Char(0)) head++
        while (head<=tail && str[tail] == Char(0)) tail--
        if (head > tail) return ""
        return str.substring(head, tail+1)
    }
}
