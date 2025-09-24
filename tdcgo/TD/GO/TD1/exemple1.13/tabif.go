package main

import "fmt"

func main() {
	var tab []int = []int{1, 2, 3, 1000, 27, 5, 2}
	for i := 0; i < len(tab); i++{
			if tab[i]< 5{
			fmt.Println(tab[i])
		}
	}
}