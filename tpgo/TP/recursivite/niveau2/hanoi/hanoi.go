package hanoi

/*
Le problème des tours de Hanoï est un jeu mathématique assez connu et dont une solution se code naturellement de manière récursive.

Le problème : on considère des disques de rayons tous différents répartis en trois piles (A, B et C) telles que les disques d'une pile sont en ordre de rayons croissants (le plus grand rayon en bas de la pile, le plus petit en haut) et on souhaite constituer une unique pile de disques en ordre de rayons croissants en réalisant uniquement des déplacements de disques qui respectent les règles suivantes :
- on déplace un seul disque à la fois, d'une pile vers une autre,
- on ne place jamais un disque sur un autre disque dont le rayon est plus petit.
En général, et dans cet exercice, on considère que tous les disques sont initialement dans la pile A et qu'on souhaite les déplacer vers la pile C, mais le problème se généralise à n'importe quelle situation initiale (et n'importe quelle situation finale).

Une solution récursive : pour déplacer n disques de la pile A à la pile C on peut utiliser la méthode suivante :
1. déplacer n-1 disques de la pile A à la pile B,
2. déplacer un disque de la pile A à la pile C (le plus grand des n, donc, puisqu'on vient de déplacer les n-1 plus petits sur la pile B),
3. déplacer n-1 disques de la pile B à la pile C.
On peut prouver que cette méthode suit bien les règles du problème et donne une séquence de mouvements de disques qui soit la plus courte possible. Cependant, ce n'est pas l'objet de l'exercice et on va admettre que cette méthode fonctionne et se contenter de l'appliquer.

La fonction hanoi doit coder la méthode de résolution du problème des tours de Hanoï proposée ci-dessus. Elle retournera une séquence de mouvements sous la forme d'un tableau d'entiers (l'entier d'indice 0 représentant le premier mouvement, l'entier d'indice 1 représentant le deuxième mouvement, etc). Des constantes ont été définies pour représenter les déplacements par des entiers, n'hésitez pas à les utiliser pour rendre votre code plus clair. Notez bien que les 6 constantes définies sont suffisantes pour décrire les mouvements car quand on déplace un disque depuis une pile c'est toujours celui qui est au sommet et on le place toujours au sommet de sa pile d'arrivée.

# Entrée
- numDisques : le nombre de disques considérés. Initialement ils sont tous dans la pile A.

# Sortie
- mouvements : la séquence des mouvements à effectuer pour amener tous les disques sur la pile C en respectant l'algorithme proposé. Le premier mouvement à effectuer est mouvements[0] puis on doit faire mouvements[1], mouvements[2], etc.


# Info
2024-2025, test 1, exercice 9
*/

const (
	AversB int = iota // 0
	AversC            // 1
	BversA            // 2
	BversC            // 3
	CversA            // 4
	CversB            // 5
)

func hanoi(numDisques int) (mouvements []int) {
	return
}
