package dev.cavefish.yamul.backend.common.controller

import dev.cavefish.yamul.backend.auth.controller.BasicAuthInterceptor
import dev.cavefish.yamul.backend.character.controller.CharacterServiceController
import dev.cavefish.yamul.backend.login.controller.LoginServiceController
import io.grpc.ServerBuilder
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.Banner
import org.springframework.boot.WebApplicationType
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.context.annotation.ComponentScan
import javax.annotation.PostConstruct

const val LOGIN_SERVICE_PORT = 8087

const val CHARACTER_SERVICE_PORT = 8088

@SpringBootApplication
@ComponentScan(basePackages = ["dev.cavefish.yamul.backend"])
class ServiceMain @Autowired constructor(
    val loginServiceController: LoginServiceController,
    val characterServiceController: CharacterServiceController,
    val basicAuthInterceptor: BasicAuthInterceptor
    ) {

    @PostConstruct
    fun runServices() {
        val loginServer = ServerBuilder.forPort(LOGIN_SERVICE_PORT).addService(loginServiceController).build()
        loginServer.start()
        println("Running Login server on port $LOGIN_SERVICE_PORT")

        val characterServer = ServerBuilder.forPort(CHARACTER_SERVICE_PORT).intercept(basicAuthInterceptor)
            .addService(characterServiceController).build()
        characterServer.start()
        println("Running Character server on port $CHARACTER_SERVICE_PORT")

        println("Running ...")
        loginServer.awaitTermination()
        characterServer.awaitTermination()
    }

}

fun main(args: Array<String>) {
    runApplication<ServiceMain>(*args) {
        setBannerMode(Banner.Mode.OFF)
        webApplicationType = WebApplicationType.NONE
    }
}
