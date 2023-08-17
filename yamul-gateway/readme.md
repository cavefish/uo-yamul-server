# Yamul Gateway

## How-to generate gRPC client code

````shell
protoc --go_out=. --go-grpc_out=. \
 ../api-definitions/backend/yamul-backend-login.proto \ 
 --proto_path=../api-definitions/backend 
````