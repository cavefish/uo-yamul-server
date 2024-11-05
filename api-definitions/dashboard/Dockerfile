FROM ubuntu:latest AS build-extensions
RUN apt-get update && apt-get install -y wget
RUN mkdir extensions && cd extensions
RUN wget https://repo1.maven.org/maven2/org/wiremock/wiremock-grpc-extension-standalone/0.8.1/wiremock-grpc-extension-standalone-0.8.1.jar


FROM rvolosatovs/protoc AS build-protoc
RUN mkdir grpc && mkdir proto
COPY proto/*.proto proto/
RUN protoc ./proto/*.proto --proto_path=./proto --descriptor_set_out ./grpc/services.dsc


FROM holomekc/wiremock-gui:latest
COPY --from=build-extensions extensions/* extensions/
COPY --from=build-protoc grpc/* grpc/

ENTRYPOINT ["/docker-entrypoint.sh", "--global-response-templating", "--disable-gzip", "--verbose"]