package ihm.td5.controleur

import ihm.td5.modele.Bibliotheque
import ihm.td5.vue.MainVue
import javafx.event.EventHandler
import javafx.scene.input.MouseEvent

class ControleurDetailLivre(Modele : Bibliotheque, vue : MainVue) : EventHandler<MouseEvent>{
    val bibliotheque: Bibliotheque
    val vue2 : MainVue

    init {
        this.vue2 = vue
        this.bibliotheque = Modele
    }

}
