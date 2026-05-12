package exo2

import java.io.File
import java.io.FileNotFoundException
import java.util.*

class Monde(): iMonde {
   private val lesPays: MutableList<Pays> = mutableListOf()

    override fun taille(): Int {
        return lesPays.size
    }

    override fun ajouter(p: Pays): Boolean {
        if (p in lesPays){
            return false
        }
        return lesPays.add(p)
    }

    override fun toString(): String {
        var result = "Monde avec ${lesPays.size} pays:\n"
        for (pays in lesPays) {
            result += "  ${pays.donneNom()}\n"
        }
        return result
    }

    override fun remplir(nomFichier: String) {
        val file = File(nomFichier)
        if (!file.exists()) {
            throw FileNotFoundException("Fichier non trouvé: $nomFichier")
        }
        
        val lines = file.readLines()
        for (line in lines) {
            val trimmedLine = line.trim()
            if (trimmedLine.isEmpty()) continue
            
            val parts = trimmedLine.split(";")
            if (parts.size < 5) continue
            
            try {
                val pais = Pays(
                    parts[0].trim(),
                    parts[1].trim(),
                    parts[2].trim(),
                    parts[3].trim().toInt(),
                    parts[4].trim().toDouble()
                )
                lesPays.add(pais)
            } catch (e: NumberFormatException) {
                throw e
            }
        }
    }

    override fun getPays(nom: String): Pays? {
        for(i in 0 until lesPays.size){
            if(lesPays[i].donneNom() == nom){
                return lesPays[i]
            }
        }
        return null
    }

    override fun getPays2(nom: String): Pays? {
        val fauxPays = Pays(nom, "", "", 0, 0.0)
        val index = lesPays.indexOf(fauxPays)
        return if (index >= 0) lesPays[index] else null
    }

    override fun plusPeuple(ref: Pays): MutableList<Pays> {
        val result = mutableListOf<Pays>()
        for (pays in lesPays) {
            if (pays.compareTo(ref) >= 0) {
                result.add(pays)
            }
        }
        return result
    }

    override fun populationSuperieureA(limite: Int): MutableList<Pays> {
        val result = mutableListOf<Pays>()
        for (pays in lesPays) {
            if (pays.donnePopulation() > limite) {
                result.add(pays)
            }
        }
        return result
    }

    override fun monde(): MutableList<Pays> {
        return lesPays.toMutableList()
    }

    override fun mondeTrie(): MutableList<Pays> {
        val copie = lesPays.toMutableList()
        copie.sort()
        return copie
    }

    override fun mondeTrieTab(): Array<Pays> {
        val tab = lesPays.toTypedArray()
        tab.sort()
        return tab
    }

    override fun mondeTrieSelonSuperficie(): MutableList<Pays> {
        val copie = lesPays.toMutableList()
        copie.sortWith(ComparateurPays())
        return copie
    }
}
