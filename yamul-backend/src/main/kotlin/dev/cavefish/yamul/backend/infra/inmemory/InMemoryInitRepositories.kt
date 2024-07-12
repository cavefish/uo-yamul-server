package dev.cavefish.yamul.backend.infra.inmemory

import dev.cavefish.yamul.backend.game.controller.domain.Character
import dev.cavefish.yamul.backend.game.controller.domain.Coordinates
import dev.cavefish.yamul.backend.game.controller.domain.Flags
import dev.cavefish.yamul.backend.game.controller.domain.GameObject
import dev.cavefish.yamul.backend.game.controller.domain.GameObjectItem
import dev.cavefish.yamul.backend.game.controller.domain.GraphicId
import dev.cavefish.yamul.backend.game.controller.domain.Hues
import dev.cavefish.yamul.backend.game.controller.domain.Notoriety

@SuppressWarnings("MagicNumber")
object InMemoryInitRepositories {
    fun init() {
        InMemoryUserCharacterRepository.database["admin" to 0] = Character(
            id = 1,
            name = "John Doe"
        )
        InMemoryGameObjectRepository.database[1] = GameObject(
            id = 1,
            graphicId = GraphicId.BodyHumanMale,
            hue = Hues.White.hue,
            coordinates = Coordinates(x = 6787, y = 2181, z = 0),
            flags = Flags.None,
            notoriety = Notoriety.Innocent,
            items = listOf(
                GameObjectItem(2, GraphicId.Backpack, Hues.Blue.hue, 0x15),
                GameObjectItem(3, GraphicId.HairShort, Hues.Green.hue, 0x0B),
                GameObjectItem(4, GraphicId.RobeGm, Hues.White.hue, 0x16),
            )
        )
    }
}