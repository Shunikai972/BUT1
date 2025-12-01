package tristruct

/*
Un lutin maniaque veut trier des chocolats. Il souhaitent que ceux au praliné (praline vaut true) soient avant ceux sans praline (praline vaut false). Dans chacune de ces deux catégories il veut aussi que ceux aux noisettes (noisette vaut true) soient avant ceux sans noisettes (noisette vaut false). Enfin, dans chacune des quatre catégories (avec praliné et noisettes, avec praliné sans noisettes, sans praliné avec noisettes, sans praliné sans noisettes) il veut trier les chocolats du plus lourd au plus léger (ça n'a pas vraiment d'importance, mais les poids sont donnés un grammes).

La fonction tri doit ranger ces chocolats de la manière désirée par le lutin. Ce tri doit se faire en place.

# Entrée
- tab : un tableau de chocolats

# Info
2024-2025, test 3, exercice 4
*/

type chocolat struct {
	poids    int
	praline  bool
	noisette bool
}

func tri_aux(c1, c2 chocolat) bool {
	if c1.praline != c2.praline {
		return c1.praline // les pralinés avant les non-pralinés
	}
	if c1.noisette != c2.noisette {
		return c1.noisette // les noisettes avant les non-noisettes
	}
	return c1.poids > c2.poids // plus lourd avant plus léger
}

// Tri en place (tri à bulles pour l'exemple)
func tri(tab []chocolat) {
	n := len(tab)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if !tri_aux(tab[j], tab[j+1]) {
				tab[j], tab[j+1] = tab[j+1], tab[j] // échange
			}
		}
	}
}