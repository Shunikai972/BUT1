package iut.collections

open class PileChainee : iPile  {
    protected var debut : Cellule? = null

    override fun estVide() : Boolean{
        if(debut == null){
            return true
        }
        return false
    }
    override fun taille() : Int{
        if(estVide()){
            return 0
        }
        var taye = 1
        while (debut?.suivant() != null){
            taye += 1
            debut = debut?.suivant()
        }
        return taye
    }
    override fun empiler(element: Int) {
        debut = Cellule(element, debut)
    }

    override fun consulter(): Int {
        if(debut != null){
            return debut!!.valeur()
        }
        throw NoSuchElementException("C'est pas bon frérot")
    }

    override fun depiler(): Int {
        if (estVide()) {
            throw NoSuchElementException("C pas bon")
        }
        val ancien : Int = debut!!.valeur()
        debut = debut?.suivant()
        return ancien
    }

}