package compte

import "testing"

func Test5Etudiant(t *testing.T) {
	if compte(5) != "0,1,2,3,4,5" {
		t.Fail()
	}
}

// Ajoutés après le DS
func Test0(t *testing.T) {
	if compte(0) != "0" {
		t.Fail()
	}
}

func Test1Etudiant(t *testing.T) {
	if compte(1) != "0,1" {
		t.Fail()
	}
}

func Test10Etudiant(t *testing.T) {
	if compte(10) != "0,1,2,3,4,5,6,7,8,9,10" {
		t.Fail()
	}
}
