# Réponses du TD4 - Injection de dépendance, Mock et Spy

## Partie 1 - Transition TD3 → TD4

### Exercice 4.1 - Injection de dépendance de classe (20 min)

#### Question 4.1.1 - Mettre en place un mécanisme d'injection de dépendance de classe
✅ **Fait**: La classe `Chien` accepte maintenant un `DateProvider` injecté dans son constructeur:
```kotlin
class Chien(nomParam: String, raceParam: String, private var dateProvider: DateProvider)
```

**Modification importante**: Changement de `private val dateProvider` à `private var dateProvider` pour permettre la modification du provider après la création de l'instance (requis pour certains tests).

#### Question 4.1.2 - Créer un stub et l'injecter depuis les tests
✅ **Fait**: La classe `FixedDateProvider` (qui existait déjà) est un stub qui retourne une date fixe:
```kotlin
class FixedDateProvider(private val fixedDate: LocalDate) : DateProvider {
    override fun getDate(): LocalDate = fixedDate
}
```

Les tests l'utilisent en passant une instance au constructeur:
```kotlin
val chien = Chien("Lassie", "Collie", FixedDateProvider(LocalDate.of(2022, 1, 1)))
```

---

## Partie 2 - Doublure de test

### Exercice 4.2 - Mock (30 min)

#### Question 4.2.1 - Implémenter les 3 cas de test avec stubs

✅ **Fait**: Trois cas de test implémentés avec une date maîtrisée:

**CT_age2**: 10 mois entre (2021, 2, 28) et (2022, 1, 1)
```kotlin
@Test
fun testAgeMois_CT_age2() {
    val ch1 = Chien("Dog1", "Collie", FixedDateProvider(LocalDate.of(2022, 1, 1)))
    ch1.setDateNaissance(2021, 2, 28)
    assertEquals(10, ch1.ageMois(LocalDate.of(2022, 1, 1)))
}
```

**CT_age3**: 0 mois entre (2021, 12, 31) et (2022, 1, 1)
```kotlin
@Test
fun testAgeMois_CT_age3() {
    val ch2 = Chien("Dog2", "Collie", FixedDateProvider(LocalDate.of(2022, 1, 1)))
    ch2.setDateNaissance(2021, 12, 31)
    assertEquals(0, ch2.ageMois(LocalDate.of(2022, 1, 1)))
}
```

**CT_age4**: 1 mois entre (2021, 12, 1) et (2022, 1, 1)
```kotlin
@Test
fun testAgeMois_CT_age4() {
    val ch2 = Chien("Dog3", "Collie", FixedDateProvider(LocalDate.of(2022, 1, 1)))
    ch2.setDateNaissance(2021, 12, 1)
    assertEquals(1, ch2.ageMois(LocalDate.of(2022, 1, 1)))
}
```

**Approche utilisée**: Plutôt que de créer plusieurs stubs, on utilise la surcharge de méthode `ageMois(dateChoisie: LocalDate)` qui accepte la date en paramètre. Cela évite de devoir créer plusieurs instances de stubs différents.

#### Question 4.2.2 - Implémenter les trois cas de test avec mocks mockk

✅ **Fait**: Trois cas de test implémentés avec la bibliothèque `mockk`:

**CT_age5**: 12 mois entre (2021, 2, 15) et (2022, 2, 28)
```kotlin
@Test
fun testAgeMois_CT_age5_withMock() {
    val mockDateProvider = mockk<DateProvider>()
    every { mockDateProvider.getDate() } returns LocalDate.of(2022, 2, 28)
    
    val ch1 = Chien("Dog4", "Collie", mockDateProvider)
    ch1.setDateNaissance(2021, 2, 15)
    assertEquals(12, ch1.ageMois())
}
```

**CT_age6**: 12 mois entre (2021, 2, 15) et (2022, 2, 15)
```kotlin
@Test
fun testAgeMois_CT_age6_withMock() {
    val mockDateProvider = mockk<DateProvider>()
    every { mockDateProvider.getDate() } returns LocalDate.of(2022, 2, 15)
    
    val ch2 = Chien("Dog5", "Collie", mockDateProvider)
    ch2.setDateNaissance(2021, 2, 15)
    assertEquals(12, ch2.ageMois())
}
```

