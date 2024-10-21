package dev.cavefish.yamul.backend

import com.flextrade.jfixture.JFixture
import java.util.LinkedList

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

fun <T> JFixture.randomize(input: List<T>): List<T> {
    if (input.size <=1) return input
    val result = ArrayList<T>(input.size)
    val used = BooleanArray(input.size)
    for (ignored in input.indices) {
        var nextX = createIntRange(0, input.size - 1)
        while (used[nextX]) nextX = (nextX + 1) % input.size
        result.add(input[nextX])
        used[nextX] = true
    }
    return result
}