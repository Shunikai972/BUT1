package iut.but1.r203.td1

import org.junit.jupiter.api.Test
import org.junit.jupiter.api.Assertions.*

class OperationsTest {

    // ---------- ADDITIONNER ----------

    @Test
    fun testAdditionner() {
        val op = Operations()
        val c = op.additionner(arrayOf(1, 2))
        assertEquals(3, c)
    }

    @Test
    fun testAdditionner1() {
        val op = Operations()
        val c = op.additionner(arrayOf(5, 5))
        assertEquals(10, c)
    }

    @Test
    fun testAdditionner2() {
        val op = Operations()
        val c = op.additionner(arrayOf(5, 5, 5))
        assertEquals(15, c)
    }

    @Test
    fun testAdditionner3() {
        val op = Operations()
        val c = op.additionner(arrayOf(1, 2, 3))
        assertEquals(6, c)
    }

    @Test
    fun testAdditionner4() {
        val op = Operations()
        val c = op.additionner(arrayOf(2, 1, 3))
        assertEquals(6, c)
    }

    @Test
    fun testAdditionner5() {
        val op = Operations()
        val c = op.additionner(arrayOf(3, 2, 1))
        assertEquals(6, c)
    }

    @Test
    fun testAdditionner6() {
        val op = Operations()
        val c = op.additionner(arrayOf(0))
        assertEquals(0, c)
    }

    @Test
    fun testAdditionner7() {
        val op = Operations()
        val c = op.additionner(arrayOf(-4, 4))
        assertEquals(0, c)
    }

    @Test
    fun testAdditionner8() {
        val op = Operations()
        val c = op.additionner(arrayOf(-2, 0, 2))
        assertEquals(0, c)
    }


    // ---------- SOUSTRAIRE ----------

    @Test
    fun testSoustraire1() {
        val op = Operations()
        val c = op.soustraire(arrayOf(3, 3))
        assertEquals(0, c)
    }

    @Test
    fun testSoustraire2() {
        val op = Operations()
        val c = op.soustraire(arrayOf(3, 3, 3))
        assertEquals(-3, c)
    }

    @Test
    fun testSoustraire3() {
        val op = Operations()
        val c = op.soustraire(arrayOf(1, 2, 3))
        assertEquals(-4, c)
    }

    @Test
    fun testSoustraire4() {
        val op = Operations()
        val c = op.soustraire(arrayOf(2, 1, 3))
        assertEquals(-2, c)
    }

    @Test
    fun testSoustraire5() {
        val op = Operations()
        val c = op.soustraire(arrayOf(3, 2, 1))
        assertEquals(0, c)
    }

    @Test
    fun testSoustraire6() {
        val op = Operations()
        val c = op.soustraire(arrayOf(0))
        assertEquals(0, c)
    }

    @Test
    fun testSoustraire7() {
        val op = Operations()
        val c = op.soustraire(arrayOf(-5, 5))
        assertEquals(-10, c)
    }

    @Test
    fun testSoustraire8() {
        val op = Operations()
        val c = op.soustraire(arrayOf(-2, 0, 2))
        assertEquals(-4, c)
    }


    // ---------- MULTIPLIER ----------

    @Test
    fun testMultiplier1() {
        val op = Operations()
        val c = op.multiplier(arrayOf(3, 3))
        assertEquals(9, c)
    }

    @Test
    fun testMultiplier2() {
        val op = Operations()
        val c = op.multiplier(arrayOf(3, 3, 3))
        assertEquals(27, c)
    }

    @Test
    fun testMultiplier3() {
        val op = Operations()
        val c = op.multiplier(arrayOf(1, 2, 3))
        assertEquals(6, c)
    }

    @Test
    fun testMultiplier4() {
        val op = Operations()
        val c = op.multiplier(arrayOf(2, 1, 3))
        assertEquals(6, c)
    }

    @Test
    fun testMultiplier5() {
        val op = Operations()
        val c = op.multiplier(arrayOf(3, 2, 1))
        assertEquals(6, c)
    }

    @Test
    fun testMultiplier6() {
        val op = Operations()
        val c = op.multiplier(arrayOf(0))
        assertEquals(0, c)
    }

    @Test
    fun testMultiplier7() {
        val op = Operations()
        val c = op.multiplier(arrayOf(-5, 5))
        assertEquals(-25, c)
    }

    @Test
    fun testMultiplier8() {
        val op = Operations()
        val c = op.multiplier(arrayOf(-2, 0, 2))
        assertEquals(0, c)
    }

}