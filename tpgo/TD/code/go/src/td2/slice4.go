package main

import "fmt"

func main() {
	var itab []int = make([]int, 5)
	var btab []bool = make([]bool, 3)
	fmt.Println(itab, btab)
}
package main

import "fmt"

// Fonction qui crée un tableau des n premiers multiples de 7
func multiplesDe7(n int) []int {
    tab := make([]int, n) // création d'une slice de taille n
    for i := 0; i < n; i++ {
        tab[i] = 7 * (i + 1)
    }
    return tab
}

func main() {
    fmt.Println(multiplesDe7(0))   // []
    fmt.Println(multiplesDe7(5))   // [7 14 21 28 35]
    fmt.Println(multiplesDe7(10))  // [7 14 21 28 35 42 49 56 63 70]
    fmt.Println(multiplesDe7(100)) // 100 premiers multiples de 7
}
