package file2tab
import(
	
	"bufio"
	"os"
)

/*
On souhaite lire un fichier et stocker les mots qu'il contient dans un tableau de chaines de caractères. C'est le travail de la fonction file2tab.

On pourra considérer que les fichiers dont le nom est donné à la fonction existent toujours et ne contiennent qu'une seule ligne où les mots sont séparés par des espaces (un seul espace entre chaque mot).

Le paquet strings est interdit.

# Entrée
- fName : le nom du fichier à traiter

# Sortie
- res : un tableau de chaîne de caractères qui devra contenir les mots du fichiers (un par case) dans l'ordre

# Info
2023-2024, test 3, exercice 7
*/

func file2tab(fName string) (res []string) {
	file, _:= os.Open(fName)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if 	scanner.Scan(){
		line := scanner.Text()
		currentWord := ""
		for i:=0; i<len(line); i++{
			c := line[i]
			if c == ' '{
				if currentWord != ""{
				res = append(res, currentWord)
				currentWord = ""
				}
			} else {
				currentWord += string(c)
			}
		}
		if currentWord != ""{
			res = append(res, currentWord)
		}
	}
	return
}
