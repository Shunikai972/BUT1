package main

import (
	"fmt"
)

func main() {
	var objectif float64
	fmt.Print("Montant souhaité (€) : ")
	fmt.Scan(&objectif)

	// Initialisation
	age := 0
	solde := 100.0
	taux := 0.035

	// Boucle infinie, on sort quand on atteint l'objectif
	for {
		if solde >= objectif {
			fmt.Printf("Anaïs devra avoir %d ans pour disposer d'au moins %.2f €.\n", age, objectif)
			break
		}
		// Anniversaire suivant
		age++
		// Ajout des intérêts
		solde += solde * taux
		// Versement du grand-père
		solde += float64(age * 10)
	}
}