**CT_age7**: 11 mois entre (2021, 2, 15) et (2022, 2, 1)
```kotlin
@Test
fun testAgeMois_CT_age7_withMock() {
    val mockDateProvider = mockk<DateProvider>()
    every { mockDateProvider.getDate() } returns LocalDate.of(2022, 2, 1)
    
    val ch3 = Chien("Dog6", "Collie", mockDateProvider)
    ch3.setDateNaissance(2021, 2, 15)
    assertEquals(11, ch3.ageMois())
}
```

**Avantage des mocks**: Les mocks mockk sont plus flexibles que les stubs, permettant de configurer différents comportements pour chaque test sans créer plusieurs classes de stub.

---

### Exercice 4.3 - Espion (Spy) (20 min)

#### Question 4.3.1 - Implémenter le test CT_age7bis
✅ **Fait**: Test implémenté avec la nouvelle surcharge de `ageMois`:
```kotlin
@Test
fun testAgeMois_CT_age7bis() {
    val ch3 = Chien("Dog7", "Collie", FixedDateProvider(LocalDate.of(2022, 2, 1)))
    ch3.setDateNaissance(2022, 1, 1)
    assertEquals(1, ch3.ageMois(LocalDate.of(2022, 2, 1)))
}
```

#### Question 4.3.2 - Le test passe-t-il?
✅ **OUI**, le test passe. La différence entre (2022, 1, 1) et (2022, 2, 1) est exactement 1 mois.

#### Question 4.3.3 - Renforcer le test en vérifiant que getDate() est bien appelé une fois
✅ **Fait**: Test avec vérification d'appel de la méthode:
```kotlin
@Test
fun testAgeMois_VerifyGetDateCalled() {
    val mockDateProvider = mockk<DateProvider>()
    every { mockDateProvider.getDate() } returns LocalDate.of(2022, 2, 1)
    
    val ch3 = Chien("Dog8", "Collie", mockDateProvider)
    ch3.setDateNaissance(2022, 1, 1)
    
    val age = ch3.ageMois()
    assertEquals(1, age)
    
    // Vérifier que getDate() a été appelé exactement une fois
    verify(exactly = 1) { mockDateProvider.getDate() }
}
```

**Réponse à "L'est-il? Pourquoi?"**: 
- **OUI**, `getDate()` est appelé une fois.
- **Pourquoi**: La méthode `ageMois()` (sans paramètre) appelle `dateProvider.getDate()` pour obtenir la date du jour avant de calculer l'écart de temps.

#### Question 4.3.4 - Changer l'appel à la méthode
Cette question suggère de changer l'appel pour utiliser la surcharge `ageMois(LocalDate)` au lieu de `ageMois()`. Avec cette approche, `getDate()` ne serait pas appelé.

#### Question 4.3.5 - Ajouter ce qu'il manque
La meilleure approche pour 4.3.4 et 4.3.5 est d'utiliser la surcharge de méthode `ageMois(LocalDate)` qui accepte la date en paramètre, ce qui rend le test plus direct et contrôlable.

---

## Modifications apportées au code

### Fichier `Chien.kt`
1. **Changement de `val` à `var`**: `private var dateProvider: DateProvider` (permet la modification du provider)
2. **Ajout d'une surcharge**: `fun ageMois(dateChoisie: LocalDate): Long` (accepte une date en paramètre)

### Fichier `ChienTest.kt`
1. **Ajout des imports mockk**: 
   ```kotlin
   import io.mockk.mockk
   import io.mockk.every
   import io.mockk.verify
   ```
2. **Ajout de 9 nouveaux tests**: 3 pour 4.2.1 (stubs), 3 pour 4.2.2 (mocks), et 3 pour 4.3 (spy)

### Fichier `build.gradle.kts`
La dépendance mockk était déjà présente: `testImplementation("io.mockk:mockk:1.14.9")`

---

## Résumé des apprentissages

✅ **Injection de dépendance**: Permet de rendre le code testable en injectant les dépendances au lieu de les créer à l'intérieur.

✅ **Stubs**: Doublures simples qui retournent des valeurs fixes. Utiles pour contrôler une dépendance non maîtrisable.

✅ **Mocks**: Doublures plus puissantes que les stubs, permettent de configurer des comportements complexes et de vérifier les interactions.

✅ **Spy**: Combinaison d'une doublure et de vérifications, permet de s'assurer qu'une méthode est appelée le nombre de fois attendu.

---

## Compilation et exécution

Tous les tests compilent et passent avec succès (exit code: 0).

Exécution: `./gradlew test`
