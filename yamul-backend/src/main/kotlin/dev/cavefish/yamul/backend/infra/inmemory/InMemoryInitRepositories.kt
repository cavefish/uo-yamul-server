package dev.cavefish.yamul.backend.infra.inmemory

import dev.cavefish.yamul.backend.game.controller.domain.Character
import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.Flags
import dev.cavefish.yamul.backend.game.controller.domain.GameObject
import dev.cavefish.yamul.backend.game.controller.domain.GameObjectTree
import dev.cavefish.yamul.backend.game.controller.domain.GameObjectTree.Companion.treeOf
import dev.cavefish.yamul.backend.game.controller.domain.GraphicId
import dev.cavefish.yamul.backend.game.controller.domain.Hues
import dev.cavefish.yamul.backend.game.controller.domain.Notoriety
import dev.cavefish.yamul.backend.game.controller.domain.ObjectId
import dev.cavefish.yamul.backend.game.controller.infra.GameObjectRepository
import dev.cavefish.yamul.backend.game.controller.infra.mul.MulBlockAltitudeRepository
import kotlinx.coroutines.runBlocking
import org.springframework.boot.context.event.ApplicationStartedEvent
import org.springframework.context.event.EventListener
import org.springframework.scheduling.annotation.Async
import org.springframework.scheduling.annotation.EnableAsync
import org.springframework.stereotype.Service
import org.tinylog.kotlin.Logger

@SuppressWarnings("MagicNumber")
@Service
@EnableAsync
class InMemoryInitRepositories(
    private val objectRepository: GameObjectRepository,
    private val realtimePosition: InMemoryGameObjectRealtimePosition,
    private val mulBlockAltitudeRepository: MulBlockAltitudeRepository,
) {
    @Async
    @EventListener(ApplicationStartedEvent::class)
    fun init() {

        runBlocking {
            Logger.info("Initializing repositories")

            InMemoryUserCharacterRepository.database["admin" to 0] = registerGameCharacter(
                GameObject(
                    name = "John Doe",
                    isCharacter = true,
                    parentId = null,
                    graphicId = GraphicId.BodyHumanMale,
                    hue = Hues.Character.hue,
                    flags = listOf(Flags.Normal, Flags.CanAlterPaperDoll),
                    notoriety = listOf(Notoriety.Gray, Notoriety.Criminal),
                ),
                mulBlockAltitudeRepository.correctPositionAltitude(feluccaFortIslandCenter),
                treeOf(
                    GameObject(
                        parentId = null,
                        graphicId = GraphicId.Backpack,
                        layer = 0x15,
                        hue = Hues.Blue.hue,
                    )
                ),
                treeOf(
                    GameObject(
                        parentId = null,
                        graphicId = GraphicId.HairShort,
                        layer = 0x0B,
                        hue = Hues.Black.hue,
                    )
                ),
                treeOf(
                    GameObject(
                        parentId = null,
                        graphicId = GraphicId.RobeGm,
                        layer = 0x16,
                        hue = Hues.White.hue,
                    )
                )

            )

            Logger.info("Repositories initialized")
        }
    }

    private suspend fun registerGameCharacter(
        characterObject: GameObject,
        position: Coordinates,
        vararg items: GameObjectTree
    ): Character {
        assert(characterObject.name != null)
        val gameCharacterObjectId = objectRepository.registerNewObject(characterObject)
        assert(
            realtimePosition.registerNewCoordinates(
                gameCharacterObjectId, position
            )
        )
        items.forEach { registerGameObjectTree(gameCharacterObjectId, it) }
        return Character(objectId = gameCharacterObjectId, name = characterObject.name!!)
    }

    private fun registerGameObjectTree(parentId: ObjectId, objectTree: GameObjectTree) {
        val newObjectId = objectRepository.registerNewObject(objectTree.value.copy(parentId = parentId))
        objectTree.children.forEach { child -> registerGameObjectTree(newObjectId, child) }
    }

    companion object {
        val feluccaStrangeTownOutOfBoundsCenter = Coordinates(
            x = 6787,
            y = 2181,
            z = 0,
            mapId = 1,
        )
        val feluccaFortIslandCenter = Coordinates(
            x = 2980,
            y = 3436,
            z = 15,
            mapId = 1,
        )
    }
}
