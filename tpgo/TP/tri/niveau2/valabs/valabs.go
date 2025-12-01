package valabs

/*
La fonction valabs doit trier un tableau d'entiers de la plus petite valeur absolue
à la plus grande valeur absolue. En cas
d'égalité de valeur absolue, les nombres
négatifs doivent être placés avant les
nombres positifs.

# Entrée
- tab : un tableau d'entiers
*/
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func valabs(tab []int) {
	for i := 0; i < len(tab); i++ {
		index_min := i
		for j := i + 1; j < len(tab); j++ {
			if(abs(tab[j]) < abs(tab[index_min])) || (abs(tab[j]) == abs(tab[index_min]) && tab[j] < tab[index_min]) {
				index_min = j
			}
		}
		tab[i], tab[index_min] = tab[index_min], tab[i]
	}
}
