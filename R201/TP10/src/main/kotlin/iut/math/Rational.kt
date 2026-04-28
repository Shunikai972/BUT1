package iut.math
import kotlin.math.pow

/**
 * Fabrique un rationnel à partir de deux entiers passés en paramètre ;
 * @param numerator le numérateur
 * @param denominator le dénominateur
 * @throws IllegalArgumentException si le denominateur vaut 0
 */
class Rational(numerator : Int, denominator : Int = 1)  {
    private var numerator : Int
    private var denominator : Int

    init {
        this.numerator = numerator

        if(denominator == 0){
            throw IllegalArgumentException()
        }
        this.denominator = denominator
    }
    //TODO : implementer la classe Rational comme décrite dans le README

    /**
     * renvoie une chaine de caractères représentant le rationnel
     * @return cette chaine
     */
    override fun toString(): String {
        return ("$numerator/$denominator")
    }

    /**
     * realise la somme d'un rationnel avec le rationnel courant
     * @param rational le rationnel à ajouter
     * @return la somme
     */
    operator fun plus(rational : Rational) : Rational {
        return Rational((numerator*rational.denominator + denominator*rational.numerator),(denominator*rational.denominator))
    }

    operator fun plus(n : Int) : Rational {
        return Rational((numerator+ n*denominator), denominator)
    }

    fun opposite() : Rational{
        return Rational(-numerator, denominator)
    }

    operator fun minus(rational: Rational) : Rational{
        return plus(rational.opposite())
    }

    fun reduce() : Rational{
        val p = pgcd(numerator, denominator)
        var j = numerator/p
        var i = denominator/p
        if(i<0){
            j = -j
            i = -i
        }
        return Rational(j,i)
    }
    


    // TODO : continuer en suivant le README


}

/**
 * A IMPLEMENTER QUESTION 8
 *
 * Fonction donnant une approximation rationnelle d'un flottant à une approximation près
 * @param x le flottant à convertir
 * @param n l'approximation à considérer
 * @return un Rationnel approximant x
 */
fun toRational(x : Double, n : Int) : Rational {
    TODO("retourne un rationnel approximant la valeur réelle")
}
