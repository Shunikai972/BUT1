package ihm.td3.ecouteurs

import ihm.td3.MaFenetre
import javafx.event.EventHandler
import javafx.scene.input.MouseEvent

class ClicPanneauEcouteur(private val appli: MaFenetre) : EventHandler<MouseEvent> {

    override fun handle(event: MouseEvent?) {
        val nbClics = appli.labelNbClicPanneau.text.toInt()
        appli.labelNbClicPanneau.text = (nbClics + 1).toString()
    }
}