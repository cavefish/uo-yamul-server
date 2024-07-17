package dev.cavefish.yamul.backend.common.controller

import dev.cavefish.yamul.backend.auth.controller.BasicAuthInterceptor
import dev.cavefish.yamul.backend.character.controller.CharacterServiceController
import dev.cavefish.yamul.backend.game.controller.GameServiceController
import dev.cavefish.yamul.backend.infra.inmemory.InMemoryInitRepositories
import dev.cavefish.yamul.backend.login.controller.LoginServiceController
import io.grpc.ServerBuilder
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.Banner
import org.springframework.boot.WebApplicationType
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.context.annotation.ComponentScan
import org.tinylog.kotlin.Logger
import javax.annotation.PostConstruct

const val LOGIN_SERVICE_PORT = 8087
const val CHARACTER_SERVICE_PORT = 8088
const val GAME_SERVICE_PORT = 8089

@SpringBootApplication
@ComponentScan(basePackages = ["dev.cavefish.yamul.backend"])
class ServiceMain @Autowired constructor(
    val loginServiceController: LoginServiceController,
    val characterServiceController: CharacterServiceController,
    val gameServiceController: GameServiceController,
    val basicAuthInterceptor: BasicAuthInterceptor,
    val initMemoryRepositories: InMemoryInitRepositories,
) {

    @PostConstruct
    fun runServices() {
        val loginServer = ServerBuilder.forPort(LOGIN_SERVICE_PORT).addService(loginServiceController).build()
        loginServer.start()
        Logger.info("Running Login server on port {0}", LOGIN_SERVICE_PORT)

        val characterServer = ServerBuilder.forPort(CHARACTER_SERVICE_PORT).intercept(basicAuthInterceptor)
            .addService(characterServiceController).build()
        characterServer.start()
        Logger.info("Running Character server on port {0}", CHARACTER_SERVICE_PORT)

        val gameServer =
            ServerBuilder.forPort(GAME_SERVICE_PORT).intercept(basicAuthInterceptor).addService(gameServiceController)
                .build()
        gameServer.start()
        Logger.info("Running Game server on port {0}", GAME_SERVICE_PORT)

        initMemoryRepositories.init()

        Logger.info("Running ...")
        loginServer.awaitTermination()
        characterServer.awaitTermination()
        gameServer.awaitTermination()
    }

}

fun main(args: Array<String>) {
    runApplication<ServiceMain>(*args) {
        setBannerMode(Banner.Mode.OFF)
        webApplicationType = WebApplicationType.NONE
    }
}
