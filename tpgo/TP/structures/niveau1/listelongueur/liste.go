package liste

/*
On considère des listes chainées, telles que vues en cours. Elles sont codées ici par la structure listeChainee. On souhaite savoir combien il y a d'éléments dans une liste chainée. C'est le rôle de la fonction numElements.

# Entrée
- l : une liste d'éléments (bien formée : elle ne contient pas de cycles)

# Sortie
- num : le nombre d'éléments dans la liste l

# Info
2024-2025, test 3, exercice 6
*/

type listeChainee struct {
	head *element
}

type element struct {
	v    int
	next *element
}

func numElements(l listeChainee) (num int) {
	current := l.head
	for current != nil {
		num ++
		current = current.next
	}
	return num
}
