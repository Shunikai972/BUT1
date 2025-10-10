package main

import "fmt"

// fonction qui ajoute 1 à chaque élément du tableau
func ajouterUn(tab []int) {
	for i := range tab {
		tab[i] += 1
	}
}

func main() {
	// test avec plusieurs tableaux
	t1 := []int{1, 2, 3}
	t2 := []int{10, 20, 30, 40}
	t3 := []int{0, -1, -2}

	fmt.Println("Avant :", t1)
	ajouterUn(t1)
	fmt.Println("Après :", t1)

	fmt.Println("Avant :", t2)
	ajouterUn(t2)
	fmt.Println("Après :", t2)

	fmt.Println("Avant :", t3)
	ajouterUn(t3)
	fmt.Println("Après :", t3)
}
