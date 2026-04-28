import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
    kotlin("jvm") version "2.0.21"
    id("org.openjfx.javafxplugin") version "0.1.0"
    application
}

group = "jacquin-c"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}


javafx {
    version = "17"
    modules("javafx.controls", "javafx.fxml")
}


tasks.withType<KotlinCompile> {
    kotlinOptions.jvmTarget = "21"
}
