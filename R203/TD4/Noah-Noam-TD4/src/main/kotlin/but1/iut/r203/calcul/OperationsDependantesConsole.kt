package but1.iut.r203.calcul

open class OperationsDependantesConsole : IntProvider {
    override fun getParam(): Int {
        var paramConsole = Integer.valueOf(readLine())
        return paramConsole
    }
}