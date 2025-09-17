package main

import "fmt"

func main() {
	var temp float64
	var tableau []int = []int{0, 20}
	taille := len(tableau)
	for i := 0; i < len(tableau); i++ {
		temp = temp + float64(tableau[i])
		taille = len(tableau)
	}
	x := temp / float64(taille)
	fmt.Println(x)
}
