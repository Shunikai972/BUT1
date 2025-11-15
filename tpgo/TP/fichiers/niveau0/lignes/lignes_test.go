package lignes

import "testing"

func TestPasDeFichierEtudiant(t *testing.T) {
	if lignes("cefichiernexistepas") != -1 {
		t.Fail()
	}
}

func TestFichierVideEtudiant(t *testing.T) {
	if lignes("../../fichiers-tests/vide") != 0 {
		t.Fail()
	}
}

func TestFichier1Etudiant(t *testing.T) {
	if lignes("../../fichiers-tests/fichier1") != 5 {
		t.Fail()
	}
}

func TestFichier2Etudiant(t *testing.T) {
	if lignes("../../fichiers-tests/fichier2") != 5 {
		t.Fail()
	}
}

func TestFichier3Etudiant(t *testing.T) {
	if lignes("../../fichiers-tests/fichier3") != 50 {
		t.Fail()
	}
}
