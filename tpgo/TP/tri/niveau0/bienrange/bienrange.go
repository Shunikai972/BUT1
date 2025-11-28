package bienrange

/*
La fonction bienrange indique si un tableau d'entiers est bien trié par ordre
croissant ou pas.

# Entrée
- tab : le tableau d'entiers à analyser

# Sortie
- estrange : un booléen qui vaut true si les entiers de tab sont bien triés en
ordre croissant et false s'ils ne sont pas bien triés.
*/

func bienrange(tab []int) (estrange bool) {
	if len(tab) < 2{
		return true
	}
	croissnt := true
	decroissnt := true
	for i := 0; i< len(tab)-1; i++ {
		if tab[i+1]> tab[i]{
			decroissnt = false 
		}else if tab[i+1]< tab[i]{
			croissnt = false 
		}
		
	} 
	return croissnt || decroissnt
}
