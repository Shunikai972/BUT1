package doublons1

/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient
exactement les nombres de 1 à n dans l'ordre. On voudrait vérifier cela. C'est
le travail de la fonction doublons.

# Entrée
- tab : un tableau d'entiers

# Sortie
- ok : un booléen qui doit valoir true si tab contient exactement les entiers de
1 à len(tab) dans l'ordre et false sinon
*/

func doublons(tab []int) bool {
	for i := 0; i < len(tab); i++ {
		if tab[i] != i+1 {
			return false
		}
	}
	return true
}
