package contient

/*
On dispose d'une structure de liste chaînée contenant des éléments avec chacun une valeur. On souhaite, étant donnée une liste, savoir si elle contient une valeur particulière (savoir si un de ses éléments a cette valeur) grâce à la méthode contient. Cette méthode ne doit pas modifier la liste.

On considère qu'il n'y a jamais de cycles dans nos listes : en allant de suivant en suivant on arrive obligatoirement au bout d'un moment sur un dernier élément (dont le suivant est nil).

Pour cet exercice les boucles for sont interdites.

# Entrées
- l : une liste
- v : un entier

# Sorties
- existe : un booléen qui vaut true si un des éléments de l a pour valeur v et qui vaut false sinon

# Info
2023-2024, test 2, exercice 5
2024-2025, test 3, exercice 7
*/

type liste struct {
	debut *element
}

type element struct {
	valeur  int
	suivant *element
}

func (l liste) contient(v int) (existe bool) {
	var contient_aux()
	if e == nil{
		return false
	}
	if 	e.valeur == v { 
		return true
	} else{
		return contient(e.next, v)
	}
	return
}
