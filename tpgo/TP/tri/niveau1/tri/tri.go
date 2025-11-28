package tri

/*
La fonction tri doit trier un tableau d'entiers du plus petit au plus grand.
Cette fonction ne doit pas modifier le tableau donné en entrée.

# Entrée
- tinit : un tableau d'entiers qui ne doit pas être modifié.

# Sortie
- tfin : un tableau contenant les mêmes entiers que tinit mais triés du plus
         petit au plus grand.

# Info
2021-2022, test2, exercice 6
*/

func tri(tinit []int) (tfin []int) {
    tfin = make([]int, len(tinit))
    copy(tfin, tinit)

    for i := 0; i < len(tfin); i++ {
        index_min := i
        for j := i + 1; j < len(tfin); j++ {
            if tfin[j] < tfin[index_min] {
                index_min = j
            }
        }
        tfin[i], tfin[index_min] = tfin[index_min], tfin[i]
    }

    return tfin
}