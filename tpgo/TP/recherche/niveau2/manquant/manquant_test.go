package manquant

import (
	"testing"
)

func TestManquant(t *testing.T) {
	if trouveManquant([]uint{0, 1, 3, 4}) != 2 {
		t.Fail()
	}
}

// Ajoutés après le test

func TestManquant2(t *testing.T) {
	if trouveManquant([]uint{}) != 0 {
		t.Fail()
	}
}

func TestManquant3(t *testing.T) {
	if trouveManquant([]uint{2, 5, 0, 1, 4}) != 3 {
		t.Fail()
	}
}

func TestManquant4(t *testing.T) {
	if trouveManquant([]uint{0, 1, 2, 3, 4, 5, 6}) != 7 {
		t.Fail()
	}
}

func TestManquant5(t *testing.T) {
	if trouveManquant([]uint{1, 2, 3, 4, 5, 6, 7}) != 0 {
		t.Fail()
	}
}
