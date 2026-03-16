package exo1

open class Equipement(desc: String) {
    var description : String = desc
        private set
    override fun toString():String {
        return description
    }
}