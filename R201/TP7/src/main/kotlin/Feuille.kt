import info.but1.collections.NutArray
import info.but1.collections.nutArrayOfNulls

class Feuille(h : Int = 100, l : Int = 100){
    private var hauteur : Int = h
    private var largeur : Int = l
    var compteurObjets : Int = 0
        private set
    private var objets = nutArrayOfNulls<ObjetGraphique>(100)
    fun inserer(nouvelObjet:ObjetGraphique){
        TODO()
    }
}