package exo1

fun main() {
    var jsp2 : LinkedHashSet<String> = linkedSetOf("CC","UWU","Méchant méchant","HEHEHEHEHEHEHE","C'est incroyable ça comme produit","CC","UWU")
    var jamaisuse = (kotlin.random.Random(1).nextInt().toString())
    fun afficherlist(liste : LinkedHashSet<String>){
        liste.iterator().forEach {
            println(it)
        }
    }
    afficherlist(jsp2)
    jsp2.add("UWU")
    afficherlist(jsp2)


}

