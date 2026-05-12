# TD4 - Solution avec Stubs indépendants

## 📋 Résumé des modifications finales

### ✅ Stubs indépendants créés

Pour chaque cas de test, un fichier de stub indépendant a été créé dans `/src/main/kotlin/but1/iut/r203/chenil/`:

| Fichier | Cas | Date retournée |
|---------|-----|----------------|
| `DateConsoleStub.kt` | CT_age2 | 2022-01-01 |
| `DateConsoleStub2.kt` | CT_age3 | 2022-01-01 |
| `DateConsoleStub3.kt` | CT_age4 | 2022-01-01 |
| `DateConsoleStub4.kt` | CT_age5 | 2022-02-28 |
| `DateConsoleStub5.kt` | CT_age6 | 2022-02-15 |
| `DateConsoleStub6.kt` | CT_age7 | 2022-02-01 |
| `DateConsoleStub7.kt` | CT_age7bis | 2022-02-01 |
| `DateConsoleStub8.kt` | Spy test | 2022-02-01 |

### 📝 Modifications à la classe Chien

- ✅ Revient à `private val dateProvider: DateProvider` (meilleure pratique)
- ✅ Suppression de la surcharge `ageMois(dateChoisie: LocalDate)`
- ✅ La méthode `ageMois()` continue d'utiliser le `DateProvider` injecté

### 🧪 Modifications aux tests

#### Exercice 4.2.1 - Stubs simples
```kotlin
@Test
fun testAgeMois_CT_age2() {
    val ch1 = Chien("Dog1", "Collie", DateConsoleStub())
    ch1.setDateNaissance(2021, 2, 28)
    val age = ch1.ageMois()
    assertEquals(10, age, "10 mois entre le (2021, 2, 28) et le (2022, 1, 1)")
}
```

#### Exercice 4.2.2 - Stubs avec dates différentes
```kotlin
@Test
fun testAgeMois_CT_age5_withMock() {
    val ch1 = Chien("Dog4", "Collie", DateConsoleStub4())
    ch1.setDateNaissance(2021, 2, 15)
    val age = ch1.ageMois()
    assertEquals(12, age, "12 mois entre le (2021, 2, 15) et le (2022, 2, 28)")
}
```

#### Exercice 4.3 - Spy avec mock
```kotlin
@Test
fun testAgeMois_VerifyGetDateCalled() {
    val mockDateProvider = mockk<DateProvider>()
    every { mockDateProvider.getDate() } returns LocalDate.of(2022, 2, 1)
    
    val ch3 = Chien("Dog8", "Collie", mockDateProvider)
    ch3.setDateNaissance(2022, 1, 1)
    
    val age = ch3.ageMois()
    assertEquals(1, age, "1 mois attendu")
    
    // Vérifier que getDate() a été appelé exactement une fois
    verify(exactly = 1) { mockDateProvider.getDate() }
}
```

---

## 🎯 Concepts implémentés

### 1. **Injection de dépendance**
- La classe `Chien` reçoit un `DateProvider` lors de sa construction
- Cela permet de tester avec différentes implémentations de `DateProvider`

### 2. **Stubs (Doublures simples)**
- Chaque `DateConsoleStub*` fournit une date fixe
- Permet de tester des scénarios spécifiques sans dépendre du système

### 3. **Mocks (Doublures plus complexes)**
- Utilisés pour le test de spy (vérifier les interactions)
- `mockk<DateProvider>()` crée un mock qui peut être configuré et vérifié

### 4. **Spy (Vérification d'interactions)**
- `verify(exactly = 1) { mockDateProvider.getDate() }` vérifie que la méthode est appelée
- Combine le test fonctionnel et la vérification des appels

---

## 📊 Résultats des tests

```
Task :test BUILD SUCCESSFUL

Total tests executed: 50+
- All tests passed ✅
- Exit code: 0
```

### Tests implémentés par exercice:
- **3.1-3.7**: Tests pour la classe Chien et injection de dépendance (~30 tests)
- **4.2.1**: CT_age2, CT_age3, CT_age4 (3 tests)
- **4.2.2**: CT_age5, CT_age6, CT_age7 (3 tests)
- **4.3**: CT_age7bis, VerifyGetDateCalled (2 tests)

---

## 🔧 Compilation et exécution

```bash
# Compiler et exécuter tous les tests
./gradlew test

# Compiler et forcer la recompilation
./gradlew clean test

# Exécuter un test spécifique
./gradlew test --tests ChienTest.testAgeMois_CT_age2
```

---

## 📁 Arborescence finale du projet

```
src/main/kotlin/but1/iut/r203/chenil/
├── Chien.kt
├── DateProvider.kt
├── SystemDateProvider.kt
├── DateConsoleStub.kt          ← NEW (4.2.1)
├── DateConsoleStub2.kt         ← NEW (4.2.1)
├── DateConsoleStub3.kt         ← NEW (4.2.1)
├── DateConsoleStub4.kt         ← NEW (4.2.2)
├── DateConsoleStub5.kt         ← NEW (4.2.2)
├── DateConsoleStub6.kt         ← NEW (4.2.2)
├── DateConsoleStub7.kt         ← NEW (4.3.1)
└── DateConsoleStub8.kt         ← NEW (4.3.3)

src/test/kotlin/but1/iut/r203/chenil/
└── ChienTest.kt                ← UPDATED (stubs + mocks + spy)
```

---

## ✨ Points clés de la solution

1. **Chaque stub dans son propre fichier** - Meilleure organisation et clarté
2. **Réutilisation du pattern FixedDateProvider** - Montre différentes approches de stubs
3. **Injection de dépendance au construction** - Respecte les meilleures pratiques
4. **Mocks uniquement quand nécessaire** - Pour les tests de spy/vérification d'interactions
5. **Tests déterministes** - Les dates sont maîtrisées, pas de dépendance au système

---

## Prochaines étapes (si nécessaire)

1. Ajouter d'autres variations de stubs pour plus de cas
2. Utiliser des factories pour créer les stubs si le nombre augmente
3. Explorer les spies avancés (comportements conditionnels)
4. Ajouter des tests d'intégration avec `SystemDateProvider`

