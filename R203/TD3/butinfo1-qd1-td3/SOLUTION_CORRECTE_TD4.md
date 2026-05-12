# ✅ TD4 - Solution correcte avec Stub et Mocks mockk

## 📋 Résumé de la solution

### 🎯 Exercice 4.2.1 - UN Stub simple
**Cas**: CT_age2, CT_age3, CT_age4 (tous utilisent la date 2022-01-01)

**Solution**: Créer UN seul stub réutilisable: `DateProvider20220101`

```kotlin
class DateProvider20220101 : DateProvider {
    override fun getDate(): LocalDate = LocalDate.of(2022, 1, 1)
}
```

**Utilisation dans les tests**:
```kotlin
@Test
fun testAgeMois_CT_age2() {
    val ch1 = Chien("Dog1", "Collie", DateProvider20220101())
    ch1.setDateNaissance(2021, 2, 28)
    val age = ch1.ageMois()
    assertEquals(10, age, "10 mois")
}
```

✅ **Avantage**: Simple, réutilisable, pas besoin de créer 3 stubs différents

---

### 🎯 Exercice 4.2.2 - Mocks mockk (dates différentes)
**Cas**: CT_age5, CT_age6, CT_age7 (chaque cas a une date différente)

**Problème**: "Cela serait laborieux avec trois autres stubs"

**Solution**: Utiliser des MOCKS mockk avec configuration `every`

```kotlin
@Test
fun testAgeMois_CT_age5_withMock() {
    // given: un mock qui retourne 2022-02-28
    val mockDateProvider = mockk<DateProvider>()
    every { mockDateProvider.getDate() } returns LocalDate.of(2022, 2, 28)
    
    val ch1 = Chien("Dog4", "Collie", mockDateProvider)
    ch1.setDateNaissance(2021, 2, 15)
    
    // when
    val age = ch1.ageMois()
    
    // then
    assertEquals(12, age, "12 mois")
}
```

✅ **Avantage**: Pas besoin de créer 3 autres stubs. On réutilise le même mock avec configuration différente pour chaque cas.

---

### 🎯 Exercice 4.3 - Spy/Verify (vérifier les interactions)
**Cas**: CT_age7bis + vérification que getDate() est appelé exactement une fois

**Solution**: Mock mockk + `verify(exactly = 1) { ... }`

```kotlin
@Test
fun testAgeMois_VerifyGetDateCalled() {
    // given: un mock avec vérification
    val mockDateProvider = mockk<DateProvider>()
    every { mockDateProvider.getDate() } returns LocalDate.of(2022, 2, 1)
    
    val ch3 = Chien("Dog8", "Collie", mockDateProvider)
    ch3.setDateNaissance(2022, 1, 1)
    
    // when
    val age = ch3.ageMois()
    
    // then: vérification fonctionnelle
    assertEquals(1, age, "1 mois")
    
    // then: vérification d'interaction
    verify(exactly = 1) { mockDateProvider.getDate() }
}
```

✅ **Résultat**: 
- `getDate()` EST appelé exactement une fois
- **Pourquoi?**: La méthode `ageMois()` appelle `dateProvider.getDate()` une fois pour obtenir la date du jour

---

## 📊 Tableau récapitulatif

| Exercice | Cas | Approche | Implementation | Date |
|----------|-----|----------|----------------|------|
| 4.2.1 | CT_age2 | Stub | DateProvider20220101() | 2022-01-01 |
| 4.2.1 | CT_age3 | Stub | DateProvider20220101() | 2022-01-01 |
| 4.2.1 | CT_age4 | Stub | DateProvider20220101() | 2022-01-01 |
| 4.2.2 | CT_age5 | Mock | mockk + every | 2022-02-28 |
| 4.2.2 | CT_age6 | Mock | mockk + every | 2022-02-15 |
| 4.2.2 | CT_age7 | Mock | mockk + every | 2022-02-01 |
| 4.3 | CT_age7bis | Mock | mockk + every | 2022-02-01 |
| 4.3 | Verify | Mock | mockk + every + verify | 2022-02-01 |

