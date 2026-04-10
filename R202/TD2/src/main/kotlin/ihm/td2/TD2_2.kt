package ihm.td2

import javafx.application.Application
import javafx.geometry.Pos
import javafx.scene.Scene
import javafx.scene.control.*
import javafx.scene.layout.BorderPane
import javafx.scene.layout.GridPane
import javafx.scene.layout.HBox
import javafx.scene.layout.VBox
import javafx.stage.Stage


const val TEXTE= ("voici un texte relativement long à "
        + "lire et qui n'a aucune sorte d'intérêt à part"
        + " celui de prendre beaucoup de place et donc "
        + "d'occuper de l'espace dans la TextArea... ")
const val ADMINISTRATEUR = "Administrateur"
const val ETUDIANT = "Etudiant(e)"
const val ENSEIGNANT = "Enseignant(e)"
const val ETAT = "Etat de l'application > en cours d'identification"


class TD2_2: Application() {

    override fun start(primaryStage: Stage) {
        val root = BorderPane()
        
        // Left side - TextArea
        val textArea = TextArea(TEXTE)
        textArea.isWrapText = true
        textArea.isEditable = false
        textArea.prefWidth = 300.0
        
        // Right side - Form
        val formGrid = GridPane()
        formGrid.style = "-fx-border-color: lightgrey; -fx-padding: 10; -fx-gap: 10;"
        formGrid.vgap = 10.0
        formGrid.hgap = 10.0
        
        // Bienvenue title
        val welcomeLabel = Label("Bienvenue")
        welcomeLabel.style = "-fx-font: bold 16;"
        formGrid.add(welcomeLabel, 0, 0)
        GridPane.setColumnSpan(welcomeLabel, 2)
        
        // Login
        val loginLabel = Label("Login:")
        val loginField = TextField()
        formGrid.add(loginLabel, 0, 1)
        formGrid.add(loginField, 1, 1)
        
        // Password
        val passwordLabel = Label("Password:")
        val passwordField = PasswordField()
        formGrid.add(passwordLabel, 0, 2)
        formGrid.add(passwordField, 1, 2)
        
        // Role selection
        val roleCombo = ComboBox<String>()
        roleCombo.items.setAll(ADMINISTRATEUR, ETUDIANT, ENSEIGNANT)
        roleCombo.value = ETUDIANT
        formGrid.add(roleCombo, 0, 3)
        GridPane.setColumnSpan(roleCombo, 2)
        
        // Connection button - positioned using GridPane static methods
        val connectBtn = Button("Connexion")
        formGrid.add(connectBtn, 1, 3)
        GridPane.setHalignment(connectBtn, javafx.geometry.HPos.RIGHT)
        
        // Formation selection
        val formationLabel = Label("Choix formation")
        formationLabel.style = "-fx-font: bold;"
        formGrid.add(formationLabel, 0, 4)
        GridPane.setColumnSpan(formationLabel, 2)
        
        val info1 = RadioButton("info1")
        val info2 = RadioButton("info2")
        val info3 = RadioButton("info3")
        val toggleGroup = ToggleGroup()
        info1.toggleGroup = toggleGroup
        info2.toggleGroup = toggleGroup
        info3.toggleGroup = toggleGroup
        
        formGrid.add(info1, 0, 5)
        formGrid.add(info2, 0, 6)
        formGrid.add(info3, 0, 7)
        
        // Course selection
        val courseLabel = Label("Choix Parcours")
        courseLabel.style = "-fx-font: bold;"
        formGrid.add(courseLabel, 0, 8)
        GridPane.setColumnSpan(courseLabel, 2)
        
        val course1 = CheckBox("Parcours 1")
        val course2 = CheckBox("Parcours 2")
        val course3 = CheckBox("Parcours 3")
        
        formGrid.add(course1, 0, 9)
        formGrid.add(course2, 0, 10)
        formGrid.add(course3, 0, 11)
        
        // Main horizontal layout
        val centerBox = HBox(10.0)
        centerBox.children.addAll(textArea, formGrid)
        
        // Bottom status bar
        val statusBar = ProgressBar(0.3)
        statusBar.prefWidth = Double.MAX_VALUE
        val statusLabel = Label(ETAT)
        val bottomBox = HBox(10.0)
        bottomBox.children.addAll(statusBar, statusLabel)
        bottomBox.style = "-fx-padding: 5;"
        
        root.center = centerBox
        root.bottom = bottomBox
        
        val scene = Scene(root, 800.0, 500.0)
        primaryStage.title = "TD2_2 en javaFX"
        primaryStage.scene = scene
        primaryStage.show()
    }
}


fun main() {
    Application.launch(TD2_2::class.java)
}





