package exo2

import java.util.Comparator

class ComparateurPays : Comparator<Pays> {
    override fun compare(o1: Pays, o2: Pays): Int {
        return o1.donneSuperficie().compareTo(o2.donneSuperficie())
    }
}