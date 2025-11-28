package classement

/*
La fonction classer doit trier un tableau de booléen en séparant d'un côté les booléens qui valent true et de l'autre les booléens qui valent false

# Entrée
- t : le tableau de booléens à trier

# Sortie
- tri : un tableau de booléens qui contient exactement les mêmes éléments que t mais avec tous les booléens qui valent true au début et tous les booléens qui valent false à la fin

# Info
- 2023-2024, test 1, exercice 4
*/

package classement

func classer(t []bool) (tri []bool) {
	for i := range t {          // parcours avec i
		if t[i] {
			tri = append([]bool{t[i]}, tri...) // true au début
		} else {
			tri = append(tri, t[i])           // false à la fin
		}
	}
	return
}
