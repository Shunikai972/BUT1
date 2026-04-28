package ihm.td1

import javafx.scene.canvas.GraphicsContext
import javafx.scene.layout.Pane
import javafx.scene.paint.Color


class Tableau : Pane() {

    init {
        this.style="-fx-background-color: white"
    }

    fun peindre(g: GraphicsContext) {
        tracerQuadrillage(g)


       //Décommentez les lignes suivantes au furet à mesure de vos développements


      /* var rectangle = Rectangle(Point(150.0, 150.0), 50.0, 200.0,Color.ORANGE);
        rectangle.dessiner(g)
        var point = Point(250.0, 200.0)
        point.dessiner(g)
        var cercle = Cercle(Point(300.0, 300.0), 30.0)
        cercle.dessiner(g)

        var groupe = Groupe(Color.BLUE)
        groupe.ajouter(Point(20.0, 30.0))
        groupe.ajouter(Rectangle (Point(100.0, 200.0), 200.0, 80.0 ))
        groupe.ajouter(Cercle(Point(250.0, 80.0), 50.0))
        groupe.dessiner(g)
*/
    }

    private fun tracerQuadrillage(g: GraphicsContext) {

        g.stroke=Color.LIGHTGRAY

        for( x in 1..1000 step 10) {
            g.strokeLine(x.toDouble(), 0.0, x.toDouble(), 1000.0)
        }

        for (x in 0..1000 step 10){
            g.strokeLine(0.0, x.toDouble(), 1000.0, x.toDouble())
        }
    }

}
