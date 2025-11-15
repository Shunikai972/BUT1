package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
)

func main() {

	// Déclarer n comme big.Int
	n := new(big.Int)
	n.SetString("999999999999999999999999", 10)

	// Créer le fichier
	myFile, err := os.Create("table.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer myFile.Close()

	// Boucle big.Int
	limit := new(big.Int)
	limit.SetString("999999999999999999999999999999", 10)

	i := new(big.Int)
	one := big.NewInt(1)

	for i.SetInt64(0); i.Cmp(limit) <= 0; i.Add(i, one) {

		result := new(big.Int)
		result.Mul(i, n)

		_, err = fmt.Fprintf(myFile, "%s x %s = %s\n", i.String(), n.String(), result.String())
		if err != nil {
			log.Fatal(err)
		}
	}

}
