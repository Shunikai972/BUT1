package majeur

/*
On dispose d'une structure pour représenter des personnes. Chaque personne a un prénom, un nom, et un age en années. De plus, on veut indiquer si la personne est majeure (elle à au moins 18 ans) ou non. On souhaite écrire une méthode verifieMajeur qui positionne le champs estMajeur de notre structure à true si la personne est majeure et à false sinon. Aucun autre champs de la structure ne doit être modifié par verifieMajeur.

# Entrée
- p : une structure représentant une personne (c'est cette structure qui va être modifiée par la méthode, il n'y a donc pas de sorties)

# Info
2023-2024, test 2, exercice 4
*/

type personne struct {
	prenom    string
	nom       string
	age       int
	estMajeur bool
}

func (p *personne) verifieMajeur() {
	if p.age >= 18 {
		p.estMajeur = true
	} else {
		p.estMajeur = false
	}
}
