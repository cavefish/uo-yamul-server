package dev.cavefish.yamul.backend.game.controller

import dev.cavefish.yamul.backend.game.api.MsgType
import dev.cavefish.yamul.backend.game.controller.processors.MessageProcessor
import org.springframework.context.ApplicationContext
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration


@Configuration
class GameControllerBeanProvider {

    @Bean
    fun processorMap(context: ApplicationContext): Map<MsgType, MessageProcessor<*>> {
        return context.getBeansOfType(MessageProcessor::class.java).mapKeys { entry -> entry.value.getType() }
    }
}