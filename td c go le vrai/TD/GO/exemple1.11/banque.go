package main

import (
	"fmt"
)

func main() {
	var Thune float32
	var moula float32
	var i float32
	moula = 100
	fmt.Println("Thune?")
	fmt.Scan(&Thune)
	for Thune <= moula {
		moula += (moula * (i * 10) * 3.5)
		i++

	}

	fmt.Printf("%d", i)
}
