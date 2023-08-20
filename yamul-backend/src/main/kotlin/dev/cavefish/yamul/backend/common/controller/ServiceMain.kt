package dev.cavefish.yamul.backend.common.controller

import dev.cavefish.yamul.backend.auth.controller.BasicAuthInterceptor
import dev.cavefish.yamul.backend.character.controller.CharacterServiceController
import dev.cavefish.yamul.backend.login.controller.LoginServiceController
import io.grpc.ServerBuilder

private const val LOGIN_SERVICE_PORT = 8087

private const val CHARACTER_SERVICE_PORT = 8088

object ServiceMain {

    @JvmStatic
    fun main(args: Array<String>) {
        val loginServer = ServerBuilder.forPort(LOGIN_SERVICE_PORT).addService(LoginServiceController()).build()
        loginServer.start()
        println("Running Login server on port $LOGIN_SERVICE_PORT")

        val characterServer = ServerBuilder.forPort(CHARACTER_SERVICE_PORT)
            .intercept(BasicAuthInterceptor())
            .addService(CharacterServiceController()).build()
        characterServer.start()
        println("Running Character server on port $CHARACTER_SERVICE_PORT")

        println("Running ...")
        loginServer.awaitTermination()
        characterServer.awaitTermination()
    }

}