package lirepointeur

import "testing"

func TestEtudiant(t *testing.T) {
	var n int = 5
	if lire(&n) != 5 {
		t.Fail()
	}
}

// Ajoutés après le DS

func Test1(t *testing.T) {
	var n int = 6
	if lire(&n) != 6 {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	var n int = 6
	if n != 6 {
		t.Fail()
	}
}
