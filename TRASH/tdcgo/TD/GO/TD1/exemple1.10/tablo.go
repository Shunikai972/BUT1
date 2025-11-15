package main

import "fmt"

func main() {
	var temp int
	var tableau []int = []int{1, 2, 3, 1000, 27, 5, 2}
	for i := 0; i < len(tableau); i++ {
		temp = temp + tableau[i]
	}
	fmt.Println(temp)
}
