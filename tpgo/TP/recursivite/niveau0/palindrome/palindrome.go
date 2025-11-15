package palindrome

/*
Un palindrome est une chaîne de caractères qui se lit exactement pareil dans les deux sens (sans prendre en compte les espaces et les accents). On peut aussi donner une définition récursive d'un palindrome (une fois les espaces et accents retirés d'une chaîne) : la chaîne xsy (avec x, y des lettres et s une chaîne) est un palindrome si et seulement si x = y et s est un palindrome. La fonction estPalindrome doit implanter cette définition récursive.

# Entrée
- s : une chaîne de caractère (sans espaces, sans accents, différente de nil)

# Sortie
- b : un booléen indiquant si s est un palindrome ou pas

# Exemples
estPalindrome("bonjour") = false
estPalindrome("kayak") = true
*/

func estPalindrome(s string) bool {
    // Fonction récursive interne
    var helper func(str string) bool
    helper = func(str string) bool {
        // Si la chaîne est vide ou a un seul caractère, c'est un palindrome
        if len(str) <= 1 {
            return true
        }
        // Comparer le premier et le dernier caractère
        if str[0] != str[len(str)-1] {
            return false
        }
        // Appel récursif sur la sous-chaîne sans le premier et le dernier caractère
        return helper(str[1 : len(str)-1])
    }
    return helper(s)
}