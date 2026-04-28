package but1.iut.r203

import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class OperationsBinairesTest {

    private val ops = OperationsBinaires()



    @Test
    fun additionner() {
        assertEquals(-8, ops.additionner(-5, -3))
        assertEquals(-5, ops.additionner(-5, 0))
        assertEquals(-2, ops.additionner(-5, 3))
        assertEquals(0, ops.additionner(0, 0))
        assertEquals(4, ops.additionner(0, 4))
        assertEquals(7, ops.additionner(4, 3))
        assertEquals(1, ops.additionner(4, -3))
        assertEquals(Int.MIN_VALUE, ops.additionner(Int.MAX_VALUE, 1))
        assertEquals(Int.MAX_VALUE*2, ops.additionner(Int.MAX_VALUE, Int.MAX_VALUE))
    }









    @Test
    fun soustraire() {
        assertEquals(-2, ops.soustraire(-5, -3))
        assertEquals(-5, ops.soustraire(-5, 0))
        assertEquals(-8, ops.soustraire(-5, 3))
        assertEquals(0, ops.soustraire(0, 0))
        assertEquals(2, ops.soustraire(5, 3))
        assertEquals(-2, ops.soustraire(3, 5))
        assertEquals(8, ops.soustraire(5, -3))
        assertEquals(Int.MAX_VALUE, ops.soustraire(Int.MIN_VALUE, 1))
    }
    @Test
    fun multiplier() {
        assertEquals(15, ops.multiplier(-5, -3))
        assertEquals(0, ops.multiplier(-5, 0))
        assertEquals(-15, ops.multiplier(-5, 3))
        assertEquals(0, ops.multiplier(0, 0))
        assertEquals(0, ops.multiplier(0, 4))
        assertEquals(12, ops.multiplier(4, 3))
        assertEquals(-12, ops.multiplier(4, -3))
        assertEquals(-2, ops.multiplier(Int.MAX_VALUE, 2))
    }
    @Test
    fun diviserNaturel() {
        assertEquals(0.0f, ops.diviserNaturel(0, 3))
        assertEquals(3.0f, ops.diviserNaturel(6, 2))
        assertEquals(3.0f, ops.diviserNaturel(9, 3))
        assertEquals(3.0f, ops.diviserNaturel(7, 2))
        assertThrows(ArithmeticException::class.java) {
            ops.diviserNaturel(5, 0)
        }
        assertThrows(ArithmeticException::class.java) {
            ops.diviserNaturel(-1, 2)
        }
        assertThrows(ArithmeticException::class.java) {
            ops.diviserNaturel(5, -2)
        }
        assertThrows(ArithmeticException::class.java) {
            ops.diviserNaturel(-4, 0)
        }
    }
    @Test
    fun factorielle() {
        assertEquals(1, ops.factorielle(0))
        assertEquals(1, ops.factorielle(1))
        assertEquals(120, ops.factorielle(5))
        assertEquals(3628800, ops.factorielle(10))
        assertEquals(5040, ops.factorielle(7))
        assertThrows(IllegalArgumentException::class.java) {
            ops.factorielle(-1)
        }
        assertThrows(IllegalArgumentException::class.java) {
            ops.factorielle(-100)
        }

        assertThrows(ArithmeticException::class.java) {
            ops.factorielle(13)
        }
        assertEquals(479001600, ops.factorielle(12))
    }
}