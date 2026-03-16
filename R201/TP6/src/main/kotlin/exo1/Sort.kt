package exo1

class Sort(nom:String,desc : String) {
    var nom : String = nom
        private set
    private var description : String = desc
    override fun toString(): String {
        return "$nom : $description"
    }
    override fun equals(other: Any?): Boolean {
        TODO()
    }
}