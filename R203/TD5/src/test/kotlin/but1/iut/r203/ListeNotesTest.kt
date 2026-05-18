package but1.iut.r203

import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class ListeNotesTest {
    @Test
    fun moyenneCT1ListeVide() {
        assertEquals(0.0F, ListeNotes(intArrayOf()).moyenne(), 0.0001F)
    }

    @Test
    fun moyenneCT2UneNoteNulle() {
        assertEquals(0.0F, ListeNotes(intArrayOf(0)).moyenne(), 0.0001F)
    }

    @Test
    fun moyenneCT3PlusieursNotesNulles() {
        assertEquals(0.0F, ListeNotes(intArrayOf(0, 0, 0)).moyenne(), 0.0001F)
    }

    @Test
    fun moyenneCT4UneNoteNonNulle() {
        assertEquals(5.0F, ListeNotes(intArrayOf(5)).moyenne(), 0.0001F)
    }

    @Test
    fun moyenneCT5DeuxNotesEgales() {
        assertEquals(5.0F, ListeNotes(intArrayOf(5, 5)).moyenne(), 0.0001F)
    }

    @Test
    fun moyenneCT6DeuxNotesDifferentes() {
        assertEquals(7.5F, ListeNotes(intArrayOf(5, 10)).moyenne(), 0.0001F)
    }

    @Test
    fun nombreOccurenceCT1NoteNegative() {
        assertThrows(IllegalArgumentException::class.java) {
            ListeNotes(intArrayOf()).nombreOccurence(-5)
        }
    }

    @Test
    fun nombreOccurenceCT2NoteSuperieureA20() {
        assertThrows(IllegalArgumentException::class.java) {
            ListeNotes(intArrayOf()).nombreOccurence(50)
        }
    }

    @Test
    fun nombreOccurenceCT3ListeVide() {
        assertEquals(0, ListeNotes(intArrayOf()).nombreOccurence(5))
    }

    @Test
    fun nombreOccurenceCT4UneOccurrence() {
        assertEquals(1, ListeNotes(intArrayOf(5)).nombreOccurence(5))
    }

    @Test
    fun nombreOccurenceCT5DeuxOccurrencesConsecutives() {
        assertEquals(2, ListeNotes(intArrayOf(5, 5)).nombreOccurence(5))
    }

    @Test
    fun nombreOccurenceCT6DeuxOccurrencesSeparees() {
        assertEquals(2, ListeNotes(intArrayOf(5, 0, 5)).nombreOccurence(5))
    }

    @Test
    fun nombreOccurenceCT7DeuxOccurrencesDansListePlusLongue() {
        assertEquals(2, ListeNotes(intArrayOf(0, 5, 0, 5)).nombreOccurence(5))
    }

    @Test
    fun nombreOccurenceCT8AucuneOccurrence() {
        assertEquals(0, ListeNotes(intArrayOf(0, 5, 0, 5)).nombreOccurence(6))
    }

    @Test
    fun noteMaximaleListeVide() {
        assertThrows(IllegalArgumentException::class.java) {
            ListeNotes(intArrayOf()).noteMaximaleEtNombreEleves()
        }
    }

    @Test
    fun noteMaximaleUneSeuleNote() {
        assertEquals(
            StatistiqueNote(12, 1),
            ListeNotes(intArrayOf(12)).noteMaximaleEtNombreEleves()
        )
    }

    @Test
    fun noteMaximaleListeNonTrieeAvecDoublons() {
        assertEquals(
            StatistiqueNote(20, 2),
            ListeNotes(intArrayOf(7, 20, 13, 20, 4)).noteMaximaleEtNombreEleves()
        )
    }

    @Test
    fun noteMaximaleValeurFrontiere() {
        assertEquals(
            StatistiqueNote(20, 2),
            ListeNotes(intArrayOf(0, 20, 0, 20)).noteMaximaleEtNombreEleves()
        )
    }

    @Test
    fun noteMinimaleListeVide() {
        assertThrows(IllegalArgumentException::class.java) {
            ListeNotes(intArrayOf()).noteMinimaleEtNombreEleves()
        }
    }

    @Test
    fun noteMinimaleUneSeuleNote() {
        assertEquals(
            StatistiqueNote(12, 1),
            ListeNotes(intArrayOf(12)).noteMinimaleEtNombreEleves()
        )
    }

    @Test
    fun noteMinimaleListeNonTrieeAvecDoublons() {
        assertEquals(
            StatistiqueNote(4, 2),
            ListeNotes(intArrayOf(7, 4, 13, 4, 20)).noteMinimaleEtNombreEleves()
        )
    }

    @Test
    fun noteMinimaleValeurFrontiere() {
        assertEquals(
            StatistiqueNote(0, 2),
            ListeNotes(intArrayOf(0, 20, 0, 20)).noteMinimaleEtNombreEleves()
        )
    }
}
