package ateddybearpicnic

/*
On considère le jeu suivant avec des ours en peluche.

Au début du jeu vous recevez un certain nombre d'ours. Vous pouvez rendre des ours, mais seulement selon les règles suivantes. Le but du jeu est de réussir à garder exactement 42 ours.

Les règles pour rendre les ours sont :
- Si vous avez un nombre pair d'ours, vous pouvez en rendre exactement la moitié.
- Si vous avez un nombre d'ours divisible par 3 ou 4 vous pouvez multiplier les deux derniers chiffres de ce nombre d'ours et en rendre cette quantité. Par exemple, si vous avez 152 ours (divisible par 5) vous pouvez en rendre 5 * 2 = 10.
- Si vous avez un nombre d'ours divisible par 5 vous pouvez en rendre 42.

La fonction winnable doit dire si, pour un nombre d'ours reçu en début de partie, il est possible de gagner ou pas.

# Entrée
- numReceived : le nombre d'ours reçus en début de partie

# Sortie
- winnable : un booléen qui vaut true si il est possible de gagner la partie en ayant reçu numReceived ours et false si il n'est pas possible de gagner la partie dans cette situation

# Info
- Cet exercice est issu de https://home.cs.colorado.edu/~main/projects/Assn0901-FunWithRecursion.html
- 2023-2024, test 1, exercice 9
*/

func bears(numReceived uint) (winnable bool) {
	return
}
