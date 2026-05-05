package ihm.td4.calculatriceMVC.controleur

import ihm.td4.calculatriceMVC.tools.Expression
import ihm.td4.calculatriceMVC.vue.Vue
import javafx.event.ActionEvent
import javafx.event.EventHandler
import javafx.scene.control.Button

class ControleurBoutonChiffreOperateur(private val vue: Vue): EventHandler<ActionEvent> {
    override fun handle(event: ActionEvent) {
        val bouton = event.source as Button
        val nouvelleExpression = this.vue.getTexte() + bouton.text

        if (Expression(nouvelleExpression).isValidExpression()) {
            this.vue.setTexte(nouvelleExpression)
        } else {
            println("Expression incorrecte : $nouvelleExpression")
            this.vue.setTexte("")
        }
    }
}
