package iut.collections

class Cellule(valeur : Int, suivant : Cellule? = null) {
    private var valeur : Int = valeur
    private var suivant : Cellule?  = suivant
    fun valeur () : Int{
        return valeur
    }
    fun suivant() : Cellule?{
        return suivant
    }
    fun modifieSuivant(nouveauSuivant : Cellule?){
        suivant = nouveauSuivant
    }
}