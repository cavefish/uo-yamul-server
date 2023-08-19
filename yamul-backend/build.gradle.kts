plugins {
    kotlin("jvm") version "1.9.0"
    id("com.google.protobuf") version "0.9.4"
    // id("org.jlleitschuh.gradle.ktlint") version "11.3.2"
    application
}

group = "dev.cavefish.yamul"
version = "1.0-SNAPSHOT"
val grpcVersion = "1.54.1"
val grpcKotlinVersion = "1.3.0" // CURRENT_GRPC_KOTLIN_VERSION
val protobufVersion = "3.24.0"
val coroutinesVersion = "1.7.0"

repositories {
    mavenCentral()
    google()
}

dependencies {
    runtimeOnly("io.grpc:grpc-netty:$grpcVersion")

    api("org.jetbrains.kotlinx:kotlinx-coroutines-core:$coroutinesVersion")

    api("io.grpc:grpc-protobuf:$grpcVersion")
    api("com.google.protobuf:protobuf-kotlin:$protobufVersion")
    api("io.grpc:grpc-kotlin-stub:$grpcKotlinVersion")

    testImplementation(kotlin("test"))
}

tasks.test {
    useJUnitPlatform()
}

kotlin {
    jvmToolchain(8)
}

protobuf {
    protoc {
        artifact = "com.google.protobuf:protoc:$protobufVersion"
    }
    plugins {
        create("grpc") {
            artifact = "io.grpc:protoc-gen-grpc-java:$grpcVersion"
        }
        create("grpckt") {
            artifact = "io.grpc:protoc-gen-grpc-kotlin:$grpcKotlinVersion:jdk8@jar"
        }
    }
    generateProtoTasks {
        all().forEach {
            it.plugins {
                create("grpc")
                create("grpckt")
            }
            it.builtins {
                create("kotlin")
            }
        }
    }
}

java.sourceSets["main"].proto {
    srcDirs("../api-definitions/backend/")
}

tasks.register<JavaExec>("RunLoginService") {
    dependsOn("classes")
    classpath = sourceSets["main"].runtimeClasspath
    mainClass.set("dev.cavefish.yamul.backend.login.controller.LoginServiceMain")
}