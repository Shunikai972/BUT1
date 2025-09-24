package main

import (
	"fmt"
)


func SommeDesDiviseurs(n int) int {
	if n <= 1 {
		return 0
	}

	somme := 1 
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			somme += i
			if i != n/i {
				somme += n / i
			}
		}
	}
	return somme
}


func SontAmicaux(a, b int) bool {
	if a <= 0 || b <= 0 {
		return false 
	}
	return SommeDesDiviseurs(a) == b && SommeDesDiviseurs(b) == a
}

func main() {
	
	fmt.Println("Tests de SommeDesDiviseurs:")
	fmt.Printf("Somme des diviseurs de 0 : %d\n", SommeDesDiviseurs(0))  // Cas particulier
	fmt.Printf("Somme des diviseurs de 1 : %d\n", SommeDesDiviseurs(1))  // Cas particulier
	fmt.Printf("Somme des diviseurs de 2 : %d\n", SommeDesDiviseurs(2))  // Premier
	fmt.Printf("Somme des diviseurs de 6 : %d\n", SommeDesDiviseurs(6))  // Parfait
	fmt.Printf("Somme des diviseurs de 28 : %d\n", SommeDesDiviseurs(28)) // Parfait

	// Tests de la fonction SontAmicaux
	fmt.Println("\nTests de SontAmicaux:")
	fmt.Printf("220 et 284 : %v (doit être vrai)\n", SontAmicaux(220, 284))
	fmt.Printf("1184 et 1210 : %v (doit être vrai)\n", SontAmicaux(1184, 1210))
	fmt.Printf("6 et 28 : %v (doit être faux)\n", SontAmicaux(6, 28))
	fmt.Printf("2 et 3 : %v (doit être faux)\n", SontAmicaux(2, 3))
	fmt.Printf("0 et 0 : %v (doit être faux)\n", SontAmicaux(0, 0))
}
