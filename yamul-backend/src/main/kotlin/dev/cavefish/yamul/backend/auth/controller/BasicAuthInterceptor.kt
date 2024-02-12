package dev.cavefish.yamul.backend.auth.controller


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
        if (isValidAuthorization(authorizationHeader)) {
            return next!!.startCall(call, headers)
        }
        throw StatusRuntimeException(Status.UNAUTHENTICATED)
    }

    private fun isValidAuthorization(authorizationHeader: String?): Boolean {
        if (authorizationHeader == null) return false
        val parts = authorizationHeader.split(' ')
        if (parts.size != 2) return false
        return checkUser(parts[0], parts[1])
    }

    private fun checkUser(user: String, password: String): Boolean {
        if (user != "admin") return false
        if (password != "admin") return false
        return true
    }

}
