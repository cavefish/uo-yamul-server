package dev.cavefish.yamul.backend.login.controller

import io.grpc.ServerBuilder

private const val SERVICE_PORT = 8087

object LoginServiceMain {

    @JvmStatic
    fun main(args: Array<String>) {
        val server = ServerBuilder.forPort(SERVICE_PORT).addService(LoginServiceController()).build()
        println("Running server on port $SERVICE_PORT")
        server.start()
        println("Running ...")
        server.awaitTermination()
    }

}

