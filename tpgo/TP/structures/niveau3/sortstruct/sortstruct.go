package sortstruct

/*
On souhaite classer un groupe d'étudiants en fonction des notes obtenues à un contrôle. On représente les étudiants par la structure student et on les place dans un tableau. L'objectif de la fonction sortByMark est de trier ce tableau de manière à ce que les étudiants soient classés de celle ou celui qui a eu la meilleure note à celle ou celui qui à eu la moins bonne note.

On pourra considérer qu'il n'y a jamais deux étudiants qui ont la même note.

# Entrée
- t : un tableau d'étudiants qui doit être trié en place de la meilleure note à la moins bonne note

# Info
2023-2024, test 3, exercice 3
*/

type student struct {
	firstName string
	lastName  string
	age       int
	mark      float64
}

func sortByMark(t []student) {}
