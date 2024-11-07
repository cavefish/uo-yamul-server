echo "Building dependencies"
docker run --rm `
    -v ${pwd}/..\api-definitions\dashboard\proto:/proto `
    -v ${pwd}\lib\generated\grpc:/grpc `
    rvolosatovs/protoc `
    /proto/*.proto --proto_path=/proto --dart_out=grpc:/grpc/ -Iproto
