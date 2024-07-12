package dev.cavefish.yamul.backend.game.controller.infra

import dev.cavefish.yamul.backend.game.controller.domain.LoggedUser
import dev.cavefish.yamul.backend.game.controller.domain.Character

interface UserCharacterRepository {
    fun getCharacterByOrder(user: LoggedUser, order: Int): Character?
}