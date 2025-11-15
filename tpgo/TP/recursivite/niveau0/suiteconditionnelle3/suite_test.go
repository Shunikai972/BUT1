package suite

import (
	"testing"
)

func Test0Etudiant(t *testing.T) {
	if u(0) != 5 {
		t.Fail()
	}
}

func Test4Etudiant(t *testing.T) {
	if u(4) != 34 {
		t.Fail()
	}
}

func Test5Etudiant(t *testing.T) {
	if u(5) != 68 {
		t.Fail()
	}
}

// Ajoutés après le DS

func Test10(t *testing.T) {
	if u(10) != 514 {
		t.Fail()
	}
}

func Test15(t *testing.T) {
	if u(15) != 4097 {
		t.Fail()
	}
}

func Test20(t *testing.T) {
	if u(20) != 65540 {
		t.Fail()
	}
}
