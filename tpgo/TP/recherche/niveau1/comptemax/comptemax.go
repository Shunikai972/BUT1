package comptemax

/*
La fonction compteMax doit trouver le plus grand entier dans un tableau et indiquer combien de fois cet entier apparaît dans le tableau.

# Entrée
- t : un tableau d'entiers

# Sortie
- val : le plus grand entier dans t
- nombre : le nombre de fois que val apparaît dans t

# Info
- 2023-2024, test 1, exercice 7
*/

func compteMax(t []int) (val, nombre int) {
	if len(t) == 0 {
		return 0, 0
	}

	val = t[0]
	nombre = 1

	for i := 1; i < len(t); i++ {
		if t[i] > val {
			val = t[i]
			nombre = 1
		} else if t[i] == val {
			nombre++
		}
	}

	return val, nombre
}
