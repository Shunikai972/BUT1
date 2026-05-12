package exo2

class Pays(no : String, cap : String, cont : String, pop : Int, sup : Double) : Comparable<Pays> {
    private var nom : String = no
    private var capitale : String = cap
    private var continent : String = cont
    private var population : Int = pop
    private var superficie : Double = sup
    fun donneNom() : String{
        return nom
    }
    fun donneCapitale(): String{
        return capitale
    }
    fun donneContinent(): String{
        return continent
    }
    fun donnePopulation(): Int{
        return population
    }
    fun donneSuperficie(): Double{
        return superficie
    }
    override fun toString(): String{
        return "$nom ($capitale) - $continent - Population: $population - Superficie: $superficie"
    }
    override fun equals(other: Any?): Boolean{
        if (other !is Pays) return false
        return this.nom == other.nom
    }
    override fun compareTo(other: Pays) : Int{
        return this.population.compareTo(other.population)
    }

}