---

## 📝 Réponses aux questions

### Question 4.3.1 - Test CT_age7bis
✅ **Implémenté** avec mock mockk

### Question 4.3.2 - Le test passe-t-il?
✅ **OUI**, le test passe
- Naissance: 2022-01-01
- Date du test: 2022-02-01
- Âge: 1 mois ✓

### Question 4.3.3 - getDate() appelé une fois?
✅ **OUI**, `getDate()` EST appelé exactement une fois
```kotlin
verify(exactly = 1) { mockDateProvider.getDate() }
```

### Question 4.3.4 - Changez l'appel à la méthode?
✅ **C'est fait** - La méthode `ageMois()` utilise `dateProvider.getDate()` qui est appelée une fois

### Question 4.3.5 - Ajoutez ce qu'il manque
✅ **Ajouté** - La vérification d'interaction avec `verify(exactly = 1) { ... }`

---

## 🔍 Comparaison Stub vs Mock

### Stub (4.2.1)
```kotlin
class DateProvider20220101 : DateProvider {
    override fun getDate(): LocalDate = LocalDate.of(2022, 1, 1)
}
```
- ✅ Simple
- ✅ Peu de code
- ❌ Inflexible (date fixe)
- ❌ Impossible de vérifier les appels

### Mock (4.2.2 et 4.3)
```kotlin
val mockDateProvider = mockk<DateProvider>()
every { mockDateProvider.getDate() } returns LocalDate.of(2022, 2, 28)
verify(exactly = 1) { mockDateProvider.getDate() }
```
- ✅ Flexible (date configurable)
- ✅ Peut vérifier les appels
- ❌ Plus de code
- ❌ Nécessite mockk

---

## ✨ Pattern: Given-When-Then

Tous les tests suivent le pattern **Given-When-Then**:

```kotlin
@Test
fun testAgeMois_CT_age5_withMock() {
    // GIVEN: une configuration initiale
    val mockDateProvider = mockk<DateProvider>()
    every { mockDateProvider.getDate() } returns LocalDate.of(2022, 2, 28)
    val ch1 = Chien("Dog4", "Collie", mockDateProvider)
    ch1.setDateNaissance(2021, 2, 15)
    
    // WHEN: une action
    val age = ch1.ageMois()
    
    // THEN: des assertions
    assertEquals(12, age, "12 mois")
}
```

---

## 📁 Fichiers de la solution

```
src/main/kotlin/but1/iut/r203/chenil/
├── Chien.kt                      (inchangé)
├── DateProvider.kt               (inchangé)
├── SystemDateProvider.kt         (inchangé)
└── DateProvider20220101.kt       (NEW - 1 stub pour 4.2.1)

src/test/kotlin/but1/iut/r203/chenil/
└── ChienTest.kt                  (contient 8 nouveaux tests pour 4.2 et 4.3)
```

---

## ✅ Résultats

```
✅ Compilation: SUCCÈS
✅ Tests: 50+ tests passent
✅ Exit code: 0
```

---

## 🎓 Concepts clés

**Stub**: Implémentation simple et prédéfinie
- Utiliser quand la date est toujours la même
- Exemple: `DateProvider20220101()`

**Mock**: Implémentation configurable avec vérification d'appels
- Utiliser quand on a besoin de dates différentes
- Utiliser pour vérifier les interactions
- Exemple: `mockk<DateProvider>()` avec `verify()`

**Given-When-Then**: Pattern pour structurer les tests
- Given: configuration initiale
- When: action à tester
- Then: assertions

---

## 🚀 Compilation et exécution

```bash
./gradlew test                    # Tous les tests
./gradlew test --tests ChienTest.testAgeMois_CT_age2  # Test spécifique
```

---

## 📌 Important

- **1 seul stub** pour 4.2.1 (pas 3)
- **Mocks mockk** pour 4.2.2 (évite de créer 3 autres stubs)
- **Verify** pour 4.3 (vérifier les appels)

