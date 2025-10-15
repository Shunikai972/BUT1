package doublons6

/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient exactement les nombres de k à k + n - 1 dans func doublons(tab []int) bool {
	for i := 0; i < len(tab); i++ {
		if tab[i] != i+1 {
			return false
		}
	}
	return true
}
doublons.

# Entrée
- tab : un tableau d'entiers

# Sortie
- ok : un booléen qui doit valoir true si tab contient exactement les entiers de k à k + len(tab) - 1 dans l'ordre et qui vaut false sinon

# Info
- 2023-2024, test 2, exercice 0
*/

func doublons(tab []int) bool {
	var k int = tab[0]
	for i := 0; i < len(tab); i++ {
		if tab[i] != k+i {
			return false
		}
	}
	return true
}
