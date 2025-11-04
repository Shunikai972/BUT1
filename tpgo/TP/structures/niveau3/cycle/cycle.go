package cycle

/*
On considère une structure de liste chaînée constituée de nœuds. On veut vérifier qu'une liste est bien formée, c'est-à-dire qu'elle ne contient aucun cycle. En d'autres termes, on veut vérifier que si on part du premier nœud et qu'on progresse de suivant en suivant on arrive toujours au bout d'un moment à un suivant qui vaut nil (donc on ne retombe jamais sur un nœud qu'on a déjà vu). La fonction estCyclique sert à vérifier cela.

# Entrée
- l : une liste

# Sortie
- contientCycle : un booléen qui vaut "true" si la liste l contient un cycle et "false" sinon.

# Info
2024-2025, test 2, exercice 8
*/

type liste struct {
	tete *noeud
}

type noeud struct {
	valeur  int
	suivant *noeud
}

func estCyclique(l liste) (contientCycle bool) {
	return
}
