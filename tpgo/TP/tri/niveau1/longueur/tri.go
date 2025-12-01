package tri

/*
On souhaite trier un ensemble de chaînes de caractères selon leur longueur : la plus courte en premier, la plus longue en dernier. C'est le rôle de la fonction classer.

# Entrée
- t : un tableau de chaînes de caractères

# Sortie
- enOrdre : un tableau de chaînes de caractères qui contient les mêmes chaînes que t mais triées en ordre croissant de leur longueur. Si deux chaînes ont la même longueur leur ordre relatif n'a pas d'importance.

# Info
2024-2025, test 2, exercice 3
*/

func classer(t []string) (enOrdre []string) {
	n := len(t) 
	enOrdre = make([]string, n)
	copy(enOrdre, t)

	for i := 0; i < n; i++ {
		index_min := i 
		for a := i + 1; a < n; a++{
			if len(enOrdre[a]) < len(enOrdre[index_min]) {
				index_min = a
			}
		}
		enOrdre[i], enOrdre[index_min] = enOrdre[index_min], enOrdre[i]
	}

	return enOrdre}
