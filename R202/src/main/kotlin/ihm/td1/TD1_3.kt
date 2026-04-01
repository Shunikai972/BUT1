package ihm.td1

import javafx.application.Application
import javafx.scene.Scene
import javafx.scene.canvas.Canvas
import javafx.scene.canvas.GraphicsContext
import javafx.stage.Stage

class TD1_3: Application() {

    override fun start(stage: Stage) {

        val canvas = Canvas(600.0, 400.0)
        val g: GraphicsContext = canvas.getGraphicsContext2D()
        var tableau = Tableau()
        tableau.peindre(g)
        tableau.children.add(canvas)
        val scene = Scene(tableau, canvas.width, canvas.height)
        stage.title="TD1_3 en javaFX"
        stage.scene=scene
        stage.show();
    }
}
    fun main() {
        Application.launch(TD1_3::class.java)

    }

