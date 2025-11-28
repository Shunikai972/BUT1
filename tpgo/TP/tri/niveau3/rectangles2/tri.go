package tri

/*
On dispose d'un ensemble de rectangle, chacun portant un nom (représentés par une structure "rectangle"). On souhaite trier ces rectangles de la manière suivante :
- de la plus petite surface à la plus grande surface (on rappelle que la surface d'un rectangle est le produit de sa largeur par sa longueur)
- en cas d'égalité de surface, du nom le plus court au nom le plus long
- en cas d'égalité de surface et de longueur de nom, en ordre alphabétique des noms
C'est le rôle de la fonction ranger de réaliser ce tri.

# Entrée
- t : un tableau de rectangles

# Sortie
- enOrdre : un tableau contenant les mêmes rectangles mais trié selon la méthode décrite ci-dessus

# Info
2024-2025, test 2, exercice 7
*/

type rectangle struct {
	nom      string
	largeur  int
	longueur int
}

func ranger(t []rectangle) (enOrdre []rectangle) {
	return
}
