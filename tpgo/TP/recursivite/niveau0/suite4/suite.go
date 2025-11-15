package suite

/*
On considère la suite (Un) définie de la façon suivante :
- U(0) = 5
- U(1) = 4
- U(n) = 2 * U(n-2) - 3
La fonction u doit calculer les termes de cette suite.

Attention, pour cet exercice les boucles for sont interdites.

# Entrée
- n : un entier

# Sortie
- un : le terme n de la suite (Un), c'est-à-dire U(n)

# Info
2024-2025, test 2, exercice 0
*/

func u(n int) (un int) {
	if n == 0 {
		return 5
	}
	if n == 1 {
		return 4
	}

	return 2*u(n-2) - 3
}
