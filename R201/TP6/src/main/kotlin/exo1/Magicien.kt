package exo1

import h.ST_inside_t
import info.but1.collections.NutArray
import info.but1.collections.nutArrayOfNulls

class Magicien(nom: String, force: Int, agi:Int,intel : Int, mag : Int = 1, sort0: Sort): Personnage(nom,force,agi,intel) {
        private var carac_magie : Int = mag
        var nbSortsConnus: Int = 1
                private set
        private var grimoire = nutArrayOfNulls<Sort>(10)
        override fun listeCaracteristiques(): String{
                return super.listeCaracteristiques()+",MAG=$carac_magie"
        }
        fun leGrimoireContient(sort:Sort):Boolean{
                for(i in 0 .. nbSortsConnus){
                        if(grimoire[i] == sort){
                                return true
                        }
                }
                return false
        }
        fun apprend(nouveau:Sort): Boolean{
            TODO()
        }
        fun lance(sort:Sort):Boolean{
                TODO()
        }
        fun listeDesSorts(): NutArray<Sort?>{
                TODO()
        }
}