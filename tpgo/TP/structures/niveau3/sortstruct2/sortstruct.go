package sortstruct2

/*
On souhaite classer un groupe d'étudiants en fonction des notes obtenues à un contrôle. On représente les étudiants par la structure student et on les place dans un tableau. L'objectif de la fonction sortByMark est de trier ce tableau de manière à ce que les étudiants soient classés de celle ou celui qui a eu la meilleure note à celle ou celui qui à eu la moins bonne note.

Si deux étudiants ont la même note on les placera en premier dans le classement celle ou celui dont le nom est le premier en ordre alphabétique. Si ils ont en plus le même nom, on placera en premier celle ou celui dont le prénom est le premier en ordre alphabétique.

On pourra considérer qu'il n'y a pas deux étudiants qui ont à la fois la même note, le même nom et le même prénom.

On rappel que la comparaison "<" peut être utilisées sur des chaînes de caractères pour savoir la quelle est la première en ordre alphabétique.

# Entrée
- t : un tableau d'étudiants qui doit être trié en place de la meilleure note à la moins bonne note, en traitant les égalités comme indiqué ci-dessus.

# Info
2023-2024, test 3, exercice 4
*/

type student struct {
	firstName string
	lastName  string
	age       int
	mark      float64
}

func sortByMarkThenLastNameThenFirstName(t []student) {}
