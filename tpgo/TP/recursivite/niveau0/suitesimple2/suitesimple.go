package suitesimple

/*
Soit la suite (Un) définie de la façon suivante :
- U(n) = -5 * U(n-1) + 5
- U(0) = 0
La fonction terme doit calculer le terme numéro n de (Un).

Pour cet exercice, les boucles for sont interdites.

# Entrée
- n, un entier positif ou nul

# Sortie
- un, le terme U(n) de la suite

# Info
2023-2024, test 1, exercice 1
*/

func terme(n uint) (un int) {

	if(n == 0){
		return 0
	}




	return -5 * terme(n-1)+5
}
