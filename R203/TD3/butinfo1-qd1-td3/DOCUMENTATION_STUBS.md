# 📚 Documentation des Stubs créés

## Vue d'ensemble

8 stubs ont été créés pour tester différents cas. Chaque stub retourne une date fixe, ce qui permet de tester le calcul d'âge de manière déterministe.

---

## DateConsoleStub

```kotlin
class DateConsoleStub : DateProvider {
    override fun getDate(): LocalDate = LocalDate.of(2022, 1, 1)
}
```

**Cas de test**: CT_age2  
**Date fournie**: 2022-01-01  
**Test associé**: `testAgeMois_CT_age2()`  
**Calcul attendu**: Naissance 2021-02-28 → 10 mois  
**Exercice**: 4.2.1

---

## DateConsoleStub2

```kotlin
class DateConsoleStub2 : DateProvider {
    override fun getDate(): LocalDate = LocalDate.of(2022, 1, 1)
}
```

**Cas de test**: CT_age3  
**Date fournie**: 2022-01-01  
**Test associé**: `testAgeMois_CT_age3()`  
**Calcul attendu**: Naissance 2021-12-31 → 0 mois (moins d'un mois)  
**Exercice**: 4.2.1

---

## DateConsoleStub3

```kotlin
class DateConsoleStub3 : DateProvider {
    override fun getDate(): LocalDate = LocalDate.of(2022, 1, 1)
}
```

**Cas de test**: CT_age4  
**Date fournie**: 2022-01-01  
**Test associé**: `testAgeMois_CT_age4()`  
**Calcul attendu**: Naissance 2021-12-01 → 1 mois  
**Exercice**: 4.2.1

---

## DateConsoleStub4

```kotlin
class DateConsoleStub4 : DateProvider {
    override fun getDate(): LocalDate = LocalDate.of(2022, 2, 28)
}
```

**Cas de test**: CT_age5  
**Date fournie**: 2022-02-28  
**Test associé**: `testAgeMois_CT_age5_withMock()`  
**Calcul attendu**: Naissance 2021-02-15 → 12 mois  
**Exercice**: 4.2.2

**Note**: Ce stub utilise `2022-02-28` pour tester un cas spécifique avec février de l'année non-bissextile.

---

## DateConsoleStub5

```kotlin
class DateConsoleStub5 : DateProvider {
    override fun getDate(): LocalDate = LocalDate.of(2022, 2, 15)
}
```

**Cas de test**: CT_age6  
**Date fournie**: 2022-02-15  
**Test associé**: `testAgeMois_CT_age6_withMock()`  
**Calcul attendu**: Naissance 2021-02-15 → 12 mois (jour anniversaire exact)  
**Exercice**: 4.2.2

---

## DateConsoleStub6

```kotlin
class DateConsoleStub6 : DateProvider {
    override fun getDate(): LocalDate = LocalDate.of(2022, 2, 1)
}
```

**Cas de test**: CT_age7  
**Date fournie**: 2022-02-01  
**Test associé**: `testAgeMois_CT_age7_withMock()`  
**Calcul attendu**: Naissance 2021-02-15 → 11 mois (avant le jour anniversaire)  
**Exercice**: 4.2.2

**Note**: Cas limite - 1 jour avant le 12ème mois.

---

## DateConsoleStub7

```kotlin
class DateConsoleStub7 : DateProvider {
    override fun getDate(): LocalDate = LocalDate.of(2022, 2, 1)
}
```

**Cas de test**: CT_age7bis  
**Date fournie**: 2022-02-01  
**Test associé**: `testAgeMois_CT_age7bis()`  
**Calcul attendu**: Naissance 2022-01-01 → 1 mois  
**Exercice**: 4.3.1

**Identique à DateConsoleStub6** - Stub réutilisé pour le test CT_age7bis.

---

## DateConsoleStub8

```kotlin
class DateConsoleStub8 : DateProvider {
    override fun getDate(): LocalDate = LocalDate.of(2022, 2, 1)
}
```

**Cas de test**: Spy/Mock verification  
**Date fournie**: 2022-02-01  
**Test associé**: `testAgeMois_VerifyGetDateCalled()`  
**Calcul attendu**: Naissance 2022-01-01 → 1 mois  
**Exercice**: 4.3.3

**Spécificité**: Ce test utilise un `mock` au lieu du stub réel pour vérifier que `getDate()` est appelé exactement une fois.

---

## 📊 Tableau de synthèse

| Stub | Cas | Date | Naissance | Âge | Exercice | Cas limite |
|------|-----|------|-----------|-----|----------|-----------|
| DateConsoleStub | CT_age2 | 2022-01-01 | 2021-02-28 | 10 | 4.2.1 | - |
| DateConsoleStub2 | CT_age3 | 2022-01-01 | 2021-12-31 | 0 | 4.2.1 | Juste après naissance |
| DateConsoleStub3 | CT_age4 | 2022-01-01 | 2021-12-01 | 1 | 4.2.1 | Mois exact |
| DateConsoleStub4 | CT_age5 | 2022-02-28 | 2021-02-15 | 12 | 4.2.2 | Non-bissextile |
| DateConsoleStub5 | CT_age6 | 2022-02-15 | 2021-02-15 | 12 | 4.2.2 | Jour anniversaire |
| DateConsoleStub6 | CT_age7 | 2022-02-01 | 2021-02-15 | 11 | 4.2.2 | Avant anniversaire |
| DateConsoleStub7 | CT_age7bis | 2022-02-01 | 2022-01-01 | 1 | 4.3.1 | Mois exact |
| DateConsoleStub8 | Spy | 2022-02-01 | 2022-01-01 | 1 | 4.3.3 | Avec mock |

---

## 🎯 Patterns utilisés

### Pattern Stub
Chaque classe retourne une `LocalDate` fixe, ce qui en fait un "stub" - une implémentation simple pour tester un comportement spécifique.

```kotlin
override fun getDate(): LocalDate = LocalDate.of(YYYY, M, D)
```

### Avantages
- ✅ Simple et lisible
- ✅ Une responsabilité par classe
- ✅ Facile à tester
- ✅ Pas de dépendance externe
- ✅ Déterministe

### Alternative avec Mock
Pour le test de spy (4.3.3), un mock est utilisé au lieu d'un stub pour pouvoir vérifier les interactions:

```kotlin
val mockDateProvider = mockk<DateProvider>()
every { mockDateProvider.getDate() } returns LocalDate.of(2022, 2, 1)
verify(exactly = 1) { mockDateProvider.getDate() }
```

---

## 🔄 Flux d'utilisation dans les tests

```
┌─────────────────────────┐
│  Test (ex: CT_age2)     │
└────────────┬────────────┘
             │
             ├─ crée une instance DateConsoleStub()
             │
             ├─ passe au constructeur Chien("Dog", "Race", stub)
             │
             ├─ appelle ch.setDateNaissance(...)
             │
             ├─ appelle ch.ageMois()
             │  ├─ demande la date à dateProvider
             │  ├─ DateConsoleStub retourne 2022-01-01
             │  └─ calcule l'écart
             │
             └─ vérifie le résultat avec assertEquals()
```

---

## 📝 Conventions de nommage

Les stubs suivent une convention de nommage claire:
- `DateConsoleStub` → CT_age2 (cas 2)
- `DateConsoleStub2` → CT_age3 (cas 3)
- `DateConsoleStub3` → CT_age4 (cas 4)
- etc.

Cette convention facilite la navigation et la compréhension des tests.

---

## ✨ Résumé

- **8 stubs créés** pour les 8 cas de test différents
- **Chacun retourne une date fixe** pour un contrôle total
- **Simples et maintenables** - une responsabilité par classe
- **Respectent SOLID** - Single Responsibility Principle
- **Testables et réutilisables** pour d'autres scénarios

