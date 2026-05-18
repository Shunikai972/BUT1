package but1.iut.r203

/**
 * Cette classe stocke des notes et calcule des statistiques
 * @param listesDesNotes : les notes entrées en paramètre du constructeur
 * @property listesDesNotes : attributs de classe stockant les notes de l'instance de classe
 * @author mottu-jm
 */
class ListeNotes(var listeDesNotes: IntArray) {

    /**
      * @return renvoie la moyenne des notes
     */
    fun moyenne(): Float {
        var somme = 0
        var index = 0
        while (index < listeDesNotes.size) {
            somme += listeDesNotes[index]
            index++
        }
        return if (listeDesNotes.isNotEmpty()) {
            somme.toFloat() / listeDesNotes.size
        } else {
            0.0F
        }
    }

    /**
     * Compte les occurences d'une note dans la liste des notes
     * @param elem
     * @return renvoie le nombre de notes trouvées
     * @throws IllegalArgumentException
     */
    fun nombreOccurence(elem: Int): Int {
        var compteurOccurence = 0
        if (elem > 20 || elem < 0) throw IllegalArgumentException()
        var i = 0
        while (i < listeDesNotes.size) {
            if (elem == listeDesNotes[i]) {
                compteurOccurence++
            }
            i++
        }
        return compteurOccurence
    }

    /**
     * Calcule la note maximale et le nombre d'élèves qui l'ont obtenue.
     * @throws IllegalArgumentException si la liste est vide
     */
    fun noteMaximaleEtNombreEleves(): StatistiqueNote {
        require(listeDesNotes.isNotEmpty()) { "La liste de notes ne peut pas être vide." }

        var noteMaximale = listeDesNotes[0]
        var nombreEleves = 1
        var index = 1

        while (index < listeDesNotes.size) {
            val note = listeDesNotes[index]
            if (note > noteMaximale) {
                noteMaximale = note
                nombreEleves = 1
            } else if (note == noteMaximale) {
                nombreEleves++
            }
            index++
        }

        return StatistiqueNote(noteMaximale, nombreEleves)
    }

    /**
     * Calcule la note minimale et le nombre d'élèves qui l'ont obtenue.
     * @throws IllegalArgumentException si la liste est vide
     */
    fun noteMinimaleEtNombreEleves(): StatistiqueNote {
        require(listeDesNotes.isNotEmpty()) { "La liste de notes ne peut pas être vide." }

        var noteMinimale = listeDesNotes[0]
        var nombreEleves = 1
        var index = 1

        while (index < listeDesNotes.size) {
            val note = listeDesNotes[index]
            if (note < noteMinimale) {
                noteMinimale = note
                nombreEleves = 1
            } else if (note == noteMinimale) {
                nombreEleves++
            }
            index++
        }

        return StatistiqueNote(noteMinimale, nombreEleves)
    }

    fun afficherListe() {
        for (i in listeDesNotes.indices) {
            print(listeDesNotes[i])
            print(" ; ")
        }
        println("Il y a " + listeDesNotes.size + " elements dans la liste.")
    }
}

/**
 * Associe une note à son nombre d'occurrences dans la liste.
 */
data class StatistiqueNote(val note: Int, val nombreEleves: Int)
