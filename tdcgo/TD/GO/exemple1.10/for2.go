package main

import "fmt"

func main() {
	var i int
	for i = 0; i < 300; i = i + 2 {
		fmt.Println(i)
	}
	fmt.Println("AAAAAAAAA")
	fmt.Println(i)
}
