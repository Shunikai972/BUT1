import exo2.*
import exo3.*

fun main() {
    println("=== Test Exo 2: Monde ===")
    
    // Test Pays
    val france = Pays("France", "Paris", "Europe", 67000000, 643801.0)
    println("France: $france")
    
    val allemagne = Pays("Allemagne", "Berlin", "Europe", 83370000, 357022.0)
    println("Allemagne: $allemagne")
    
    // Test Monde
    val monde = Monde()
    println("Taille initiale: ${monde.taille()}")
    
    val ajoutFrance = monde.ajouter(france)
    println("France ajoutée: $ajoutFrance")
    
    val ajoutAllemagne = monde.ajouter(allemagne)
    println("Allemagne ajoutée: $ajoutAllemagne")
    
    val ajoutFranceAgain = monde.ajouter(france)
    println("France ajoutée à nouveau: $ajoutFranceAgain")
    
    println("Taille après ajout: ${monde.taille()}")
    
    // Test getPays
    val franceTrouvee = monde.getPays("France")
    println("getPays('France'): $franceTrouvee")
    
    // Test getPays2
    val franceTrouvee2 = monde.getPays2("France")
    println("getPays2('France'): $franceTrouvee2")
    
    // Test comparaison
    val plusPeuple = monde.plusPeuple(allemagne)
    println("Pays plus peuplés que l'Allemagne (${allemagne.donnePopulation()}): ")
    for (p in plusPeuple) {
        println("  - ${p.donneNom()}: ${p.donnePopulation()}")
    }
    
    // Test mondeTrie
    val mondeTrie = monde.mondeTrie()
    println("Monde trié par population:")
    for (p in mondeTrie) {
        println("  - ${p.donneNom()}: ${p.donnePopulation()}")
    }
    
    // Test mondeTrieSelonSuperficie
    val mondeTrieSuperficie = monde.mondeTrieSelonSuperficie()
    println("Monde trié par superficie:")
    for (p in mondeTrieSuperficie) {
        println("  - ${p.donneNom()}: ${p.donneSuperficie()}")
    }
    
    println("\n=== Test Exo 3: Tableau aléatoire ===")
    val tableau = tableauAleatoireDistinct(10)
    print("Tableau aléatoire distinct: [")
    for (v in tableau)
        print("$v ")
    println("]")
}