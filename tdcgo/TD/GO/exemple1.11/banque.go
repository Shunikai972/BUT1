package main

import (
	"fmt"
)

func agePourMontant(montantCible float64) int {
	// Initialisation
	montant := 100.0         // Dépôt initial à la naissance (en centimes)
	taux := 0.035            // Taux d'intérêt annuel de 3,5%
	age := 0

	for montant < montantCible {
		age++
		// Ajout de l'intérêt
		montant += montant * taux
		// Versement de l'anniversaire : 10 * âge (en centimes)
		montant += 10 * float64(age)
	}
	return age
}

func main() {
	var montantEuros float64
	fmt.Print("Entrez le montant en euros souhaité : ")
	fmt.Scan(&montantEuros)

	// Conversion euros en centimes
	montantCible := montantEuros * 100

	age := agePourMontant(montantCible)
	fmt.Printf("Anaïs devra avoir %d ans pour disposer d'au moins %.2f euros sur son compte.\n", age, montantEuros)
}
