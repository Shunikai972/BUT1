package ihm.td1

import javafx.application.Application
import javafx.scene.Scene
import javafx.scene.layout.BorderPane
import javafx.stage.Stage


class TD1_1(): Application(){

    init {
        print("je suis dans le constructeur")
    }

        override fun init() {
                print("je suis dans la méthode x")
        }

        override fun start(stage : Stage) {

                stage.setTitle("GROS");
                stage.isResizable = false
                stage.alwaysOnTopProperty()
                val scene = Scene(BorderPane(), 1000.0, 800.0)
                stage.scene=scene
                stage.show()
        }

}

fun main() {

        println("main avant launch")
        Application.launch(TD1_1::class.java)
        println("main après launch")
}




