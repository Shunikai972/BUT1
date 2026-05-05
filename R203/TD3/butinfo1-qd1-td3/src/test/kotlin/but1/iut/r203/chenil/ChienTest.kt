package but1.iut.r203.chenil

import org.junit.jupiter.api.Test
import org.junit.jupiter.api.DisplayName
import org.junit.jupiter.api.BeforeEach
import java.time.LocalDate
import java.time.DateTimeException
import kotlin.test.assertEquals
import kotlin.test.assertThrows
import kotlin.test.assertTrue

/**
 * Stub pour tester avec une date fixe
 * Ceci est un "test double" utilisé pour rendre les tests déterministes
 */
class FixedDateProvider(private val fixedDate: LocalDate) : DateProvider {
    override fun getDate(): LocalDate = fixedDate
}

/**
 * Suite de tests pour la classe Chien
 * Démontre comment utiliser l'injection de dépendance pour tester du code
 * qui dépend de la date et du temps
 */
@DisplayName("Tests de la classe Chien avec injection de dépendance")
class ChienTest {

    // ===== EXERCICE 3.1: Tests fonctionnels de setDateNaissance =====
    @Test
    @DisplayName("3.1.1 - setDateNaissance avec une date valide")
    fun testSetDateNaissanceValid() {
        val fixedDate = LocalDate.of(2023, 5, 15)
        val chien = Chien("Lassie", "Collie", FixedDateProvider(fixedDate))
        
        chien.setDateNaissance(2021, 2, 28)
        // Le test passe si pas d'exception levée
        assertTrue(true)
    }

    @Test
    @DisplayName("3.1.1 - setDateNaissance avec date invalide (mois 13)")
    fun testSetDateNaissanceInvalidMonth() {
        val fixedDate = LocalDate.of(2023, 5, 15)
        val chien = Chien("Lassie", "Collie", FixedDateProvider(fixedDate))
        
        assertThrows<DateTimeException> {
            chien.setDateNaissance(2023, 13, 32)
        }
    }

    @Test
    @DisplayName("3.1.1 - setDateNaissance avec date invalide (jour 32)")
    fun testSetDateNaissanceInvalidDay() {
        val fixedDate = LocalDate.of(2023, 5, 15)
        val chien = Chien("Lassie", "Collie", FixedDateProvider(fixedDate))
        
        assertThrows<DateTimeException> {
            chien.setDateNaissance(2023, 2, 30)
        }
    }

    @Test
    @DisplayName("3.1.1 - toString contient le nom du chien")
    fun testToStringContainsName() {
        val fixedDate = LocalDate.of(2023, 5, 15)
        val chien = Chien("Médor", "Chihuahua", FixedDateProvider(fixedDate))
        chien.setDateNaissance(2021, 2, 15)
        
        val str = chien.toString()
        assertTrue(str.contains("Médor"), "Le toString doit contenir le nom du chien")
    }

    @Test
    @DisplayName("3.1.1 - toString contient la race du chien")
    fun testToStringContainsBreed() {
        val fixedDate = LocalDate.of(2023, 5, 15)
        val chien = Chien("Médor", "Chihuahua", FixedDateProvider(fixedDate))
        chien.setDateNaissance(2021, 2, 15)
        
        val str = chien.toString()
        assertTrue(str.contains("Chihuahua"), "Le toString doit contenir la race du chien")
    }

    // ===== EXERCICE 3.2 & 3.3: Tests d'âge avec dates fixées =====
    @Test
    @DisplayName("3.3 - Calcul d'âge: exactement 1 mois")
    fun testAgeMoisExactlyOneMonth() {
        val birthDate = LocalDate.of(2023, 5, 15)
        val testDate = LocalDate.of(2023, 6, 15)
        
        val chien = Chien("Lassie", "Collie", FixedDateProvider(testDate))
        chien.setDateNaissance(2023, 5, 15)
        
        val age = chien.ageMois()
        assertEquals(1, age, "À 1 mois de jour, l'âge doit être 1 mois")
    }

    @Test
    @DisplayName("3.3 - Calcul d'âge: exactement 1 an")
    fun testAgeMoisExactlyOneYear() {
        val testDate = LocalDate.of(2022, 5, 15)
        
        val chien = Chien("Lassie", "Collie", FixedDateProvider(testDate))
        chien.setDateNaissance(2021, 5, 15)
        
        val age = chien.ageMois()
        assertEquals(12, age, "À 1 an de jour, l'âge doit être 12 mois")
    }

