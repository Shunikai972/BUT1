package ihm.td4.calendrierMVC


import ihm.td4.calendrierMVC.controleur.ControleurComboBox
import ihm.td4.calendrierMVC.controleur.ControleurTouche
import ihm.td4.calendrierMVC.modele.Calendrier
import ihm.td4.calendrierMVC.vue.Vue
import javafx.application.Application
import javafx.scene.Scene
import javafx.stage.Stage


class Main: Application() {


   override fun start(primaryStage: Stage) {
       val modele = Calendrier()
       val vue = Vue()

       vue.populateMonth(modele.giveAllMonths())
       vue.update(modele.getYear(), modele.getMonth(), modele.getDayNumber(), modele.getDay())
       vue.subscribeControleurComboBox(ControleurComboBox(modele, vue))
       vue.subscribeControleurTouche(ControleurTouche(modele, vue))
       vue.isFocusTraversable = true

       val scene = Scene(vue, 360.0, 90.0)
       primaryStage.title = "Calendrier"
       primaryStage.scene = scene
       primaryStage.show()
       vue.requestFocus()
   }

}

fun main(){
    Application.launch(Main::class.java)
}
