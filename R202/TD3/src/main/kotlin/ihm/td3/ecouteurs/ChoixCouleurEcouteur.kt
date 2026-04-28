package ihm.td3.ecouteurs

import ihm.td3.MaFenetre
import javafx.event.ActionEvent
import javafx.event.EventHandler


class ChoixCouleurEcouteur(private val appli: MaFenetre) : EventHandler<ActionEvent> {

    override fun handle(event: ActionEvent?) {
        val couleur = appli.choixCouleur.value
        appli.panneauCouleurs.style = when (couleur) {
            "Bleu" -> "-fx-background-color: blue;"
            "Vert" -> "-fx-background-color: green;"
            "Rouge" -> "-fx-background-color: red;"
            else -> ""
        }
    }
}
