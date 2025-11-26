package unmot

import "errors"

/*
Étant donné un nom de fichier, on souhaite obtenir le premier mot du fichier ou, si ce n'est pas possible (le fichier n'existe pas ou ne contient pas de mot), une erreur. Pour simplifier, on considérera qu'il n'existe que deux sortes de fichiers :
- des fichiers contenant exactement un mot sur leur première ligne (et éventuellement d'autres lignes ensuite)
- des fichiers ne contenant aucun mot

# Entrée
- fName : le nom du fichier à traiter

# Sorties
- mot : le premier mot du fichier fName s'il existe (n'import quoi sinon)
- err : nil s'il existe un mot dans le fichier fName, errImpossible sinon

# Info
2022-2023, test3, exercice 5
*/

var errImpossible error = errors.New("Le fichier n'existe pas ou il ne contient aucun mot")

func premiermot(fName string) (mot string, err error) {
    // Ouvrir le fichier
    f, e := os.Open(fName)
    if e != nil {
        return "", errImpossible
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

    // Lecture de la première ligne
    if !scanner.Scan() {
        return "", errImpossible
    }

    line := scanner.Text()
    fields := strings.Fields(line)

    // Vérifier s'il y a un mot
    if len(fields) == 0 {
        return "", errImpossible
    }

    // Le premier mot est fields[0]
    return fields[0], nil
}