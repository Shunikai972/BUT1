package suite

/*
On considère la suite U telle que
- U(0) = 1 et
- U(n) = 5 * U(n-1) - 2
La fonction u calcule les termes de cette suite.

# Entrée
- n : le numéro du terme à calculer

# Sortie
- un : la valeur de U(n)

# Info
2024-2025, test 1, exercice 1
*/

func u(n int) (un int) {
	if n == 0 {
		return 1
	}

	return 5*u(n-1) - 2
}
