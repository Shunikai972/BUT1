package ihm.td3.ecouteurs

import ihm.td3.MaFenetre
import javafx.event.ActionEvent
import javafx.event.EventHandler
import javafx.scene.image.Image
import javafx.scene.image.ImageView

class BoutonAfficheImageEcouteur(private val appli: MaFenetre) : EventHandler<ActionEvent> {

    private var imageAffichee = false
    private val imageView = ImageView(Image("file:image/Mac.jpg"))

    override fun handle(event: ActionEvent?) {
        if (!imageAffichee) {
            imageView.fitWidth = 80.0
            imageView.fitHeight = 80.0
            imageView.isPreserveRatio = true

            appli.panneauBas.children.add(imageView)
            appli.boutonImage.text = "Cacher\nimage"

            imageAffichee = true
        } else {
            appli.panneauBas.children.remove(imageView)
            appli.boutonImage.text = "Afficher\n image"

            imageAffichee = false
        }
    }
}