package ihm.td5


import ihm.td5.modele.Bibliotheque
import ihm.td5.vue.MainVue
import javafx.application.Application
import javafx.scene.Scene

import javafx.stage.Stage

class Main: Application() {

    override fun start(primaryStage: Stage) {
       val vue = MainVue()
        val bibliotheque = Bibliotheque()
        bibliotheque.preremplir()
        vue.updateContenuPanneauDroit(bibliotheque.courant, bibliotheque.donneLivre())
        vue.updateContenuPanneauGauche(bibliotheque.donneTousLesLivres())
        val scene = Scene(vue, 550.0, 350.0)
        primaryStage.title="TD5"

        primaryStage.scene=scene
        primaryStage.show()
    }
}

fun main(){
    Application.launch(Main::class.java)
}