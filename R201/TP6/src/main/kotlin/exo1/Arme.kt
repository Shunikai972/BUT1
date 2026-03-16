package exo1

class Arme(desc:String, att:Int,def:Int): Equipement(desc) {
    private var defense : Int
    private var attaque : Int
    init {
        this.defense = def
        this.attaque = att
    }
    override fun toString():String {
        return "$description ($attaque,$defense)"
    }
}