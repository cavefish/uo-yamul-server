package dev.cavefish.yamul.backend.login.controller


import dev.cavefish.yamul.backend.login.api.LoginRequest
import dev.cavefish.yamul.backend.login.api.LoginResponse
import dev.cavefish.yamul.backend.login.api.LoginServiceGrpc
import dev.cavefish.yamul.backend.utils.StringUtils
import io.grpc.stub.StreamObserver
import org.springframework.stereotype.Component
import org.tinylog.Logger

@Component
class LoginServiceController : LoginServiceGrpc.LoginServiceImplBase() {

    override fun validateLogin(request: LoginRequest?, responseObserver: StreamObserver<LoginResponse>?) {
        val username = request?.username?.run(StringUtils::trimZeros)!!
        val result = if (username == "admin") {
            Logger.info("Valid login {}", username)
            LoginResponse.LoginResponseValue.valid
        } else {
            Logger.info("Invalid login {}", username)
            LoginResponse.LoginResponseValue.invalid
        }
        responseObserver?.onNext(LoginResponse.newBuilder().setValue(result).build())
        responseObserver?.onCompleted()
    }

}
