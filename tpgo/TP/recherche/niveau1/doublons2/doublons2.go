package doublons2

/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient
exactement une fois chaque nombre compris entre 1 et n. On voudrait vérifier
cela. C'est le travail de la fonction doublons.

Coder la fonction doublons de manière à ne parcourir le tableau tab qu'une seule
fois rapportera plus de points.

# Entrée
- tab : un tableau d'entiers

# Sortie
- ok : un booléen qui vaut true si tab contient bien exactement une
fois chaque entier entre 1 et len(tab) et false sinon

# Info
2021-2022, test 1, exercice 8
*/

func doublons(tab []int) (ok bool) {
	var present []bool = make([]bool, len(tab))
	for i := 0; i < len(tab); i++ {
		if !(tab[i] >= 1 && tab[i] <= len(tab)) {
			return false
		}
		if present[tab[i]-1] {
			return false
		}
		present[tab[i]-1] = true
	}
	return true
}
