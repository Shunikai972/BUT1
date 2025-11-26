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
package sortlist

type List struct {
	Content *Element
}

type Element struct {
	V    int
	Next *Element
}

func (l List) Sort() (sortedL List) {
	var values []int
	for e := l.Content; e != nil; e = e.Next {
		values = append(values, e.V)
	}

	for i := 0; i < len(values); i++ {
		min := i
		for j := i + 1; j < len(values); j++ {
			if values[j] < values[min] {
				min = j
			}
		}
		values[i], values[min] = values[min], values[i]
	}

	var head, tail *Element
	for _, v := range values {
		elem := &Element{V: v}
		if head == nil {
			head = elem
		} else {
			tail.Next = elem
		}
		tail = elem
	}

	sortedL.Content = head
	return
}
