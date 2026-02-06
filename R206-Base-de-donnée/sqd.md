### Question 1: Liste des communes qui se trouvent dans le département 49 (in et exists)

**Forme IN :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Commune, Departement | - | NOM_COMMUNE | Voir Annexe A | 0.004 secondes |

**Forme EXISTS :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Commune, Departement | - | NOM_COMMUNE | Voir Annexe A | 0.003 secondes |

### Question 2: Liste des noms des communes avec le nom du département qui disposent de la 5G (in et exists)

**Forme IN :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Commune, Distribution, Operateur | Distribution sur Commune (CODE_INSEE), Operateur sur Distribution (NUMFO) | NOM_COMMUNE, NOMDEP | Voir Annexe A | 0.061 secondes |

**Forme EXISTS :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Commune, Distribution, Operateur | Distribution sur Commune (CODE_INSEE), Operateur sur Distribution (NUMFO) | NOM_COMMUNE, NOMDEP | Voir Annexe A | 0.045 secondes |

### Question 3: Liste des communes avec le nom du département qui ne possèdent pas la 5G (not in et not exists)

**Forme NOT IN :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Commune, Distribution, Operateur | Distribution sur Commune (CODE_INSEE), Operateur sur Distribution (NUMFO) | NOM_COMMUNE, NOMDEP | Voir Annexe C | 0.014 secondes |

**Forme NOT EXISTS :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Commune, Distribution, Operateur | Distribution sur Commune (CODE_INSEE), Operateur sur Distribution (NUMFO) | NOM_COMMUNE, NOMDEP | Voir Annexe C | 0.025 secondes |

### Question 4: Liste des communes qui ne possèdent pas la 5G et qui se trouvent dans le département 44 (not in et not exists)

**Forme NOT IN :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Commune, Departement, Distribution, Operateur | Departement sur Commune (NOMDEP), Distribution sur Commune (CODE_INSEE), Operateur sur Distribution (NUMFO) | NOM_COMMUNE | Voir Annexe C | 0.021 secondes |

**Forme NOT EXISTS :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Commune, Departement, Distribution, Operateur | Departement sur Commune (NOMDEP), Distribution sur Commune (CODE_INSEE), Operateur sur Distribution (NUMFO) | NOM_COMMUNE | Voir Annexe C | 0.025 secondes |

### Question 5: Afficher pour chaque département le nombre de communes (group by et sous requête corrélée)

**Forme GROUP BY :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Departement, Commune | Departement sur Commune (NOMDEP) | NOMDEP, COUNT(CODE_INSEE) | Voir Annexe D | 0.007 secondes |

**Forme Sous requête corrélée :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Departement, Commune | - | NOMDEP, NB_COMMUNES | Voir Annexe D | 0.005 secondes |

### Question 6: Liste de toutes les communes de Loire-Atlantique avec le nombre d'antennes 5G (inner join et sous requête select)

**Forme INNER JOIN :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Commune, Distribution, Operateur, Departement | Distribution sur Commune (CODE_INSEE), Operateur sur Distribution (NUMFO), Departement sur Commune (NOMDEP) | NOM_COMMUNE, NB_5G | Voir Annexe D | 0.031 secondes |

**Forme Sous requête select :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Commune, Departement, Distribution, Operateur | Departement sur Commune (NOMDEP), Distribution sur Commune (CODE_INSEE), Operateur sur Distribution (NUMFO) | NOM_COMMUNE, NB_5G | Voir Annexe D | 0.253 secondes |

### Question 7: Idem que la question précédente avec au moins 10 antennes (group by/having et sous requête corrélée dans le where)

**Forme GROUP BY/HAVING :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Commune, Distribution, Operateur, Departement | Distribution sur Commune (CODE_INSEE), Operateur sur Distribution (NUMFO), Departement sur Commune (NOMDEP) | NOM_COMMUNE, NB_5G | Voir Annexe E | 0.026 secondes |

**Forme Sous requête corrélée dans le WHERE :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Commune, Departement, Distribution, Operateur | Departement sur Commune (NOMDEP), Distribution sur Commune (CODE_INSEE), Operateur sur Distribution (NUMFO) | NOM_COMMUNE, NB_5G | Voir Annexe E | 0.279 secondes |

### Question 8: Liste de toutes les communes de Loire-Atlantique avec le nombre d'antennes 5G et 4G (select)

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Commune, Departement, Distribution, Operateur | Departement sur Commune (NOMDEP), Distribution sur Commune (CODE_INSEE), Operateur sur Distribution (NUMFO) | NOM_COMMUNE, NB_5G, NB_4G | - | 0.86 secondes |

### Question 9: Liste de tous les fournisseurs avec le nombre d'antennes 5G de 3,5Ghz, 2,1 GHZ et 700MHZ (group by et sous requête corrélée annexe f)

**Forme GROUP BY :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Operateur, Distribution | Distribution sur Operateur (NUMFO) | NOMFO, TECHNOLOGIE, NB_ANTENNES | Voir Annexe F | 0.063 secondes |

**Forme Sous requête corrélée annexe F :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Operateur, Distribution | Distribution sur Operateur (NUMFO) | NOMFO, TECHNOLOGIE, NB_ANTENNES | Voir Annexe F | 0.021 secondes |

### Question 10: Idem que la question précédente pour la ville de Nantes (group by et sous requêtes corrélée)

**Forme GROUP BY :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Operateur, Distribution, Commune | Distribution sur Operateur (NUMFO), Commune sur Distribution (CODE_INSEE) | NOMFO, TECHNOLOGIE, NB_ANTENNES | Voir Annexe F | 0.035 secondes |

**Forme Sous requêtes corrélée :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Operateur, Distribution, Commune | Distribution sur Operateur (NUMFO), Commune sur Distribution (CODE_INSEE) | NOMFO, TECHNOLOGIE, NB_ANTENNES | Voir Annexe F | 0.025 secondes |

### Question 11: Idem pour le département de Loire-Atlantique (group by et sous requêtes corrélée)

**Forme GROUP BY :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Operateur, Distribution, Commune, Departement | Distribution sur Operateur (NUMFO), Commune sur Distribution (CODE_INSEE), Departement sur Commune (NOMDEP) | NOMFO, TECHNOLOGIE, NB_ANTENNES | Voir Annexe F | 0.132 secondes |

**Forme Sous requêtes corrélée :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Operateur, Distribution, Commune, Departement | Distribution sur Operateur (NUMFO), Commune sur Distribution (CODE_INSEE), Departement sur Commune (NOMDEP) | NOMFO, TECHNOLOGIE, NB_ANTENNES | Voir Annexe F | 0.023 secondes |

### Question 12: Listes des communes qui possèdent le déploiement de toutes les technologies des différents opérateurs (not exist et not in)

**Forme NOT EXIST :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Commune, Operateur, Distribution | Operateur sur Distribution (NUMFO), Commune sur Distribution (CODE_INSEE) | NOM_COMMUNE | Voir Annexe C | 1.763 secondes |

**Forme NOT IN :**

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Commune, Operateur, Distribution | Operateur sur Distribution (NUMFO), Commune sur Distribution (CODE_INSEE) | NOM_COMMUNE | Voir Annexe C | 1.992 secondes |

### Question 13: Listes des communes qui possèdent le déploiement des technologies 5G des différents opérateurs (select)

| Tables | Jointures | Sortie | Type de transformation utilisé | Temps d'exécution |
| - | - | - | - | - |
| Commune, Operateur, Distribution | Operateur sur Distribution (NUMFO), Commune sur Distribution (CODE_INSEE) | NOM_COMMUNE | - | 1.625 secondes |

