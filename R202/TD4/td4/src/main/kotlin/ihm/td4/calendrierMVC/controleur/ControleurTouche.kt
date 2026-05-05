package ihm.td4.calendrierMVC.controleur


import ihm.td4.calendrierMVC.modele.Calendrier
import ihm.td4.calendrierMVC.vue.Vue
import javafx.event.EventHandler
import javafx.scene.input.KeyCode
import javafx.scene.input.KeyEvent

class ControleurTouche(private val modele: Calendrier, private val vue: Vue): EventHandler<KeyEvent> {
   override fun handle(event: KeyEvent) {
      when (event.code) {
         KeyCode.LEFT -> this.modele.addDay(-1)
         KeyCode.RIGHT -> this.modele.addDay(1)
         else -> return
      }

      this.vue.update(
         this.modele.getYear(),
         this.modele.getMonth(),
         this.modele.getDayNumber(),
         this.modele.getDay()
      )
   }
}
