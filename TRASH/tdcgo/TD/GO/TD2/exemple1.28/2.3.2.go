package main

import "fmt"

func main() {
    // Déclaration de tableaux
    var tab1 [5]int
    var tab2 [3]string
    var tab3 [1]bool

    // Affichage des valeurs par défaut
    fmt.Println("Valeurs par défaut du tableau [5]int :", tab1)  // Valeur par défaut : [0 0 0 0 0]
    fmt.Println("Valeurs par défaut du tableau [3]string :", tab2) // Valeur par défaut : ["", "", ""]
    fmt.Println("Valeurs par défaut du tableau [1]bool :", tab3)   // Valeur par défaut : [false]
}
