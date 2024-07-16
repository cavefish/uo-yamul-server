package dev.cavefish.yamul.backend

import dev.cavefish.yamul.backend.game.controller.domain.LoggedUser
import io.grpc.Context

object Constants {
    val AUTH_CONTEXT_LOGGED_USER: Context.Key<LoggedUser> = Context.key("AUTH_CONTEXT_LOGGED_USER")
}
