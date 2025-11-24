package file2struct
import(
	"bufio"
	"os"
	"strconv"
	"strings"
)
/*
On dispose de fichiers contenant des informations sur des étudiants (nom, prénom, age) et on souhaite enregistrer ces informations dans une structure de données en Go.

La fonction file2struct lira un fichier contenant les informations sur un étudiant et utilisera ces informations pour remplir les champs d'une variable de type student.

On pourra considérer que les fichiers sont toujours bien formés : ils existent, peuvent s'ouvrir, la première ligne contient le prénom (et rien d'autre), la deuxième ligne contient le nom (et rien d'autre), la troisième ligne contient l'age (et rien d'autre) représenté par deux chiffres (par exemple 12 et pas douze ni 12,5).

# Entrée
- fName : le nom du fichier à lire

# Sortie
- res : une variable de type student qui devra contenir les informations trouvées dans le fichier

# Info
2023-2024, test 3, exercice 6
*/

type student struct {
	firstName string
	lastName  string
	age       int
}

func file2struct(fName string) (res student) {
	file, _:= os.Open(fName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	res.firstName = strings.TrimSpace(scanner.Text())

	scanner.Scan()
	res.lastName = strings.TrimSpace(scanner.Text())
	
	scanner.Scan()
	ageStr := strings.TrimSpace(scanner.Text())
	res.age, _ = strconv.Atoi(ageStr)

	return
}
