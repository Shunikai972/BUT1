package main

import (
	"fmt"
)


func main() {
	var montantEuros float64
	fmt.Print("Entrez le prix ouhaité : ")
	fmt.Scan(&montantEuros)

	// Conversion euros en centimes
	if montantEuros > 300 {
		montantEuros = montantEuros * 0.95
	}
	

	fmt.Printf("Prix final:\n %f", montantEuros)
}
