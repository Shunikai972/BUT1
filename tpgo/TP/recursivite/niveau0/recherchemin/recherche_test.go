package recherche

import "testing"

func TestPetitEtudiant(t *testing.T) {
	if trouvePlusPetit([]int{2, 1, 3}) != 1 {
		t.Fail()
	}
}

func TestGrandEtudiant(t *testing.T) {
	if trouvePlusPetit([]int{4, 5, 2, 1, 3, 7, 6}) != 1 {
		t.Fail()
	}
}

// Ajoutés après le DS

func TestAutre(t *testing.T) {
	if trouvePlusPetit([]int{1, 1, 1, 1, 1}) != 1 {
		t.Fail()
	}
}

func TestAutre2(t *testing.T) {
	if trouvePlusPetit([]int{1024}) != 1024 {
		t.Fail()
	}
}

func TestAutre3(t *testing.T) {
	if trouvePlusPetit([]int{-1, -1000, -2, -2000, 2000, 2, 1000, 1}) != -2000 {
		t.Fail()
	}
}
