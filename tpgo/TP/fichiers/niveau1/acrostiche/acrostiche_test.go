package acrostiche

import "testing"

func TestFichier1Etudiant(t *testing.T) {
	if acrostiche("../../fichiers-tests/fichier1") != "lllll" {
		t.Fail()
	}
}

func TestFichier2Etudiant(t *testing.T) {
	if acrostiche("../../fichiers-tests/fichier2") != "lll" {
		t.Fail()
	}
}

func TestBonjourEtudiant(t *testing.T) {
	if acrostiche("../../fichiers-tests/bonjour") != "Bonjour" {
		t.Fail()
	}
}
