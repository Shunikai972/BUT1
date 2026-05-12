# TD4 - Résumé des modifications

## Fichiers modifiés

### 1. `src/main/kotlin/but1/iut/r203/chenil/Chien.kt`

**Modifications:**
- Changement de `private val dateProvider` à `internal var dateProvider` (permet la modification après construction et l'accès depuis les tests)
- Ajout d'une surcharge de la méthode `ageMois()` qui accepte un paramètre `LocalDate`:
  ```kotlin
  fun ageMois(dateChoisie: LocalDate): Long {
      val ecart = Period.between(dateNaissance, dateChoisie)
      return ecart.years * 12L + ecart.months
  }
  ```

### 2. `src/test/kotlin/but1/iut/r203/chenil/ChienTest.kt`

**Modifications:**
- Ajout des imports pour mockk:
  ```kotlin
  import io.mockk.mockk
  import io.mockk.every
  import io.mockk.verify
  ```
- Ajout de variables de classe pour les tests 4.2 et 4.3:
  ```kotlin
  private lateinit var ch1: Chien
  private lateinit var ch2: Chien
  private lateinit var ch3: Chien
  ```
- Ajout de la méthode `@BeforeEach` pour initialiser les chiens
- Ajout de 9 nouveaux tests:
  - **Exercice 4.2.1**: `testAgeMois_CT_age2()`, `testAgeMois_CT_age3()`, `testAgeMois_CT_age4()` (avec stubs/paramètres)
  - **Exercice 4.2.2**: `testAgeMois_CT_age5_withMock()`, `testAgeMois_CT_age6_withMock()`, `testAgeMois_CT_age7_withMock()` (avec mocks mockk)
  - **Exercice 4.3**: `testAgeMois_CT_age7bis()`, `testAgeMois_VerifyGetDateCalled()` (avec spy/verify)

### 3. `build.gradle.kts`
- Aucune modification (la dépendance mockk était déjà présente)

---

## Récapitulatif des cas de test implémentés

### Exercice 4.2.1 - Stubs
| Cas | Naissance | Test Date | Âge attendu |
|-----|-----------|-----------|------------|
| CT_age2 | 2021-02-28 | 2022-01-01 | 10 mois |
| CT_age3 | 2021-12-31 | 2022-01-01 | 0 mois |
| CT_age4 | 2021-12-01 | 2022-01-01 | 1 mois |

### Exercice 4.2.2 - Mocks
| Cas | Naissance | Mock Date | Âge attendu |
|-----|-----------|-----------|------------|
| CT_age5 | 2021-02-15 | 2022-02-28 | 12 mois |
| CT_age6 | 2021-02-15 | 2022-02-15 | 12 mois |
| CT_age7 | 2021-02-15 | 2022-02-01 | 11 mois |

### Exercice 4.3 - Spy
| Cas | Naissance | Test Date | Âge attendu | Vérification |
|-----|-----------|-----------|------------|--------------|
| CT_age7bis | 2022-01-01 | 2022-02-01 | 1 mois | getDate() appelé 1x |

---

## Concepts testés

### Injection de dépendance
✅ La classe `Chien` accepte un `DateProvider` injecté, permettant de contrôler la source de la date dans les tests.

### Stubs
✅ La classe `FixedDateProvider` fournit une date fixe, idéale pour les tests avec des dates maîtrisées.
✅ Alternative: utiliser directement un paramètre `LocalDate` dans la méthode `ageMois(dateChoisie: LocalDate)`.

### Mocks
✅ Utilisation de `mockk<DateProvider>()` pour créer des mocks.
✅ Configuration avec `every { ... } returns ...` pour définir le comportement.

### Spy/Verify
✅ Utilisation de `verify(exactly = 1) { mockDateProvider.getDate() }` pour vérifier que la méthode a été appelée.

---

## Résultats des tests

✅ **Compilation**: Succès
✅ **Exécution des tests**: Tous les tests passent (exit code: 0)
✅ **Total des tests**: 50+ (incluant tous les exercices 3.1-3.7 et 4.2-4.3)

---

## Commandes utiles

```bash
# Exécuter tous les tests
./gradlew test

# Exécuter un test spécifique
./gradlew test --tests ChienTest.testAgeMois_CT_age2

# Afficher plus de détails
./gradlew test --info
```

---

## Conclusions

Le TD4 démontre comment utiliser l'injection de dépendance, les stubs, les mocks et les spies pour écrire des tests unitaires robustes et maintenables. Les modifications minimales à la classe `Chien` (rendre `dateProvider` modifiable et ajout d'une surcharge de méthode) permettent une testabilité complète sans compromettre la conception de la classe.
