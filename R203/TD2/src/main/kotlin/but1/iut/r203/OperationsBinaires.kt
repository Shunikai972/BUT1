package but1.iut.r203

/**
 * @author mottu-jm
 */
class OperationsBinaires {
    private val COURT = 0
    private val LONG = 1
    private val TRESLONG: Int = 2

    /**
     * Additionner deux entiers
     * @param op1, op2 : int Opérandes à additionner
     * @return somme : int
     * @throws ArithmeticException : out of Int bounds
     */
    fun additionner(op1: Int, op2: Int): Int {
        var retour = op1 + op2
        this.afficherResultat(retour, "addition", COURT)
        return retour
    }

    /**
     * Soustraire deux entiers
     * @param op1, op2 : int Opérandes à soustraire
     * @return resultat des soustractions
     * @throws ArithmeticException : out of Int bounds
     */
    fun soustraire(op1: Int, op2: Int): Int {
        var retour = op1 - op2
        this.afficherResultat(retour, "soustration", TRESLONG)
        return retour
    }

    /**
     * Multiplier  deux entiers
     * @param op1, op2 : int Opérandes à multiplier
     * @return produit : int
     * @throws ArithmeticException : out of Int bounds
     */
    fun multiplier(op1: Int, op2: Int): Int {
        var retour = op1 * op2
        this.afficherResultat(retour, "multiplication", LONG)
        return retour
    }

    /**
     * Diviser deux entiers naturels
     * @param dividende : entier naturel
     * @param diviseur : entier naturel
     * @return quotient : flottant
     * @throws ArithmeticException
     */
    fun diviserNaturel(dividende: Int, diviseur: Int): Float {
        if (diviseur == 0) {
            throw ArithmeticException("Trying to do a division by zero")
        }
        if (dividende < 0 || diviseur < 0) {
            throw ArithmeticException("Trying to do a division of non natural number(s)")
        }
        return (dividende / diviseur).toFloat()
    }

    /**
     * Calcul de la factorielle d’un entier n positif ou nul
     * @param int n un nombre dont on veut calculer la factorielle
     * @return le résultat n! = 1*2*...*n et 0! = 1
     * @throws IllegalArgumentException quand on essaie une factorielle d'un nombre négatif
     * @throws ArithmeticException : out of Int bounds
     */
    fun factorielle(n: Int): Int {
        if (n < 0) throw IllegalArgumentException("Parameter is too small")
        if (n > 12) throw ArithmeticException("Output would overreach the size of an int")
        return if (n == 0) {
            1
        } else {
            n * factorielle(n - 1)
        }
    }

    /**
     * Affiche le resultat de l'operation de differentes manieres (de COURT à TRESLONG pour le parametre de l'afficheur)
     * @param resultat : nombre à afficher
     * @param operation : opération effectuée
     * @param afficheur : constante déterminant la longueur de l'affichage
     */
    private fun afficherResultat(resultat: Int, operation: String, afficheur: Int) {
        if (afficheur == COURT) {
            println(resultat)
        } else if (afficheur == LONG) {
            println("Le résultat est $resultat")
        } else println("Le résultat de la $operation est $resultat")
    }


}




