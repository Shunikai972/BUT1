package comptevrai

/*
La fonction compteVrai doit indiquer combien de fois la valeur true apparaît dans un tableau de booléens.

# Entrée
- t : un tableau de booléens

# Sortie
- nombre : le nombre de fois que la valeur true apparaît dans t

# Info
2023-2024, test 1, exercice 2
*/

func compteVrai(t []bool) (nombre int) {

	for i := 0; i < len(t); i++ {
		if t[i] {
			nombre++
		}
	}
	return (nombre)
}
