package suite

import (
	"testing"
)

func TestU0Etudiant(t *testing.T) {
	if u(0) != 1 {
		t.Fail()
	}
}

func TestU2Etudiant(t *testing.T) {
	if u(2) != 13 {
		t.Fail()
	}
}

// Ajoutés après le test

func TestU5(t *testing.T) {
	if u(5) != 1563 {
		t.Fail()
	}
}

func TestU10(t *testing.T) {
	if u(10) != 4882813 {
		t.Fail()
	}
}
