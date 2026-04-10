package ihm.td2

import javafx.application.Application
import javafx.geometry.Pos
import javafx.scene.Scene
import javafx.scene.control.Button
import javafx.scene.layout.BorderPane
import javafx.scene.layout.HBox
import javafx.scene.layout.VBox
import javafx.stage.Stage


class TD2_12: Application() {

    override fun start(primaryStage: Stage) {
        val root = BorderPane()
        
        // Left side - numbered buttons (1-6)
        val leftBox = VBox(10.0)
        leftBox.style = "-fx-padding: 10; -fx-alignment: TOP_CENTER;"
        for (i in 1..6) {
            val btn = Button(i.toString())
            btn.prefWidth = 60.0
            btn.prefHeight = 40.0
            leftBox.children.add(btn)
        }
        
        // Center - activation button
        val centerBox = VBox()
        centerBox.alignment = Pos.CENTER
        val activationBtn = Button("Activation")
        activationBtn.prefWidth = 100.0
        activationBtn.prefHeight = 40.0
        centerBox.children.add(activationBtn)
        
        // Right side - navigation buttons
        val rightBox = VBox(10.0)
        rightBox.style = "-fx-padding: 10; -fx-alignment: TOP_CENTER;"
        val upBtn = Button("∧")
        upBtn.prefWidth = 40.0
        val leftArrowBtn = Button("<")
        leftArrowBtn.prefWidth = 40.0
        val rightArrowBtn = Button(">")
        rightArrowBtn.prefWidth = 40.0
        val downBtn = Button("∨")
        downBtn.prefWidth = 40.0
        
        rightBox.children.addAll(upBtn, leftArrowBtn, rightArrowBtn, downBtn)
        
        root.left = leftBox
        root.center = centerBox
        root.right = rightBox
        
        val scene = Scene(root, 400.0, 300.0)
        primaryStage.title = "TP2.1 en javaFX"
        primaryStage.scene = scene
        primaryStage.show()
    }
}


fun main() {
    Application.launch(TD2_12::class.java)
}

