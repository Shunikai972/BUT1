package ihm.td4.heureMVC.vue

import javafx.event.ActionEvent
import javafx.event.EventHandler
import javafx.geometry.Insets
import javafx.scene.control.Button
import javafx.scene.control.Label
import javafx.scene.layout.HBox


class Vue: HBox() {

    private val label: Label
    private val button: Button

    init {
        this.label = Label("")
        this.button = Button("Quelle heure ?")
        this.children.addAll(this.button, this.label)
        this.spacing = 10.0
        this.padding = Insets(20.0)
    }
    // pour valuer le texte du label
    fun setTexte(s: String){
        this.label.text = s
    }


// cette méthode sera développée dans la question 3
    fun fixeControleurBoutonHeure(controleur:  EventHandler<ActionEvent>) {

            this.button.onAction = controleur
    }


}
