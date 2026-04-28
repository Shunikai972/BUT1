package ihm.td3.ecouteurs


import ihm.td3.MaFenetre
import javafx.event.ActionEvent
import javafx.event.EventHandler



class BoutonGoEcouteur(appli: MaFenetre) : EventHandler<ActionEvent> {
        private val appli: MaFenetre

        //--- Constructeur ---------------------------------
        init {
            this.appli=appli
        }

        //--- Code exécuté lorsque l'événement survient ----
       override  fun handle(event: ActionEvent) {
            appli.labelNbClicBouton.text = (appli.labelNbClicBouton.text.toInt()+1).toString()
        }
    }
