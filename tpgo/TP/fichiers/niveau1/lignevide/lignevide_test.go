package lignevide

import "testing"

func TestExisteEtudiant(t *testing.T) {
	if !ligneVide("../../fichiers-tests/avec_test.txt") {
		t.Fail()
	}
}

func TestNExistePasEtudiant(t *testing.T) {
	if ligneVide("../../fichiers-tests/sans_test.txt") {
		t.Fail()
	}
}

// Ajoutés après le DS

func TestNon1(t *testing.T) {
	if ligneVide("../../fichiers-tests/vide_test.txt") {
		t.Fail()
	}
}

func TestNon2(t *testing.T) {
	if ligneVide("../../fichiers-tests/grand_test.txt") {
		t.Fail()
	}
}

func TestOui1(t *testing.T) {
	if !ligneVide("../../fichiers-tests/presquevide_test.txt") {
		t.Fail()
	}
}

func TestOui2(t *testing.T) {
	if !ligneVide("../../fichiers-tests/grand2_test.txt") {
		t.Fail()
	}
}
