package main

import (
	"fmt"
)

func signeProduit(a, b int) string {
	if a == 0 || b == 0 {
		return "nul"
	} else if (a > 0 && b > 0) || (a < 0 && b < 0) {
		return "positif"
	} else {
		return "négatif"
	}
}

func main() {
	// Test avec différentes valeurs
	testCases := [][2]int{
		{5, 3},
		{-4, 2},
		{0, 10},
		{-7, -8},
		{0, 0},
		{6, -1},
	}

	for _, pair := range testCases {
		a, b := pair[0], pair[1]
		fmt.Printf("Pour a=%d et b=%d, le produit est %s.\n", a, b, signeProduit(a, b))
	}
}
