package ihm.td4.calendrierMVC.controleur

import ihm.td4.calendrierMVC.modele.Calendrier
import ihm.td4.calendrierMVC.vue.Vue
import javafx.event.ActionEvent
import javafx.event.EventHandler

class ControleurComboBox(private val modele: Calendrier, private val vue: Vue): EventHandler<ActionEvent> {
  override fun handle(event: ActionEvent) {
    val moisSelectionne = this.vue.getIndexSelected()
    if (moisSelectionne < 0) {
      return
    }

    this.modele.moveMonth(moisSelectionne)
    this.vue.update(
      this.modele.getYear(),
      this.modele.getMonth(),
      this.modele.getDayNumber(),
      this.modele.getDay()
    )
  }
}

