package dev.cavefish.yamul.backend.infra.inmemory

import dev.cavefish.yamul.backend.game.controller.domain.Character
import dev.cavefish.yamul.backend.game.controller.domain.LoggedUser
import dev.cavefish.yamul.backend.game.controller.infra.UserCharacterRepository
import org.springframework.stereotype.Repository

@Repository
class InMemoryUserCharacterRepository: UserCharacterRepository {
    override fun getCharacterByOrder(user: LoggedUser, order: Int) = database[user.username to order]

    companion object {
        val database = HashMap<Pair<String, Int>, Character>()
    }
}
