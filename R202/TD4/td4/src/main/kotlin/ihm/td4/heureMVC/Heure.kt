package ihm.td4.heureMVC


import ihm.td4.heureMVC.modele.Model
import ihm.td4.heureMVC.controleur.ControleurBoutonHeure
import ihm.td4.heureMVC.vue.Vue
import javafx.application.Application
import javafx.scene.Scene
import javafx.stage.Stage


class Heure: Application(){

        override fun start(primaryStage: Stage) {
            val vue= Vue()
            val modele= Model()
            val controleurBoutonHeure = ControleurBoutonHeure(modele, vue)
            vue.fixeControleurBoutonHeure(controleurBoutonHeure)
            val scene = Scene(vue, 250.0, 100.0)
            primaryStage.title="Affichage heure"
            primaryStage.scene=scene
            primaryStage.show()
        }
    }


        fun main(){
            Application.launch(Heure::class.java)
        }
