import io.gitlab.arturbosch.detekt.Detekt
import io.gitlab.arturbosch.detekt.DetektCreateBaselineTask

plugins {
    kotlin("jvm") version "1.9.0"
    id("com.google.protobuf") version "0.9.4"
    id("com.github.sherter.google-java-format") version "0.9"
    id("io.gitlab.arturbosch.detekt") version "1.23.1"
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

kotlin.sourceSets["main"].kotlin {
    srcDirs("src/main/kotlin")
}

tasks.register<JavaExec>("RunLoginService") {
    dependsOn("classes")
    classpath = sourceSets["main"].runtimeClasspath
    mainClass.set("dev.cavefish.yamul.backend.login.controller.LoginServiceMain")
}

tasks["build"].dependsOn(tasks["googleJavaFormat"])
tasks["verifyGoogleJavaFormat"].mustRunAfter(tasks["googleJavaFormat"])

detekt {
    buildUponDefaultConfig = true
    allRules = false
    config.setFrom("$projectDir/config/detekt.yml")
    baseline = file("$projectDir/config/baseline.xml")
}

tasks.withType<Detekt>().configureEach {
    reports {
        html.required.set(true)
        xml.required.set(true)
        txt.required.set(true)
        sarif.required.set(true)
        md.required.set(true)
    }
}

tasks.withType<Detekt>().configureEach {
    jvmTarget = "19"
}
tasks.withType<DetektCreateBaselineTask>().configureEach {
    jvmTarget = "19"
}