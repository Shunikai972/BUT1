package ihm.td1

import javafx.application.Application
import javafx.scene.Scene
import javafx.scene.control.Button
import javafx.scene.control.ColorPicker
import javafx.scene.control.DatePicker
import javafx.scene.control.Label
import javafx.scene.control.Slider
import javafx.scene.image.Image
import javafx.scene.image.ImageView
import javafx.scene.layout.VBox
import javafx.scene.text.Font
import javafx.scene.text.FontPosture
import javafx.scene.text.FontWeight
import javafx.stage.Stage
import java.io.FileInputStream


class TD1_2 : Application() {


    override fun start(stage : Stage) {

        stage.setTitle("Premiere application");
        stage.isResizable = false
        stage.alwaysOnTopProperty()
        val root = VBox()
        val monLabel = Label("mon premier programme en javaFX")
        monLabel.font = Font.font("Verdana", FontWeight.BOLD, FontPosture.REGULAR, 20.0)
        root.children.add(monLabel)
        stage.scene= Scene(root, 400.0, 300.0)

        val boutonb =  Button("bouton1")
        boutonb.font = Font.font("Verdana", FontPosture.REGULAR, 20.0)
        boutonb.textFill = javafx.scene.paint.Color.RED

        val bouton2= Button("bouton2")
        bouton2.font = Font.font("Verdana", FontPosture.REGULAR, 10.0)
        bouton2.textFill = javafx.scene.paint.Color.RED
        bouton2.style = "-fx-background-color: Aqua"




        val input = FileInputStream("image/fleur.jpg")
        val image = Image(input)
        val bouton3= Button("bouton3",
        ImageView(image))


        val monLabel4 = DatePicker()

        val monLabel5 = ColorPicker(javafx.scene.paint.Color.PINK)


        val monlabel6 = Slider(0.0,100.0,25.0)
        monlabel6.setShowTickLabels(true)
        monlabel6.setBlockIncrement(0.1)
        monlabel6.setShowTickMarks(true)
        root.children.add(boutonb)
        root.children.add(bouton2)
        root.children.add(bouton3)
        root.children.add(monLabel4)
        root.children.add(monLabel5)
        root.children.add(monlabel6)
        stage.show()
    }

}

fun main() {
    Application.launch(TD1_2::class.java)
}