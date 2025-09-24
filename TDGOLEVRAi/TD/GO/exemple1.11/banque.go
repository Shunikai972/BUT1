package main

import (
	"fmt"
)

func agePourDisposerMontant(montantSouhaite float64) int {
	// Paramètres initiaux
	var solde float64 = 100  // Solde initial en euros
	const tauxInteret float64 = 0.035  // Taux d'intérêt annuel de 3.5%
	var age int = 0  // Anaïs commence à 0 ans

	// Boucle pour calculer le solde jusqu'à ce qu'il atteigne ou dépasse le montant souhaité
	for solde < montantSouhaite {
		age++  // Anaïs prend un an
		solde += 10 * float64(age)  // Dépôt du grand-père : 10 fois l'âge d'Anaïs
		solde *= (1 + tauxInteret)  // Application des intérêts
	}

	return age
}

func main() {
	// Demande à l'utilisateur d'entrer le montant souhaité
	var montantSouhaite float64
	fmt.Print("Entrez le montant souhaité en euros: ")
	fmt.Scanf("%f", &montantSouhaite)

	// Calculer l'âge nécessaire pour disposer du montant
	ageNecessaire := agePourDisposerMontant(montantSouhaite)

	// Afficher le résultat
	fmt.Printf("Anaïs pourra disposer de %.2f€ lorsqu'elle aura %d ans.\n", montantSouhaite, ageNecessaire)
}
