package but1.iut.r203.calcul

import io.mockk.called
import io.mockk.every
import io.mockk.mockk
import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.Test

class OperationsDependantesTest {

    @Test
    fun CT_Division1() {
        val mockProvider = mockk<IntProvider>()
        every { mockProvider.getParam() } returnsMany listOf(10, 2)
        val calcul = OperationsDependantes(mockProvider)
        assertEquals(5.0f, calcul.diviserNaturelConsole())
    }
    @Test
    fun CT_DivisionParZero() {
        val mockProvider = mockk<IntProvider>()
        every { mockProvider.getParam() } returnsMany listOf(10, 0)
        val calcul = OperationsDependantes(mockProvider)
        assertThrows(ArithmeticException::class.java) {
            calcul.diviserNaturelConsole()
        }
    }
    @Test
    fun CT_DivisionNegatif() {
        val mockProvider = mockk<IntProvider>()
        every { mockProvider.getParam() } returnsMany listOf(-10, 2)
        val calcul = OperationsDependantes(mockProvider)
        assertThrows(ArithmeticException::class.java) {
            calcul.diviserNaturelConsole()
        }
    }
    @Test
    fun CT_Factorielle1() {
        val calcul = OperationsDependantes(OperationsDependantesStub())
        assertEquals(120, calcul.factorielleConsole())
    }
    @Test
    fun CT_Factorielle2() {
        val mock = mockk<IntProvider>()
        every { mock.getParam() } returns 4
        val calcul = OperationsDependantes(mock)
        assertEquals(24, calcul.factorielleConsole())
    }
}