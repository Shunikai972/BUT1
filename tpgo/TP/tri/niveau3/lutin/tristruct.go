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

func tri(tab []chocolat) {}
