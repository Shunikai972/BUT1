import info.but1.collections.NutArray
import info.but1.collections.nutArrayOfNulls

class Groupe(objet1 : ObjetGraphique, objet2 : ObjetGraphique) : ObjetGraphique(){
     var nbObjetsRegroupes : Int = 2
        private set
    private var objetsRegroupes = nutArrayOfNulls<ObjetGraphique>(15)

    init {
        objetsRegroupes[0] = objet1
        objetsRegroupes[1] = objet2
    }
    fun regrouper(objetAjoute: ObjetGraphique) : Boolean{
        if(nbObjetsRegroupes >= objetsRegroupes.size){
            return false
        }else{
            objetsRegroupes[nbObjetsRegroupes] = objetAjoute
            nbObjetsRegroupes
            nbObjetsRegroupes++
            return true
        }
    }
    override fun deplacer(deplacement: Vecteur) {
        TODO("Not yet implemented")
    }

    override fun selectionner(ok: Boolean) {
        TODO("Not yet implemented")
    }
}