package ihm.td4.calculatriceMVC.vue

import javafx.event.ActionEvent
import javafx.event.EventHandler
import javafx.geometry.Insets
import javafx.scene.control.Button
import javafx.scene.control.TextField
import javafx.scene.layout.ColumnConstraints
import javafx.scene.layout.GridPane
import javafx.scene.layout.RowConstraints
import javafx.scene.paint.Color

class Vue: GridPane() {

    private val textField: TextField
    private val tabButton: ArrayList<Button>

    init {
        val labelsBouton = arrayOf("0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-", "+", "*", "/", "=", "C")
        this.textField = TextField("")
        this.textField.setEditable(false)
        this.textField.setPrefSize(10.0, 50.0)

        this.tabButton = ArrayList<Button>(14)
        // création des boutons de la calculatrice
        for (i in 0..labelsBouton.size - 1) {
            val button = Button(labelsBouton[i])
            button.setMaxSize(Double.MAX_VALUE, Double.MAX_VALUE)
            button.textFill= Color.BLUEVIOLET
            this.tabButton.add(button)
        }
        // ajout des boutons dans le gridpane
        this.add(textField, 0, 0, 4, 1)
        this.addRow(1,tabButton[7], tabButton[8], tabButton[9],tabButton[10] )
        this.addRow(2,tabButton[4], tabButton[5], tabButton[6],tabButton[11] )
        this.addRow(3,tabButton[1], tabButton[2], tabButton[3],tabButton[12] )
        this.addRow(4,tabButton[0], tabButton[13], tabButton[14],tabButton[15] )
        // mise en place des contraintes sur les lignes et colonnes du gridpane
        val columnConstraint = ColumnConstraints()
        columnConstraint.percentWidth = 25.00
        val rowConstraint = RowConstraints()
        rowConstraint.percentHeight = 25.00
        val tabColumnConstraints= Array(4){columnConstraint}
        this.columnConstraints.addAll(tabColumnConstraints)
        val tabRowConstraints= Array(5){rowConstraint}
        this.rowConstraints.addAll(tabRowConstraints)

        this.vgap=10.0
        this.hgap=10.0
        this.padding= Insets(10.0)
        this.style="-fx-background-color: lightblue"
    }
    // pour valuer le texte de TextField
    fun setTexte(s: String){
        this.textField.text=s
    }

    // pour récupérer le texte de textField
    fun getTexte():String=this.textField.text

    // pour concaténer s au texte de textField
    fun ajoutTexte(s: String){
       this.textField.text+=s
    }

    fun fixeControleurBoutonChiffreOperateur(controleur: EventHandler<ActionEvent>) {
        for (i in 0..13) {
            this.tabButton[i].onAction = controleur
        }
    }

    fun fixeControleurResultat(controleur: EventHandler<ActionEvent>) {
        this.tabButton[14].onAction = controleur
    }

    fun fixeControleurEfface(controleur: EventHandler<ActionEvent>) {
        this.tabButton[15].onAction = controleur
    }

}
