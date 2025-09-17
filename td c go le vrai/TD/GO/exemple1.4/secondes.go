package main

import (
	"fmt"
)

func main() {
	var temps int
	fmt.Println("Temps?")
	fmt.Scan(&temps)

	heure := temps / 3600
	restant := temps % 3600
	minutes := restant / 60
	secondes := restant % 60

	fmt.Printf("%d secondes correspond à : %d heures, %d minutes et %d secondes.\n", temps, heure, minutes, secondes)
}
