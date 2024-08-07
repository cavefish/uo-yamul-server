package dev.cavefish.yamul.backend.auth.controller

import dev.cavefish.yamul.backend.Constants
import dev.cavefish.yamul.backend.game.controller.domain.LoggedUser
import io.grpc.Context
import io.grpc.Contexts
import io.grpc.Metadata
import io.grpc.Metadata.ASCII_STRING_MARSHALLER
import io.grpc.ServerCall
import io.grpc.ServerCallHandler
import io.grpc.ServerInterceptor
import io.grpc.Status
import io.grpc.StatusRuntimeException
import org.springframework.stereotype.Component
import java.util.*

@Component
class BasicAuthInterceptor : ServerInterceptor {

    private val authorizationHeaderKey: Metadata.Key<String> by lazy {
        Metadata.Key.of("x-auth-key", ASCII_STRING_MARSHALLER)
    }

    override fun <ReqT : Any?, RespT : Any?> interceptCall(
        call: ServerCall<ReqT, RespT>?, headers: Metadata?, next: ServerCallHandler<ReqT, RespT>?
    ): ServerCall.Listener<ReqT> {
        Objects.requireNonNull(call)
        Objects.requireNonNull(headers)
        Objects.requireNonNull(next)

        val authorizationHeader = headers?.get(authorizationHeaderKey)
        val loggedUser = getValidLoggedUser(authorizationHeader)
        if (loggedUser != null) {
            val ctx = Context.current().withValue(Constants.AUTH_CONTEXT_LOGGED_USER, loggedUser)
            return Contexts.interceptCall(ctx, call, headers, next)
        }
        throw StatusRuntimeException(Status.UNAUTHENTICATED)
    }

    private fun getValidLoggedUser(authorizationHeader: String?): LoggedUser? {
        if (authorizationHeader == null) return null
        val parts = authorizationHeader.split(' ')
        if (parts.size != 2) return null
        if (checkUser(parts[0], parts[1])) {
            return LoggedUser(parts[0])
        }
        return null
    }

    private fun checkUser(user: String, password: String): Boolean {
        if (user != "admin") return false
        if (password != "admin") return false
        return true
    }

}
