package main

import "fmt"

func main() {
	var reste int = 20
	var quotient int = 0
	const diviseur int = 5
	for reste >= diviseur {
		reste = reste - diviseur
		quotient++
	}
	fmt.Println(quotient)
}
