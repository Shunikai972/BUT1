package recherche

/*
On souhaite savoir si la valeur "true" est présente au moins une fois dans un tableau de booléens. C'est le rôle de la fonction trouveVrai.

# Entrée
- t : un tableau de booléens

# Sorties
- pos : la première position à laquelle on peut trouver la valeur "true" dans t (si elle est présente)
- existe : un booléen qui vaut "true" si la valeur "true" est présente dans t et "false" sinon.

# Info
2024-2025, test 2, exercice 1
*/

func trouveVrai(t []bool) (pos int, existe bool) {
	for i := 0; i < len(t); i++ {
		if t[i] {
			return i, true
		}
	}
	return -1, false
}
