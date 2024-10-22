package dev.cavefish.yamul.backend.utils


object StringUtils {
    fun trimZeros(str: String): String {
        var head = 0
        var tail = str.length - 1
        while (head <= tail && str[head] == Char(0)) head++
        while (head <= tail && str[tail] == Char(0)) tail--
        if (head > tail) return ""
        return str.substring(head, tail + 1)
    }
}

fun Any.toStringByRecursion(vararg elements: Pair<String, Any>): String {
    val values = elements.joinToString(", ") {
        "${it.first}=${getPrintableValue(it.second)}"
    }
    return "${this::class.simpleName}(${values})"
}

private fun getPrintableValue(value: Any?): String {
    return when (value) {
        is List<*> -> {
            "L[${value.joinToString(", ") { getPrintableValue(it) }}]"
        }
        is Array<*> -> {
            "A[${value.joinToString(", ") { getPrintableValue(it) }}]"
        }

        else -> value.toString()
    }
}

