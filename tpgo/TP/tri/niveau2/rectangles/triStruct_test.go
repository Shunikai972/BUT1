package triStruct

import (
	"slices"
	"testing"
)

func TestPetit(t *testing.T) {
	attendu := []rectangle{
		{largeur: 1, longueur: 2},
		{largeur: 3, longueur: 4},
		{largeur: 5, longueur: 6},
	}

	res := ranger([]rectangle{attendu[1], attendu[2], attendu[0]})

	if !slices.Equal(attendu, res) {
		t.Fail()
	}
}

func TestGrand(t *testing.T) {
	attendu := []rectangle{
		{largeur: 1, longueur: 2},
		{largeur: 1, longueur: 4},
		{largeur: 2, longueur: 4},
		{largeur: 3, longueur: 4},
		{largeur: 3, longueur: 5},
		{largeur: 5, longueur: 6},
	}

	res := ranger([]rectangle{attendu[1], attendu[2], attendu[4], attendu[0], attendu[3], attendu[5]})

	if !slices.Equal(attendu, res) {
		t.Fail()
	}
}

// Ajoutés après le DS

func TestVide(t *testing.T) {
	attendu := []rectangle{}

	res := ranger([]rectangle{})

	if !slices.Equal(attendu, res) {
		t.Fail()
	}
}

func TestAutre(t *testing.T) {
	attendu := []rectangle{
		{largeur: 1, longueur: 1},
		{largeur: 1, longueur: 2},
		{largeur: 1, longueur: 3},
		{largeur: 2, longueur: 2},
		{largeur: 1, longueur: 5},
		{largeur: 2, longueur: 3},
		{largeur: 1, longueur: 7},
		{largeur: 2, longueur: 4},
		{largeur: 3, longueur: 3},
		{largeur: 2, longueur: 5},
		{largeur: 1, longueur: 11},
	}

	res := ranger([]rectangle{attendu[1], attendu[9], attendu[2], attendu[6], attendu[4], attendu[7], attendu[10], attendu[0], attendu[3], attendu[5], attendu[8]})

	if !slices.Equal(attendu, res) {
		t.Fail()
	}
}
