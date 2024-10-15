package dev.cavefish.yamul.backend

import com.flextrade.jfixture.JFixture

fun <T> JFixture.createDifferent(clazz: Class<T>, value: T): T {
    do {
        val x = create(clazz)!!
        if (x != value) return x
    } while (true)
}

fun JFixture.createIntRange(from: Int, to: Int): Int {
    val value = create(Int::class.java)
    if (value in from..to) return value

    val diff = to - from + 1 // Integer range includes limits
    val x = value % diff
    return from + x
}