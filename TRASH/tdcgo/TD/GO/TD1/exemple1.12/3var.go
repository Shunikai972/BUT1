package main

import (
	"fmt"
)

func main() {
	var a int = 50
	var b int = 5
	var c int = 30
	var fin int = a

	if b > fin {
		fin = b
	}
	if c > fin {
		fin = c
	}

	fmt.Printf("Prix final:\n%d\n", fin)
}
