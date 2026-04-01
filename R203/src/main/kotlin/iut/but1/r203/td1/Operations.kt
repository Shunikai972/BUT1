package iut.but1.r203.td1

/**
 * @author mottu-jm (inspired from rpouiller's java code)
 */
class Operations {

    private val COURT = 0
    private val LONG = 1
    private val TRESLONG = 2

    fun additionner(nombres: Array<Int>): Int {
        var retour = 0
        for (i in 0 until nombres.size) {
            retour += nombres[i]
        }
        this.afficherResultat(retour, "addition", COURT)
        return retour
    }

    fun soustraire(nombres: Array<Int>): Int {
        var retour = nombres[0]

        for (i in 1 until nombres.size) {
            retour -= nombres[i]
        }

        this.afficherResultat(retour, "soustraction", TRESLONG)
        return retour
    }

    fun multiplier(nombres: Array<Int>): Int {
        var retour = 1

        for (i in 0 until nombres.size) {
            retour *= nombres[i]
        }

        this.afficherResultat(retour, "multiplication", LONG)
        return retour
    }

    private fun afficherResultat(resultat: Int, operation: String, afficheur: Int) {
        if (afficheur == COURT) {
            println(resultat)
        } else if (afficheur == LONG) {
            println("Le résultat est $resultat")
        } else {
            println("Le résultat de la $operation est $resultat")
        }
    }
}