package ihm.td4.heureMVC.controleur

import ihm.td4.heureMVC.modele.Model
import ihm.td4.heureMVC.vue.Vue
import javafx.event.ActionEvent
import javafx.event.EventHandler



class ControleurBoutonHeure(modele: Model, vue: Vue): EventHandler<ActionEvent> {

    private val modele: Model
    private val vue: Vue

    init {
        this.modele=modele
        this.vue=vue
    }


    override  fun handle(event: ActionEvent) {
        this.vue.setTexte(this.modele.getTime())
        }

    }

