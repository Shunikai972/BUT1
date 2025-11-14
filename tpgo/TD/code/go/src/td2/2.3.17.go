package main

import "fmt"

func main() {
	var n int = 5
	var m int = 10
	t := make([][]int, n)

	for i := 0; i < n; i++ {
		t[i] = make([]int, m)
		for j := 0; j < m; j++ {
			t[i][j] = i * j
		}
	}
	fmt.Println("Tableau de con : ")
	fmt.Println(t)
}
