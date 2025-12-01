package tritri

/*
On a un tableau de tableaux d'entiers. On souhaite trier les entiers de chacun de ces tableaux en ordre croissant. De plus, on souhaite trier le tableau de tableaux en s'assurant qu'un tableau t1 est avant un tableau t2 si et seulement si
- le plus petit entier de t1 est plus petit que le plus petit entier de t2, ou
- les n plus petits entiers de t1 sont égaux aux n plus petits entiers de t2 et le n+1-ième entier de t1 est plus petit que le n+1-ième entier de t2, ou
- tous les entiers de t1 correspondent aux plus petits entiers de t2 et len(t1) <= len(t2).
C'est le rôle de la fonction tritri. Tous les tris doivent être faits en place.

# Entrée
- tab : un tableau de tableau d'entiers

# Info
2024-2025, test 3, exercice 8
*/

func tritri(tab [][]int) {
	for i := 0; i < len(tab); i++ {
		// Tri des sous-tableaux en ordre croissant
		n := len(tab[i])
		for a := 0; a < n; a++ {
			index_min := a
			for b := a + 1; b < n; b++ {
				if tab[i][b] < tab[i][index_min] {
					index_min = b
				}
			}
			tab[i][a], tab[i][index_min] = tab[i][index_min], tab[i][a]
		}
	}

	// Tri des tableaux selon les critères donnés
	for i := 0; i < len(tab); i++ {
		for j := i + 1; j < len(tab); j++ {
			k := 0
			for k < len(tab[i]) && k < len(tab[j]) && tab[i][k] == tab[j][k] {
				k++
			}
			if (k < len(tab[i]) && k < len(tab[j]) && tab[i][k] > tab[j][k]) || (k == len(tab[j]) && len(tab[i]) > len(tab[j])) {
				tab[i], tab[j] = tab[j], tab[i]
			}
		}
	}
}
