@file:OptIn(ExperimentalStdlibApi::class)

package dev.cavefish.yamul.backend

import dev.cavefish.yamul.backend.game.controller.domain.LoggedUser
import io.grpc.Context

object Constants {
    val AUTH_CONTEXT_LOGGED_USER: Context.Key<LoggedUser> = Context.key("AUTH_CONTEXT_LOGGED_USER")
    const val MULTIMA_PATH = "multima.mulfiles.dir"
    @OptIn(ExperimentalStdlibApi::class)
    val hexNumberFormat = HexFormat {
        upperCase = true
        number {
            removeLeadingZeros = false
            prefix = "0x"
        }
    }

    fun toHexFormat(input: String) = input.toByte().toHexString(hexNumberFormat)
    fun toHexFormat(input: ULong) = input.toHexString(hexNumberFormat)
    fun toHexFormat(input: Long) = input.toHexString(hexNumberFormat)
    fun toHexFormat(input: Int) = input.toHexString(hexNumberFormat)
    fun toHexFormat(input: Byte) = input.toHexString(hexNumberFormat)
    fun toHexFormat(input: Short) = input.toHexString(hexNumberFormat)
}
