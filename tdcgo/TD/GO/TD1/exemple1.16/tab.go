package main

import "fmt"

	func tab(tint []int) int{
	return tint[3]
}
func main() {
	var tint []int = []int{1, 2, 3, 1000, 27}
	fmt.Println(tab(tint))
}
