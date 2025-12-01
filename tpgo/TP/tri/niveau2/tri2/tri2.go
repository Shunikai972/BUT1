package tri2

/*
La fonction triabs doit trier un tableau d'entiers de la plus grande valeure
absolue à la plus petite valeure absolue. Cette fonction ne doit pas modifier
le tableau donné en entrée.

# Entrée
- tinit : un tableau d'entiers qui ne doit pas être modifié.

# Sortie
- tfin : un tableau contenant les mêmes entiers que tinit mais triés du plus
         grand (en valeure absolue) au plus petit (en valeure absolue).

# Info
2021-2022, test2, exercice 7
*/

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func triabs(tinit []int) (tfin []int) {
	tfin = make([]int, len(tinit))
	copy(tfin, tinit)

	for i := 0; i < len(tfin); i++ {
		index_max := i
		for j := i + 1; j < len(tfin); j++ {
			if abs(tfin[j]) > abs(tfin[index_max]) {
				index_max = j
			}
		}
		tfin[i], tfin[index_max] = tfin[index_max], tfin[i]
	}
	return tfin
}


