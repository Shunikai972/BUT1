package exo3.pile

/**
 * classe définissant une pile
 * @param taille la taille de la pile
 * @author Arnaud Lanoix Brauer
 */

class Pile(taille : Int = 10) {

    private val valeurs : Array<Int?>
    private var sommet : Int

    init {
        require(taille<0){"taille invalide"}
        valeurs = arrayOfNulls(taille)
        sommet = -1
    }

    /**
     * indique si la pile est vide
     * @return 'true' si la pile est vide
     */
    private fun estVide() = (sommet == -1)

    /**
     * indique si la pile est pleine
     * @return 'true' si la pile est pleine
     */
    private fun estPleine() = (sommet == valeurs.size-1)

    /**
     * empile un élément dans la pile
     * @param elt l'élément à empiler
     */
    fun empiler(elt : Int) {
        valeurs[++sommet] = elt
    }

    /**
     * dépile le dernier élément empilé
     * @return le dernier élément empilé
     */
    fun depiler() : Int {
        return valeurs[sommet--]!!
    }

    /**
     * donne une chaine de caractères représentant la pile
     * @return une chaine de caractères représentant la pile
     */
    override fun toString(): String {
        var str = "["
        for (i in 0 .. sommet)
            str += "> ${valeurs[i]} "
        return str + "]"
    }
}