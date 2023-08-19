package dev.cavefish.yamul.backend.login.controller


import dev.cavefish.yamul.backend.login.api.LoginRequest
import dev.cavefish.yamul.backend.login.api.LoginResponse
import dev.cavefish.yamul.backend.login.api.LoginServiceGrpc
import dev.cavefish.yamul.backend.utils.StringUtils
import io.grpc.stub.StreamObserver

class LoginServiceController : LoginServiceGrpc.LoginServiceImplBase() {

    override fun validateLogin(request: LoginRequest?, responseObserver: StreamObserver<LoginResponse>?) {
        val username = request?.username?.run(StringUtils::trimZeros)
        val result = if (username.equals("admin")) {
            println("Valid login $username")
            LoginResponse.LoginResponseValue.valid
        } else {
            println("Invalid login $username")
            LoginResponse.LoginResponseValue.invalid
        }
        responseObserver?.onNext(LoginResponse.newBuilder().setValue(result).build())
        responseObserver?.onCompleted()
    }
}
