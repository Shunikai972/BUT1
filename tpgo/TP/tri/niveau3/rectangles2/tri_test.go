package tri

import (
	"slices"
	"testing"
)

func TestSurface(t *testing.T) {
	attendu := []rectangle{
		{nom: "bb", largeur: 3, longueur: 5},
		{nom: "a", largeur: 2, longueur: 9},
	}

	res := ranger([]rectangle{attendu[1], attendu[0]})

	if !slices.Equal(attendu, res) {
		t.Fail()
	}
}

func TestLongueur(t *testing.T) {
	attendu := []rectangle{
		{nom: "b", largeur: 3, longueur: 5},
		{nom: "aa", largeur: 3, longueur: 5},
	}

	res := ranger([]rectangle{attendu[1], attendu[0]})

	if !slices.Equal(attendu, res) {
		t.Fail()
	}
}

func TestAlpha(t *testing.T) {
	attendu := []rectangle{
		{nom: "a", largeur: 3, longueur: 5},
		{nom: "b", largeur: 3, longueur: 5},
	}

	res := ranger([]rectangle{attendu[1], attendu[0]})

	if !slices.Equal(attendu, res) {
		t.Fail()
	}
}

func TestTout(t *testing.T) {
	attendu := []rectangle{
		{nom: "a", largeur: 3, longueur: 5},
		{nom: "b", largeur: 3, longueur: 5},
		{nom: "aa", largeur: 3, longueur: 5},
		{nom: "a", largeur: 2, longueur: 9},
	}

	res := ranger([]rectangle{attendu[1], attendu[3], attendu[2], attendu[0]})

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

func TestTout2(t *testing.T) {
	attendu := []rectangle{
		{nom: "Alice", largeur: 1, longueur: 1},
		{nom: "Bob", largeur: 1, longueur: 2},
		{nom: "Claire", largeur: 1, longueur: 3},
		{nom: "Denis", largeur: 2, longueur: 2},
		{nom: "Eleonore", largeur: 1, longueur: 4},
		{nom: "Fred", largeur: 1, longueur: 5},
		{nom: "Gaby", largeur: 1, longueur: 6},
		{nom: "Hector", largeur: 2, longueur: 3},
		{nom: "Iris", largeur: 1, longueur: 7},
		{nom: "Jean", largeur: 2, longueur: 4},
		{nom: "Katy", largeur: 1, longueur: 8},
		{nom: "Lucien", largeur: 1, longueur: 9},
		{nom: "Marine", largeur: 3, longueur: 3},
		{nom: "Nicolas", largeur: 1, longueur: 10},
		{nom: "Ophelie", largeur: 2, longueur: 5},
		{nom: "Pierre", largeur: 1, longueur: 11},
	}

	res := ranger([]rectangle{attendu[6], attendu[15], attendu[1], attendu[9], attendu[3], attendu[10], attendu[2], attendu[0], attendu[14], attendu[11], attendu[4], attendu[7], attendu[5], attendu[8], attendu[13], attendu[12]})

	if !slices.Equal(attendu, res) {
		t.Fail()
	}
}
