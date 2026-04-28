package ihm.td3

import ihm.td3.ecouteurs.BoutonGoEcouteur
import ihm.td3.ecouteurs.ChoixCouleurEcouteur
import ihm.td3.ecouteurs.ClicPanneauEcouteur
import ihm.td3.ecouteurs.BoutonAfficheImageEcouteur



import javafx.application.Application
import javafx.geometry.Insets
import javafx.geometry.Pos
import javafx.scene.Scene
import javafx.scene.control.*
import javafx.scene.layout.BorderPane
import javafx.scene.layout.FlowPane
import javafx.scene.layout.Pane
import javafx.scene.text.TextAlignment
import javafx.stage.Stage


class MaFenetre: Application() {

	val labelNbClicBouton:Label
	val labelNbClicPanneau: Label
	val panneauCouleurs: Pane
	val zoneTexte: TextField
	val textArea: TextArea
	val panneauHaut: FlowPane
	val panneauBas: FlowPane
	val choixCouleur: ComboBox<String>
	val boutonImage: Button

	init{
		this.labelNbClicBouton=Label()
		this.labelNbClicPanneau=Label()
		this.panneauCouleurs = Pane()
		this.zoneTexte=TextField("")
		this.textArea=TextArea()
		this.panneauHaut= FlowPane()
		this.panneauBas= FlowPane()
		this.choixCouleur=ComboBox<String>()
		this.boutonImage=Button("")
	}

	override fun start(primaryStage: Stage) {
		panneauCouleurs.style = "-fx-background-color: blue;"
		val couleurspossible = arrayOf("Bleu", "Vert", "Rouge")
		choixCouleur.items.addAll(*couleurspossible)
		choixCouleur.selectionModel.select(0)
		choixCouleur.onAction = ChoixCouleurEcouteur(this)
		
		val go = Button("go")
		go.onAction = BoutonGoEcouteur(this)

		panneauCouleurs.onMouseClicked = ClicPanneauEcouteur(this)
		boutonImage.onAction = BoutonAfficheImageEcouteur(this)
		boutonImage.text="Afficher\nimage"
		boutonImage.textAlignment= TextAlignment.CENTER
		
		zoneTexte.prefColumnCount = 15
		textArea.isWrapText = true
		textArea.isDisable = true
		textArea.prefRowCount = 10
		textArea.prefColumnCount = 10

		panneauHaut.alignment = Pos.TOP_CENTER
		panneauHaut.hgap = 20.0
		panneauHaut.padding= Insets(10.0)
		panneauHaut.children.addAll(boutonImage,go,choixCouleur, zoneTexte)

		panneauBas.hgap = 10.0
		panneauBas.alignment = Pos.TOP_CENTER
		panneauBas.padding= Insets(10.0)
		
		labelNbClicBouton.text = "0"
		labelNbClicPanneau.text = "0"
		val labelCliquer = Label("Clics sur \"go\" / panneau coloré = ")
		panneauBas.children.addAll(labelCliquer, labelNbClicBouton, labelNbClicPanneau)
		textArea.style="-fx-text-fill: black"
                
		val root = BorderPane()
		root.top = panneauHaut
		root.center = panneauCouleurs
		root.bottom = panneauBas
		root.left = textArea
		val scene = Scene(root, 800.0, 500.0)
		primaryStage.title="TD3 en javaFX"
		primaryStage.scene=scene
		primaryStage.show()
	}

}

fun main() {
	Application.launch(MaFenetre::class.java)
}
