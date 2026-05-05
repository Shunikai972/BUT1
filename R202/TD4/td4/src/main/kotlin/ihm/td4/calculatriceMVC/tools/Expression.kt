package ihm.td4.calculatriceMVC.tools

class Expression(expression: String) {
    private var expression: String

    init{
        this.expression=expression
    }

    // détermine si une expression est valide
    // 456, 65+, 78-35, 546/89= sont valides par exemple
    fun isValidExpression(): Boolean {
        val hasEquals = expression.endsWith("=")
        val expr = if (hasEquals) expression.dropLast(1) else expression

        // Regex : nombre + optionnel (opérateur + nombre OU opérateur seul)
        val pattern = Regex("^\\d+(\\s*[+\\-*/](\\s*\\d+)?)?$")

        if (!pattern.matches(expr)) return false

        if (hasEquals) {
            // expression mathématique complète (pas juste "456+" mais 678*89=)
            return Regex("\\d+\\s*[+\\-*/]\\s*\\d+").matches(expr)
        }

        return true
    }
}



fun main() {


    val tests = listOf("7","2567", "567+","=","346++", "23458+345", "32345+7896=", "++123", "123=", "+123","*", "123+=")

    tests.forEach {
        println("$it -> ${Expression(it).isValidExpression()}")
    }
}
