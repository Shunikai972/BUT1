package fieldval

/*
On dispose d'un type pour représenter des étudiants avec un nom, un prénom et un age. On souhaite récupérer l'age d'un étudiant. C'est le travail de la fonction getAge

# Entrée
- s : un étudiant

# Sortie
- a : l'age de s

# Info
2023-2024, test3, exercice 0
*/

type student struct {
	firstName string
	lastName  string
	age       int
}

func getAge(s student) (a int) {
	a = s.age
	return
}
