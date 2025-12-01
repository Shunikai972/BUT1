package decroissant

/*
La fonction decroissant doit trier un tableau d'entiers dans l'ordre décroissant.

# Entrée
- t, le tableau à trier (en place)

# Info
2022-2023, test 4, exercice 7
2024-2025, test 3, exercice 1
*/

func decroissant(tab []int) {
	n := len (tab)
	for i := 0; i < n; i++ {
		index_max := i
		for a := i + 1; a < n; a++ {
			if tab[a] > tab[index_max] {
				index_max = a
			}
		}
		tab[i], tab[index_max] = tab[index_max], tab[i]
	}

}
  