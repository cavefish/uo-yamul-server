package dev.cavefish.yamul.backend.services.global

import dev.cavefish.yamul.backend.game.controller.domain.gameevents.GameEventSendMessage
import dev.cavefish.yamul.backend.game.controller.domain.gameevents.GameStreamEventCoordinator
import org.springframework.scheduling.annotation.Scheduled
import org.springframework.stereotype.Service
import java.time.LocalDateTime

@Service
class PublicClockCronService(
    private val gameStreamEventCoordinator: GameStreamEventCoordinator,
){
    @Scheduled(fixedRate = 1000 * 10)
    fun tick() {
        gameStreamEventCoordinator.notify(GameEventSendMessage("Hello world! It is ${LocalDateTime.now()}"))
    }
}
