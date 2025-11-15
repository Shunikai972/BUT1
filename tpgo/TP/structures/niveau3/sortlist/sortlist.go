package sortlist

/*
On dispose d'une structure de liste chaînée. Étant donnée une liste l on souhaite obtenir une nouvelle liste sortedL contenant les mêmes valeurs mais triées en ordre croissant. Ceci doit être fait sans modifier l. C'est le travail de la méthode Sort.

# Entrée
- l : une liste de valeurs

# Sortie
- sortedL : une liste contenant les mêmes valeurs que l mais triées en ordre croissant

# Info
2023-2024, test 3, exercice 8
*/

type List struct {
	Content *Element
}

type Element struct {
	V    int
	Next *Element
}

func (l List) Sort() (sortedL List) {
	return
}
