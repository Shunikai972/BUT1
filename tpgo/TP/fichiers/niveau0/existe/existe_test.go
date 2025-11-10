package existe

import "testing"

func TestPresentEtudiant(t *testing.T) {
	if !existe("../../fichiers-tests/fichier1") {
		t.Fail()
	}
}

func TestAbsentEtudiant(t *testing.T) {
	if existe("cefichiernexistepas") {
		t.Fail()
	}
}
