abstract  class ObjetDessinable(x: Int, y : Int): IDessinable, ObjetGraphique(){
    private var estSelectionne  : Boolean = false
    private var couleurPremierPlan : String = "#000000"
    private var couleurArrierePlan: String = "FFFFFF"
    private var position : Point = Point(x,y)
    override fun selectionner(ok: Boolean){
        TODO()
    }
    override fun colorierPremier(nouvelleCouleur: String){
        TODO()
    }
    override fun colorierArriere(nouvelleCouleur: String){
        TODO()
    }
    override fun deplacer(deplacement : Vecteur){
        position = deplacement.appliquer(position)
    }
    override fun dessiner(): String{
        return "X=${position.x},Y=${position.y}"
    }
}