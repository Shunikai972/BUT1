package file2struct

import "testing"

func TestEtudiant(t *testing.T) {
	stud := file2struct("../../fichiers-tests/exemple.stud")
	if stud.firstName != "Luffy" {
		t.Error("Le prénom attendu est Luffy mais vous avez trouvé", stud.firstName)
	}
	if stud.lastName != "Monkey D." {
		t.Error("Le nom attendu est Monkey D. mais vous avez trouvé", stud.lastName)
	}
	if stud.age != 19 {
		t.Error("L'age attendu est 19 mais vous avez trouvé", stud.age)
	}
}

// Ajouté après le test machine
func Test2(t *testing.T) {
	stud := file2struct("../../fichiers-tests/exemple2.stud")
	if stud.firstName != "Plip" {
		t.Fail()
	}
	if stud.lastName != "Plop" {
		t.Fail()
	}
	if stud.age != 32 {
		t.Fail()
	}
}
