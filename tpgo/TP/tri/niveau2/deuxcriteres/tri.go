package tri

/*
La fonction tri doit trier un tableau d'entiers de la manière suivante : on trouvera d'abord tous les nombres strictement négatifs du tableau, dans l'ordre décroissant, puis tous les nombres positifs ou nuls, dans l'ordre croissant.

# Entrée
- t, le tableau à trier (en place)

# Info
2022-2023, test 4, exercice 8
*/

func tri(t []int) {
	for i := 0; i < len(t); i++ {
		for j := i + 1; j < len(t); j++ {
			if (t[i] < 0 && t[j] < 0 && t[i] < t[j]) || (t[i] >= 0 && t[j] >= 0 && t[i] > t[j]) || (t[i] >= 0 && t[j] < 0) {
				t[i], t[j] = t[j], t[i]
			}
		}
	}
}
