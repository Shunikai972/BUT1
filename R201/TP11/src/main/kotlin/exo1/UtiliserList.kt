package exo1

fun main() {
    var jsp :MutableList<String> = mutableListOf("Bonjour","Aurevoir","Salut","YOUHOUUUUUUUUU", "Explosif","Méchant","Gentil","Méchant","Gentil","gentil","méchant","UWU","OWO")// TODO
    var wantedlement : String = "Jsp"
    fun afficherlist(liste : MutableList<String>){
        for( i in 0 until  jsp.size){
            println(jsp[i])
        }
    }
    jsp.add(0,"WWWWWWWW")
    jsp.add(jsp.size,"qzdqdqsdqf")
    jsp.add(5,"stadsqd")

    afficherlist(jsp)
    if("s" in jsp){
        println("yes")
    } else{
        println("no")
    }
    for( i in 0 until  jsp.size){
        if(jsp[i] == wantedlement){
            println("position de l'élément voulue ="+i)
        }else{
            println("pas trouvé mdr")
        }
    }
    fun Trialpha (liste : MutableList<String>){
        var jsp2 = liste.sorted()
        for( i in 0 until  liste.size){
            println(jsp2[i])
        }
    }
    Trialpha(jsp)
}

