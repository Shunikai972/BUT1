package ihm.td4.calculatriceMVC.controleur

import ihm.td4.calculatriceMVC.vue.Vue
import javafx.event.ActionEvent
import javafx.event.EventHandler

class ControleurEfface(private val vue: Vue): EventHandler<ActionEvent> {
    override fun handle(event: ActionEvent) {
        this.vue.setTexte("")
    }

}
