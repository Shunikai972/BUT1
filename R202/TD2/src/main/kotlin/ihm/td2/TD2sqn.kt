package ihm.td2

import javafx.application.Application
import javafx.scene.Scene
import javafx.scene.control.Button
import javafx.scene.layout.BorderPane
import javafx.scene.layout.VBox
import javafx.stage.Stage


class TD2_1: Application() {

    override fun start(primaryStage: Stage) {
        val root2 = BorderPane()
        val leftBox = VBox(10.0)
        leftBox.style = "-fx-padding: 10; -fx-alignment: TOP_CENTER;"
        for (i in 1..6) {
            val btn = Button(i.toString())
            btn.prefWidth = 60.0
            btn.prefHeight = 40.0
            leftBox.children.add(btn)


            root2.left = leftBox

            val scene = Scene(root2, 400.0, 300.0)
            primaryStage.title = "TP2.1 en javaFX"
            primaryStage.scene = scene
            primaryStage.show()
        }

    }
}


fun main() {
    Application.launch(TD2_1::class.java)
}

