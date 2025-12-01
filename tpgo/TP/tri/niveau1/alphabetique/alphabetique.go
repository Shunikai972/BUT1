package alphabetique

/*
On souhaite classer un ensemble de mots par ordre alphabétique. Les mots sont donnés dans un tableau et doivent être triés en place dans ce même tableau. C'est le rôle de la fonction alphabetique.

On rappel que la comparaison "<" peut être utilisées sur des chaînes de caractères pour savoir la quelle est la première en ordre alphabétique.

# Entrée
- dico : un tableau de chaînes de caractères qui doit être trié en place en ordre alphabétique

# Info
2023-2024, test 3, exercice 2
*/

func alphabetique(dico []string) {
    n := len(dico)
    
    for i := 0; i < n; i++ {
        index_min := i
        
        for a := i + 1; a < n; a++ {
            if dico[a] < dico[index_min] {
                index_min = a
            }
        }
        
        dico[i], dico[index_min] = dico[index_min], dico[i]
    }
}
/*
func stortBymark (t []student)  {
	for i:= 0; i < len(t); i++{
		index_max := i
		for a := i +1; a < len(t); a++{
			if t[index_max].mark < t[a].mark {
				index_max = a
			}
		}
		t[index_max], t[i]= t[i], t[index max]
	}
}*/