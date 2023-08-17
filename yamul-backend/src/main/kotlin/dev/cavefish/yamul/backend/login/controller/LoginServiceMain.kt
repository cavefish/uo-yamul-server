package dev.cavefish.yamul.backend.login.controller

import io.grpc.ServerBuilder

class LoginServiceMain {

    companion object {
        @JvmStatic
        fun main(args: Array<String>) {
            val server = ServerBuilder.forPort(8087).addService(LoginServiceController()).build()
            println("Running server on port 8087")
            server.start()
            println("Running ...")
            server.awaitTermination()
        }
    }
}

