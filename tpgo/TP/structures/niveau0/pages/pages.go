package pages

/*
On dispose d'une structure de données permettant de réprésenter des livres avec un titre, un prix et un nombre de page.

On souhaite trouver dans un tableau de livres un livre qui a un nombre de pages donné. C'est ce que doit faire la fonction trouver.

# Entrée
- tab : le tableau de livres
- numPages : le nombre de pages pour lequel on souhaite trouver un livre

# Sortie
- l : un livre provenant de tab et qui fait numPages pages, s'il existe (sinon, un livre quelconque)
- trouve : un booléen qui vaut true si on a trouvé un livre de numPages pages dans tab et false sinon

# Info
2022-2023, test3, exercice 3
*/

type livre struct {
	titre    string
	prix     float64
	numPages int
}

func trouver(tab []livre, numPages int) (l livre, trouve bool) {
	for _, livre := range tab {
		if numPages == livre.numPages{
			return livre,true
		}
			

	}
 	return livre{}, false
		
}
