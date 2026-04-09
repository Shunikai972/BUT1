import iut.collections.PileChainee

fun main(args: Array<String>) {
    val pile = PileChainee()
    pile.empiler(42)
    pile.empiler(0)
    pile.empiler(-2)
    pile.empiler(4)
    pile.empiler(-42)
    pile.empiler(-4)
    pile.empiler(99)

    print(pile.depiler())
}