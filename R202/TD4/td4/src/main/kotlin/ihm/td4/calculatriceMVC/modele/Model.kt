package ihm.td4.calculatriceMVC.modele
class Model(){
    fun calculer(operande1: Int, operateur: String, operande2: Int): Int {
        return when (operateur) {
            "+" -> operande1 + operande2
            "-" -> operande1 - operande2
            "*" -> operande1 * operande2
            "/" -> operande1 / operande2
            else -> throw IllegalArgumentException("Opérateur inconnu")
        }
    }
}
