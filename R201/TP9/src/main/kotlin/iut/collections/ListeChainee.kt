package iut.collections

class ListeChainee : PileChainee(), iFile, iListe {
    override fun index(element: Int): Int{
        var next = debut
        var taille = 0
        while(next!!.suivant() != null)    {
            if(next.valeur() == element){
                return taille
            }
            taille++
            next=next.suivant()
        }
        throw NoSuchElementException("fsdgf")
    }

    override fun consulter(index: Int): Int {
        var next = debut
        if(next!!.suivant() == null){
            throw IndexOutOfBoundsException("outofbound")
        }
        while()
    }

    override fun inserer(index: Int, element: Int) {
        TODO("Not yet implemented")
    }

    override fun supprimer(index: Int) {
        TODO("Not yet implemented")
    }

    override fun insererEnQueue(element: Int) {
        TODO("Not yet implemented")
    }

    override fun consulterEnQueue(): Int {
        TODO("Not yet implemented")
    }

    override fun supprimerEnQueue() {
        TODO("Not yet implemented")
    }
}