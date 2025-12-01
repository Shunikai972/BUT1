package tri

/*
La fonction tri doit trier un tableau d'entiers de la manière suivante : on trouvera d'abord tous les nombres pairs du tableau, dans l'ordre croissant, puis tous les nombres impairs, dans l'ordre décroissant.

# Entrée
- t, le tableau à trier (en place)

# Info
2022-2023, test 2, exercice 4
*/

func tri(t []int) {
	for i := 0; i < len(t); i++ {
		for j := i + 1; j < len(t); j++ {
			if (t[i]%2 == 0 && t[j]%2 == 0 && t[i] > t[j]) || (t[i]%2 != 0 && t[j]%2 != 0 && t[i] < t[j]) || (t[i]%2 != 0 && t[j]%2 == 0) {
				t[i], t[j] = t[j], t[i]
			}
		}
	}
	
}