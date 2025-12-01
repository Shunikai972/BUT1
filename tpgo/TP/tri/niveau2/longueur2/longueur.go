package longueur

/*
On souhaite trier un tableau de chaînes de caractères (minuscules sans accents) de la plus longue à la plus courte (et par ordre alphabétique dans le cas où plusieurs chaînes ont la même longueur). La fonction trier doit réaliser ce tri en place. On rappelle que l'opérateur < sur les chaînes de caractères définit l'ordre alphabétique.

# Entrée
- tab : le tableau de chaînes de caractères à trier

# Info
2022-2023, test3, exercice 1
*/

func trier(tab []string) {
	n := len(tab)
	for i := 0; i < n; i++ {
		index_max := i
		for a := i + 1; a < n; a++ {
			if len(tab[a]) > len(tab[index_max]) || (len(tab[a]) == len(tab[index_max]) && tab[a] < tab[index_max]) {
				index_max = a
			}
		}
		tab[i], tab[index_max] = tab[index_max], tab[i]
	}
}
