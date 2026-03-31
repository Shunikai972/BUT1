import info.but1.collections.NutArray
import info.but1.collections.nutArrayOfNulls

class Document(nom : String): IEditable{
    var compteurFeuilles : Int = 1
        private set
    private var feuilles = nutArrayOfNulls<ObjetGraphique>(10)
    override fun editer(nouveau: String) {
        TODO("Not yet implemented")
    }
    fun nouvelleFeuille(hauteur : Int = 200, largeur : Int = 150) : Boolean{
        TODO()
    }
    fun donneFeuille(position: Int): Feuille?{
        TODO()
    }
}