    @Test
    @DisplayName("3.3 - Calcul d'âge: 1 an et 3 mois")
    fun testAgeMoisOneYearThreeMonths() {
        val testDate = LocalDate.of(2023, 8, 15)
        
        val chien = Chien("Lassie", "Collie", FixedDateProvider(testDate))
        chien.setDateNaissance(2022, 5, 15)
        
        val age = chien.ageMois()
        assertEquals(15, age, "1 an et 3 mois = 15 mois")
    }

    @Test
    @DisplayName("3.3 - Calcul d'âge: 2 ans et 6 mois")
    fun testAgeMoisTwoYearsSixMonths() {
        val testDate = LocalDate.of(2024, 11, 15)
        
        val chien = Chien("Lassie", "Collie", FixedDateProvider(testDate))
        chien.setDateNaissance(2022, 5, 15)
        
        val age = chien.ageMois()
        assertEquals(30, age, "2 ans et 6 mois = 30 mois")
    }

    // ===== EXERCICE 3.4: Cas limites (boundary) =====
    @Test
    @DisplayName("3.4 - Cas limite: nouveau-né (0 mois)")
    fun testAgeMoisNewborn() {
        val testDate = LocalDate.of(2023, 5, 15)
        
        val chien = Chien("Puppy", "Poodle", FixedDateProvider(testDate))
        chien.setDateNaissance(2023, 5, 15)
        
        val age = chien.ageMois()
        assertEquals(0, age, "À la date de naissance, l'âge doit être 0 mois")
    }

    @Test
    @DisplayName("3.4 - Cas limite: un jour avant le mois complet")
    fun testAgeMoisJustBeforeOneMonth() {
        val testDate = LocalDate.of(2023, 6, 14)
        
        val chien = Chien("Lassie", "Collie", FixedDateProvider(testDate))
        chien.setDateNaissance(2023, 5, 15)
        
        val age = chien.ageMois()
        assertEquals(0, age, "Un jour avant le mois complet, l'âge doit toujours être 0 mois")
    }

    @Test
    @DisplayName("3.4 - Cas limite: le jour exact du mois")
    fun testAgeMoisExactMonthDay() {
        val testDate = LocalDate.of(2023, 6, 15)
        
        val chien = Chien("Lassie", "Collie", FixedDateProvider(testDate))
        chien.setDateNaissance(2023, 5, 15)
        
        val age = chien.ageMois()
        assertEquals(1, age, "Le jour exact du mois complet, l'âge doit être 1 mois")
    }

    @Test
    @DisplayName("3.4 - Cas limite: année bissextile (29 février)")
    fun testAgeMoisLeapYearFebruary29() {
        val testDate = LocalDate.of(2024, 2, 29)
        
        val chien = Chien("Lassie", "Collie", FixedDateProvider(testDate))
        chien.setDateNaissance(2024, 2, 29)
        
        val age = chien.ageMois()
        assertEquals(0, age, "À la date de naissance (29 février), l'âge doit être 0 mois")
    }

    @Test
    @DisplayName("3.4 - Cas limite: année bissextile (anniversaire 29 février)")
    fun testAgeMoisLeapYearBirthdayInLeapYear() {
        val testDate = LocalDate.of(2025, 2, 28)
        
        val chien = Chien("Lassie", "Collie", FixedDateProvider(testDate))
        chien.setDateNaissance(2024, 2, 29)
        
        val age = chien.ageMois()
        assertEquals(11, age, "De 2024-02-29 à 2025-02-28 = 11 mois")
    }

    // ===== EXERCICE 3.5: Tests avec DateProvider testable =====
    @Test
    @DisplayName("3.5 - ageMoisDate avec une date fixe")
    fun testAgeMoisDateWithFixedProvider() {
        val testDate = LocalDate.of(2023, 8, 15)
        
        val chien = Chien("Lassie", "Collie", FixedDateProvider(testDate))
        chien.setDateNaissance(2023, 5, 15)
        
        val age = chien.ageMoisDate(FixedDateProvider(testDate))
        assertEquals(3, age, "3 mois de différence")
    }

    @Test
    @DisplayName("3.5 - Comparer ageMois et ageMoisDate avec même date")
    fun testAgeMoisMatchesAgeMoisDate() {
        val testDate = LocalDate.of(2023, 8, 15)
        val dateProvider = FixedDateProvider(testDate)
        
        val chien = Chien("Lassie", "Collie", dateProvider)
        chien.setDateNaissance(2023, 5, 15)
        
        val age1 = chien.ageMois()
        val age2 = chien.ageMoisDate(dateProvider)
        assertEquals(age1, age2, "Les deux méthodes doivent donner le même résultat avec la même date")
    }

