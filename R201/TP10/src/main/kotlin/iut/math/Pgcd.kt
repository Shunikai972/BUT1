package iut.math

import kotlin.math.abs

/**
 * renvoie le "Plus Grand Commun Diviseur" de deux entiers
 * @param aa un entier
 * @param bb un entier
 * @return le pgcd de aa et bb
 */
fun pgcd(aa : Int, bb : Int) : Int {
    val a = abs(aa)
    val b = abs(bb)
    if (b == 0){
        return a
    }
    return pgcd(b, a%bb)
}