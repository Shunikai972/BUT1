package ihm.td4.calculatriceMVC


import ihm.td4.calculatriceMVC.controleur.ControleurBoutonChiffreOperateur
import ihm.td4.calculatriceMVC.controleur.ControleurEfface
import ihm.td4.calculatriceMVC.controleur.ControleurResultat
import ihm.td4.calculatriceMVC.modele.Model
import ihm.td4.calculatriceMVC.vue.Vue
import javafx.application.Application
import javafx.scene.Scene
import javafx.stage.Stage


class Calculatrice: Application(){

        override fun start(primaryStage: Stage) {
            val vue = Vue()
            val modele = Model()

            vue.fixeControleurBoutonChiffreOperateur(ControleurBoutonChiffreOperateur(vue))
            vue.fixeControleurResultat(ControleurResultat(modele, vue))
            vue.fixeControleurEfface(ControleurEfface(vue))

            val scene = Scene(vue)
            primaryStage.title = "Calculatrice"
            primaryStage.scene = scene
            primaryStage.show()
            

        }
    }


        fun main(){
            Application.launch(Calculatrice::class.java)
        }
