package exo1

import gen.annotation.Todo
import info.but1.collections.NutArray
import info.but1.collections.nutArrayOfNulls

open class Personnage(nom : String, force: Int, agi : Int, intel : Int) {
    private var nom : String
    private var carac_force : Int
    private var carac_agilite: Int
    protected var carac_intelligence : Int
    private var inventaire = nutArrayOfNulls<Equipement>(6)
    var nbEquipementsPossedes: Int = 0
        private set
    init {
        this.nom = nom
        this.carac_force = force
        this.carac_agilite = agi
        this.carac_intelligence = intel
    }
    open fun listeCaracteristiques(): String{
        return "FOR=$carac_force,AGI=$carac_agilite,INT=$carac_intelligence"
    }
    fun prend(equipement: Equipement): Boolean{
        if(nbEquipementsPossedes <6){
            if(lEquipementContient(equipement)==false){
                inventaire[nbEquipementsPossedes] = equipement
                nbEquipementsPossedes++
                return true
            }
        }
        return false
    }
    fun attaque(arme: Arme, cible: Personnage): Boolean{
        if(lEquipementContient(arme) == true){
            println("attaque $cible avec $arme")
            return true
        }
        return false
    }
    override fun toString(): String {
        return "$nom <${listeCaracteristiques()}>"
    }
    fun lEquipementContient(equipement: Equipement): Boolean{
        for(i in 0 ..nbEquipementsPossedes){
            if (inventaire[i] == equipement) {
                return true
            }
        }
        return false
    }
    fun listeInventaire():String{
        var chaine : String = ""
        for(i in 0 .. nbEquipementsPossedes){
            if (inventaire[i] == null){
                continue
            }
            chaine += inventaire[i].toString()
        }
        return chaine
    }
}