    // ===== EXERCICE 3.6: Problèmes de testabilité identifiés et résolus =====
    @Test
    @DisplayName("3.6 - Le test ne dépend plus de la date actuelle")
    fun testIndependentFromSystemDate() {
        // Avant (non testable): ageMois() utilisait LocalDate.now()
        // Après (testable): on peut injecter n'importe quelle date
        
        val chien = Chien("Lassie", "Collie", FixedDateProvider(LocalDate.of(2023, 5, 15)))
        chien.setDateNaissance(2020, 1, 1)
        
        val age = chien.ageMois()
        // Le résultat est déterministe peu importe quand le test s'exécute
        assertEquals(40, age) // Exact entre 2020-01-01 et 2023-05-15
    }

    // ===== EXERCICE 3.7: Injection de dépendance de classe =====
    @Test
    @DisplayName("3.7.1 - Chien accepte une dépendance DateProvider")
    fun testChienAcceptesDateProvider() {
        val dateProvider = FixedDateProvider(LocalDate.of(2023, 5, 15))
        val chien = Chien("Lassie", "Collie", dateProvider)
        
        chien.setDateNaissance(2021, 2, 28)
        val age = chien.ageMois()
        
        assertTrue(age > 0, "Le chien doit avoir un âge calculé")
    }

    @Test
    @DisplayName("3.7.3 - Stub qui retourne toujours (2025, 5, 12)")
    fun testWithFixedDateStub() {
        // Ceci est le stub mentionné dans l'exercice 3.7.3
        val stubDate = LocalDate.of(2025, 5, 12)
        val dateProvider = FixedDateProvider(stubDate)
        
        val chien = Chien("Lassie", "Collie", dateProvider)
        chien.setDateNaissance(2025, 5, 12)
        
        val age = chien.ageMois()
        assertEquals(0, age, "Le jour même de la naissance, l'âge doit être 0")
    }

    @Test
    @DisplayName("3.7.4 - Les tests passent avec le stub")
    fun testPassWithStub() {
        val stubDate = LocalDate.of(2025, 5, 12)
        val dateProvider = FixedDateProvider(stubDate)
        
        val chien = Chien("Médor", "Berger Allemand", dateProvider)
        chien.setDateNaissance(2023, 1, 1)
        
        val age = chien.ageMois()
        assertEquals(28, age) // 2023-01-01 à 2025-05-12 = 28 mois
    }

    // ===== Tests supplémentaires pour couvrir plus de scénarios =====
    @Test
    @DisplayName("Tests supplémentaires - Plusieurs chiens avec âges différents")
    fun testMultipleDogsWithDifferentAges() {
        val testDate = LocalDate.of(2023, 6, 15)
        val dateProvider = FixedDateProvider(testDate)
        
        val chien1 = Chien("Lassie", "Collie", dateProvider)
        chien1.setDateNaissance(2023, 6, 15)
        val age1 = chien1.ageMois()
        assertEquals(0, age1)
        
        val chien2 = Chien("Rex", "Berger Allemand", dateProvider)
        chien2.setDateNaissance(2022, 6, 15)
        val age2 = chien2.ageMois()
        assertEquals(12, age2)
        
        val chien3 = Chien("Max", "Poodle", dateProvider)
        chien3.setDateNaissance(2021, 6, 15)
        val age3 = chien3.ageMois()
        assertEquals(24, age3)
    }

    @Test
    @DisplayName("Tests supplémentaires - Février année non-bissextile")
    fun testFebruaryNonLeapYear() {
        val testDate = LocalDate.of(2023, 3, 28)
        val dateProvider = FixedDateProvider(testDate)
        
        val chien = Chien("Lassie", "Collie", dateProvider)
        chien.setDateNaissance(2023, 2, 28)
        
        val age = chien.ageMois()
        assertEquals(1, age, "De février 28 à mars 28 = 1 mois")
    }

    @Test
    @DisplayName("Tests supplémentaires - Grand écart de temps")
    fun testLargeTimeGap() {
        val testDate = LocalDate.of(2050, 12, 31)
        val dateProvider = FixedDateProvider(testDate)
        
        val chien = Chien("Lassie", "Collie", dateProvider)
        chien.setDateNaissance(2000, 1, 1)
        
        val age = chien.ageMois()
        assertEquals(611, age, "De 2000-01-01 à 2050-12-31 ≈ 51 ans = 612 mois")
    }
}
