package tri

/*
On dispose d'un ensemble de rectangle, chacun portant un nom (représentés par une structure "rectangle"). On souhaite trier ces rectangles de la manière suivante :
- de la plus petite surface à la plus grande surface (on rappelle que la surface d'un rectangle est le produit de sa largeur par sa longueur)
- en cas d'égalité de surface, du nom le plus court au nom le plus long
- en cas d'égalité de surface et de longueur de nom, en ordre alphabétique des noms
C'est le rôle de la fonction ranger de réaliser ce tri.

# Entrée
- t : un tableau de rectangles

# Sortie
- enOrdre : un tableau contenant les mêmes rectangles mais trié selon la méthode décrite ci-dessus

# Info
2024-2025, test 2, exercice 7
*/

type rectangle struct {
	nom      string
	largeur  int
	longueur int
}

func ranger(t []rectangle) (enOrdre []rectangle) {
	n := len(t)
	enOrdre = make([]rectangle, n)
	copy(enOrdre, t)

	for i := 0; i < n; i++{
		index_min := i
		aire_i := enOrdre[i].largeur * enOrdre[i].longueur
		for a := i + 1; a < n; a++{	
			aire_a := enOrdre[a].largeur * enOrdre[a].longueur
			if aire_a < aire_i || (aire_a == aire_i && (len(enOrdre[a].nom) < len(enOrdre[index_min].nom) || (len(enOrdre[a].nom) == len(enOrdre[index_min].nom) && enOrdre[a].nom < enOrdre[index_min].nom))) {
				index_min = a
				aire_i = aire_a
			}
		}
		enOrdre[i], enOrdre[index_min] = enOrdre[index_min], enOrdre[i]
	}
	return
}
