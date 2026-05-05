package ihm.td4.calculatriceMVC.controleur

import ihm.td4.calculatriceMVC.modele.Model
import ihm.td4.calculatriceMVC.tools.Expression
import ihm.td4.calculatriceMVC.vue.Vue
import javafx.event.ActionEvent
import javafx.event.EventHandler

class ControleurResultat(private val modele: Model, private val vue: Vue): EventHandler<ActionEvent> {
    override fun handle(event: ActionEvent) {
        val expression = this.vue.getTexte()

        if (!Expression("$expression=").isValidExpression()) {
            println("Expression incomplète ou incorrecte : $expression")
            this.vue.setTexte("")
            return
        }

        try {
            val match = Regex("[+\\-*/]").find(expression)
            if (match == null) {
                println("Aucun opérateur trouvé : $expression")
                this.vue.setTexte("")
                return
            }

            val operateur = match.value
            val parties = expression.split(Regex("[+\\-*/]"))
            val operande1 = parties[0].toInt()
            val operande2 = parties[1].toInt()
            val resultat = this.modele.calculer(operande1, operateur, operande2)

            this.vue.setTexte("$expression=$resultat")
        } catch (exception: ArithmeticException) {
            println("Division par zéro impossible")
            this.vue.setTexte("")
        } catch (exception: NumberFormatException) {
            println("Expression incorrecte : $expression")
            this.vue.setTexte("")
        }
    }
}










