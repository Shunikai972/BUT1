# ⚠️ Erreurs corrigées - TD4

## ❌ Ce qui a été fait de travers

J'ai créé **8 fichiers de stub** (`DateConsoleStub.kt`, `DateConsoleStub2.kt`, etc.) ce qui était une mauvaise compréhension de l'énoncé.

### Pourquoi c'était mauvais?

**Énoncé disait**:
> "Les 3 cas de test suivants ont besoin d'une date du jour différente à chaque fois : cela serait **laborieux avec trois autres stubs**."

**Ce que j'ai fait** ❌:
- Créé 8 stubs différents
- Un pour chaque cas
- C'était exactement ce que l'énoncé disait de NE PAS faire!

---

## ✅ La bonne approche (CORRECTION)

### Pour 4.2.1 (CT_age2, CT_age3, CT_age4)
**Approche**: Créer UN SEUL stub réutilisable

```kotlin
// DateProvider20220101.kt (UN SEUL FICHIER)
class DateProvider20220101 : DateProvider {
    override fun getDate(): LocalDate = LocalDate.of(2022, 1, 1)
}
```

**Utilisation** (réutiliser le même stub pour les 3 cas):
```kotlin
@Test fun testAgeMois_CT_age2() {
    val ch1 = Chien("Dog1", "Collie", DateProvider20220101())  // ← Même stub
    // ...
}

@Test fun testAgeMois_CT_age3() {
    val ch2 = Chien("Dog2", "Collie", DateProvider20220101())  // ← Même stub
    // ...
}
```

✅ **Avantage**: 1 stub pour 3 cas au lieu de 3 stubs

---

### Pour 4.2.2 (CT_age5, CT_age6, CT_age7)
**Approche**: Utiliser des MOCKS mockk configurables

```kotlin
@Test fun testAgeMois_CT_age5_withMock() {
    val mockDateProvider = mockk<DateProvider>()
    every { mockDateProvider.getDate() } returns LocalDate.of(2022, 2, 28)
    // ...
}

@Test fun testAgeMois_CT_age6_withMock() {
    val mockDateProvider = mockk<DateProvider>()
    every { mockDateProvider.getDate() } returns LocalDate.of(2022, 2, 15)  // ← Autre date
    // ...
}
```

✅ **Avantage**: Pas besoin de créer `DateConsoleStub4`, `DateConsoleStub5`, `DateConsoleStub6`. 
Un mock peut avoir plusieurs configurations!

---

## 📊 Comparaison avant/après

| Approche | Avant ❌ | Après ✅ |
|----------|---------|---------|
| **4.2.1** | 3 stubs | 1 stub |
| **4.2.2** | 3 stubs | Mocks mockk |
| **Total** | 8 fichiers de stub | 1 fichier de stub + mocks |

---

## 🎯 Résumé des corrections

**Supprimé** ❌:
- `DateConsoleStub.kt`
- `DateConsoleStub2.kt`
- `DateConsoleStub3.kt`
- `DateConsoleStub4.kt`
- `DateConsoleStub5.kt`
- `DateConsoleStub6.kt`
- `DateConsoleStub7.kt`
- `DateConsoleStub8.kt`

**Créé** ✅:
- `DateProvider20220101.kt` (1 seul stub pour 4.2.1)

**Modifié** ✅:
- `ChienTest.kt`: 
  - 4.2.1 utilise maintenant `DateProvider20220101()`
  - 4.2.2 utilise maintenant des `mockk<DateProvider>()`
  - 4.3 utilise des mocks avec `verify()`

---

## 📝 Clés de compréhension

### Stub
- ✅ Utilisez si la date est TOUJOURS la même
- ✅ Exemple: 4.2.1 où tous les cas utilisent 2022-01-01
- ❌ Ne créez pas plusieurs stubs pour une même date

### Mock
- ✅ Utilisez si la date change entre les cas
- ✅ Configurez avec `every { ... } returns ...`
- ✅ Vérifiez les appels avec `verify(exactly = X) { ... }`
- ✅ Exemple: 4.2.2 où chaque cas a une date différente

---

## ✨ Résultat final

```
✅ 1 stub réutilisable pour 4.2.1
✅ Mocks mockk configurables pour 4.2.2 et 4.3
✅ Code plus simple et plus maintenable
✅ Tous les tests passent (exit code: 0)
```

---

## 📚 Documentation pertinente

Voir **SOLUTION_CORRECTE_TD4.md** pour la solution complète et correcte.

