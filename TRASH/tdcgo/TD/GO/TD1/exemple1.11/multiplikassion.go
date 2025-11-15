package main

import "fmt"

func main() {
	var multiplikateur int = 20
	var temp = multiplikateur
	const mult int = 5
	for temp >= mult {
		multiplikateur += mult
		temp--
	}
	fmt.Println(multiplikateur)
}
