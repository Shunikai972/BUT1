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
	if lignes("../../fichiers-tests/fichier2") != 3 {
		t.Fail()
	}
}

func TestFichier3Etudiant(t *testing.T) {
	if lignes("../../fichiers-tests/fichier3") != 37 {
		t.Fail()
	}
}

func TestFichier4Etudiant(t *testing.T) {
	if lignes("../../fichiers-tests/fichier4") != 3 {
		t.Fail()
	}
}
