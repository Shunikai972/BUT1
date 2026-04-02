package iut.exo1

import java.net.InetAddress

fun main() {
    print("Saisir un nom : ")
    val nom = readln()
    val ip = nomVersIP(nom)
    println("IP correspondante :$ip")
}


fun nomVersIP(nom : String) : String {
    try{
        return InetAddress.getByName(nom).hostAddress
    }
    catch( e : Exception ){
        print("erreur adresse")
        return "1"

    }
    return "-1"

}