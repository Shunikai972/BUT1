package lirepointeur

/*
La fonction lire doit retourner l'entier qui est stocké à l'adresse donnée par un pointeur. Elle ne doit pas modifier la valeur stockée à cette adresse.

# Entrée
- ptrx : l'adresse d'un entier

# Sortie
- x : la valeur de l'entier situé à l'adresse ptrx

# Info
2024-2025, test 3, exercice 0
*/

func lire(ptrx *int) (x int) {
	x = *ptrx
	return x
}
