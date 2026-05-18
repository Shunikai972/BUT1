package but1.iut.r203

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Assertions.assertThrows
import org.junit.jupiter.api.Test

class OperationsUnairesTest {
    private val operationsUnaires = OperationsUnaires()

    @Test
    fun factorielleCT1EntierNegatif() {
        assertThrows(IllegalArgumentException::class.java) {
            operationsUnaires.factorielle(-1)
        }
    }

    @Test
    fun factorielleCT2DepassementInt() {
        assertThrows(ArithmeticException::class.java) {
            operationsUnaires.factorielle(13)
        }
    }

    @Test
    fun factorielleCT3Zero() {
        assertEquals(1, operationsUnaires.factorielle(0))
    }

    @Test
    fun factorielleCT4Trois() {
        assertEquals(6, operationsUnaires.factorielle(3))
    }

    @Test
    fun factorielleCasSupplementaireUn() {
        assertEquals(1, operationsUnaires.factorielle(1))
    }

    @Test
    fun factorielleCasSupplementaireBorneSuperieure() {
        assertEquals(479001600, operationsUnaires.factorielle(12))
    }
}
