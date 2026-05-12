package exo3

import kotlin.random.Random

fun main() {
    val tableau = tableauAleatoireDistinct(10)
    print("[")
    for (v in tableau)
        print("$v-")
    println("]")
}

fun tableauAleatoireDistinct(n: Int): Array<Int> {
    val resultat = Array<Int>(n, { i -> 0 })
    // tableau de taille n complètement initialisé à 0
    
    val ensemble = HashSet<Int>()
    var index = 0
    
    while (ensemble.size < n) {
        val nombreAleatoire = Random.nextInt(1, n + 1)
        if (ensemble.add(nombreAleatoire)) {
            resultat[index] = nombreAleatoire
            index++
        }
    }
    
    return resultat
}