package dev.cavefish.yamul.backend.infra.inmemory

import dev.cavefish.yamul.backend.game.controller.domain.Character
import dev.cavefish.yamul.backend.game.controller.domain.Flags
import dev.cavefish.yamul.backend.game.controller.domain.GameObject
import dev.cavefish.yamul.backend.game.controller.domain.GraphicId
import dev.cavefish.yamul.backend.game.controller.domain.Hues
import dev.cavefish.yamul.backend.game.controller.domain.Notoriety
import dev.cavefish.yamul.backend.game.controller.infra.GameObjectRepository
import org.springframework.stereotype.Repository

@SuppressWarnings("MagicNumber")
@Repository
class InMemoryInitRepositories(
    private val gameObjectRepository: GameObjectRepository,
) {
    fun init() {
        val playerCharacterObject = GameObject(
            name = "John Doe",
            isCharacter = true,
            parentId = null,
            graphicId = GraphicId.BodyHumanMale,
            hue = Hues.Character.hue,
            flags = Flags.and(Flags.IgnoreMobiles, Flags.YellowBar),
            notoriety = Notoriety.and(Notoriety.Gray, Notoriety.Criminal),
        )
        val gameCharacterObjectId = gameObjectRepository.registerNewObject(playerCharacterObject)

        gameObjectRepository.registerNewObject(
            GameObject(
                parentId = gameCharacterObjectId,
                graphicId = GraphicId.Backpack,
                layer = 0x15,
                hue = Hues.Blue.hue,
            )
        )
        gameObjectRepository.registerNewObject(
            GameObject(
                parentId = gameCharacterObjectId,
                graphicId = GraphicId.HairShort,
                layer = 0x0B,
                hue = Hues.Black.hue,
            )
        )
        gameObjectRepository.registerNewObject(
            GameObject(
                parentId = gameCharacterObjectId,
                graphicId = GraphicId.RobeGm,
                layer = 0x16,
                hue = Hues.White.hue,
            )
        )

        InMemoryUserCharacterRepository.database["admin" to 0] = Character(
            objectId = gameCharacterObjectId,
            name = playerCharacterObject.name!!
        )

    }
}
