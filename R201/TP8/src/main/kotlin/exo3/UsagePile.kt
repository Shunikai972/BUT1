package exo3

import exo3.pile.Pile

fun main() {
    val pile0 = Pile(-2)
    val pile =  Pile(4)
    println(pile)
    pile.empiler(42)
    pile.empiler(-4)
    pile.empiler(2)
    println(pile)
    println(pile.depiler())
    println(pile.depiler())
    println(pile)
    pile.empiler(2)
    pile.empiler(-42)
    pile.empiler(4)
    println(pile)
    println(pile.depiler())
    println(pile.depiler())
    println(pile.depiler())
    println(pile)
}