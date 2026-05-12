package ihm.td5.vue

import ihm.td5.librairie.Auteur
import ihm.td5.librairie.Livre
import javafx.event.ActionEvent
import javafx.event.EventHandler
import javafx.geometry.Insets
import javafx.geometry.Pos
import javafx.scene.control.*
import javafx.scene.input.MouseEvent
import javafx.scene.layout.BorderPane
import javafx.scene.layout.FlowPane
import javafx.scene.layout.GridPane
import javafx.scene.paint.Color
import javafx.scene.text.Font
import javafx.scene.text.FontPosture
import javafx.scene.text.FontWeight

/**
 * classe développée par jacquin-c en 2022
 */

class MainVue: BorderPane() {

  private val panneauGauche: GridPane
  private var panneauDroit : TitledPaneLivre
   
   val boutonModification: Button
   val boutonAjout: Button
   val boutonSuppression: Button



    init {
        this.panneauGauche = GridPane()
        this.panneauDroit = TitledPaneLivre("information Livre")
        this.boutonModification = Button("Modification livre")
        this.boutonAjout = Button("Ajout livre")
        this.boutonSuppression = Button("Suppression livre")

        // le titre en haut
        val flowPaneTop = FlowPane()
        val labelTop = Label("Ma super Bibliothèque")
        labelTop.font = Font.font("Tahoma", FontWeight.BOLD, FontPosture.REGULAR, 20.0)
        labelTop.textFill = Color.BLACK
        flowPaneTop.alignment = Pos.CENTER
        labelTop.padding = Insets(10.0)
        flowPaneTop.children.add(labelTop)
        this.top = flowPaneTop

        // le panneau gauche qui contient la liste des livres qui est dans un TitledPane
        val scrollPane = ScrollPane(this.panneauGauche)
        val titledPaneLivres = TitledPane(" liste des livres", scrollPane)
        titledPaneLivres.alignment = Pos.CENTER
        titledPaneLivres.font = Font.font("sans_serif", FontWeight.BOLD, FontPosture.REGULAR, 12.0)
        titledPaneLivres.isCollapsible = false
        titledPaneLivres.padding = Insets(10.0)
        this.left = titledPaneLivres

       // le panneau droit qui contiendra un TitledPane dont le contenu se modifiera en fonction
        // des actions de l'utilisateur: visualiser un livre, le modifier, ajouter un livre
        setAlignment(panneauDroit, Pos.TOP_CENTER)
        this.center = panneauDroit
         // les trois boutons en bas pour lancer des actions de modification, ajout ou suppression d'un livre
        val flowPaneBoutons=FlowPane()
        flowPaneBoutons.hgap=10.0
        flowPaneBoutons.padding=Insets(10.0)
        flowPaneBoutons.children.addAll(boutonModification,boutonAjout,boutonSuppression)
        this.bottom=flowPaneBoutons

    }

    /**
     * permet de modifier le TitledPane qui correspond au panneau droit qui se trouve au centre de la vue
     * @param titledPaneLivre le nouveau panneau de type TitledPane
     */
    fun updatePanneauDroit(titledPaneLivre: TitledPaneLivre){
        this.panneauDroit=titledPaneLivre
        this.center=titledPaneLivre
        setAlignment(titledPaneLivre, Pos.TOP_CENTER)
    }

    /**
     * permet de mettre à jour la liste des livres dans le panneau gauche
     * et colore en bleu le livre dont l'index est index dans la liste
     * (commence à 0)
     * @param livres la liste des livres
     * si la liste de livre est vide, index=-1
     */
    fun updateContenuPanneauGauche(livres: MutableList<Livre>, index: Int=-1){
        this.panneauGauche.children.clear()

        var i = -1
        for (livre in livres) {
            i++
            val labelLivre = Label(livre.titre)
            this.panneauGauche.add(labelLivre, 0, i)
        }
        if(index >=0) {
            this.panneauGauche.children[index].style = "-fx-background-color:lightblue"
        }

    }

    /**
     * permet de mettre à jour la liste de livre dans le panneau de gauche, colore en bleu le livre dont l'index est index dans la liste
     * (commence à 0) et associe à chaque cellule de la grille un gestionnaire d'événement
     * @param livres la liste des livres
     * @param eventHandler le gestionnaire d'évènement
     *
     */
    fun updateContenuPanneauGauche(livres: MutableList<Livre>, eventHandler: EventHandler<MouseEvent>, index: Int) {
        this.updateContenuPanneauGauche(livres,index)
        for (livre in livres) {
            this.panneauGauche.children.forEach { it.onMouseClicked = eventHandler }
        }
    }


    /**
     * permet de mettre à jour le contenu du panneau de droite.
     * @param livre? le livre dont les informations sont à afficher ou null
     * @param int l'index du livre dans la bibliothèque
     */
    fun updateContenuPanneauDroit(numero: Int, livre: Livre?){
        this.panneauDroit.update(numero,livre)
    }

    /**
         * active / désactive le Bouton
         * @param etat vrai si le bouton doit être activé
         */
        fun activerBouton2PanneauDroit(etat: Boolean) {
            this.panneauDroit.bouton2.isDisable = !etat
        }
    /*
    donne le contenu du textField dans le panneau droit
    @return String le contenu
     */
    fun getContenuTextFieldPanneauDroit(): String{
    if (panneauDroit.textFieldTitre.text==null){
    return ""}
    else{
        return this.panneauDroit.textFieldTitre.text
        }
    }
    /**
   *donne l'item sélectionné dans l comboBox catégorie du panneau droit
   @return String l'item sélectionné
    */
    fun getCategorieComboBoxPanneauDroit(): String{
        return panneauDroit.comboBoxCategorie.selectionModel.selectedItem
    }
    /**
    *donne l'item sélectionné dans la ComboBox auteur du panneau droit
    @return Auteur l'item sélectionné
    */
    fun getAuteurComboBoxPanneauDroit(): Auteur {
        return panneauDroit.comboBoxAuteur.selectionModel.selectedItem
    }

        /**
         * active / désactive le Bouton permettant de passer au livre précédent
         * @param etat vrai si le bouton doit être activé
         */
        fun activerBouton1PanneauDroit(etat: Boolean) {
            this.panneauDroit.bouton1.isDisable = !etat
        }


        /**
         * ajoute un action listener au bouton bouton
         * @param bouton le bouton cible
         * @param action le listener à ajouter
         */
        fun fixeControleurBouton(bouton: Button, action: EventHandler<ActionEvent>) {
            bouton.onAction = action
        }


        /**
         * pour effacer le sélection d'un livre dans la grille de livre
         */
        fun effacerSelectionPanneauGauche() {
            this.panneauGauche.children.forEach {
                it.style = "-fx-background-color:white"
            }
        }

    /**
     * permet de colorier une ligne du GridPane en lightblue
     * @param numeroLigne le numéro de ligne à colorier
     */
    fun selectionnerLignePanneauGauche(numeroLigne:Int){
        val cellule=this.panneauGauche.children
        cellule[numeroLigne].style="-fx-background-color:lightblue"
    }

    /**
     * retourne le bouton1 du panneau droit
     * @return Button le bouton 1
     */
    fun getBouton1PanneauDroit(): Button {
        return this.panneauDroit.bouton1

    }
    /**
     * retourne le bouton2 du panneau droit
     * @return Button le bouton 2
     */
    fun getBouton2PanneauDroit(): Button {
        return this.panneauDroit.bouton2

    }

}


