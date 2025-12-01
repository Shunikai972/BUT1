package triStruct

/*
On souhaite trie un ensemble de rectangle en fonction de leur surface : de la plus petite surface à la plus grande surface. Les rectangles sont représentés par une structure rectangle qui comprend un champs largeur et un champs longueur (on rappelle que la surface d'un rectangle est le produit de sa largeur par sa longueur). C'est la fonction ranger qui doit effectuer ce tri.

# Entrée
- t : un tableau de rectangles

# Sortie
- enOrdre : un tableau de rectangles ayant exactement les mêmes éléments que t mais rangés de celui qui a la plus petite surface à celui qui a la plus grande surface. L'ordre relatif des rectangles ayant la même surface n'a pas d'importance.

# Info
2024-2025, test 2, exercice 4
*/

type rectangle struct {
	largeur  int
	longueur int
}

func ranger(t []rectangle) (enOrdre []rectangle) {
	n := len(t)
	enOrdre = make([]rectangle, n)
	copy(enOrdre, t)
	
	for i := 0; i < n; i++ {
		index_min := i
		area_i := enOrdre[i].largeur * enOrdre[i].longueur
		for a := i + 1; a < n; a++ {
			area_a := enOrdre[a].largeur * enOrdre[a].longueur
			if area_a < area_i {
				index_min = a
				area_i = area_a
			}
		}
		enOrdre[i], enOrdre[index_min] = enOrdre[index_min], enOrdre[i]
	}

	return enOrdre
}
