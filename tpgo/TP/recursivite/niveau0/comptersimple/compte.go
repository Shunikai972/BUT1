package compte

/*
Étant donné un entier n, on souhaite construire la chaîne de caractères contenant tous les nombres de 0 à n dans l'ordre et séparés par des virgules. Par exemple, si n vaut 5 on souhaite construire la chaîne "0,1,2,3,4,5". C'est le rôle de la fonction compte.

Pour cet exercice les boucles for sont interdites.

On rappelle qu'on peut utiliser la fonction Sprint du paquet fmt pour construire des chaînes de caractères. Par exemple, fmt.Sprint(1, ",", 2) construit la chaîne "1,2".

# Entrée
- n : un entier positif

# Sortie
- s : une chaîne de caractères contenant les nombres de 0 à n dans l'ordre et séparés par des virgules

# Info
2024-2025, test 3, exercice 5
*/

import "fmt"

func compte(n uint) string {
    if n == 0 {
        return "0" 
    }
   
    return compte(n-1) + "," + fmt.Sprint(n)
}
