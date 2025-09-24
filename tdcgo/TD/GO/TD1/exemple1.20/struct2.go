// struct2.go
package main

import "fmt"

type cuple struct {
	first  int
	second int
}

func main() {
	var c1 cuple = cuple{first: 50, second: 2}
	var c2 cuple = cuple{first: 60}
	var c3 cuple = cuple{second: 2}
	fmt.Println(c1, c2, c3)
}
