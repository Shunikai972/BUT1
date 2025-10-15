package suitesimple

import (
	"testing"
)

func Test0Etudiant(t *testing.T) {
	res := terme(0)
	if res != 0 {
		t.Error("Vous dites que terme(0) vaut ", res, " alors qu'il devrait valoir 0")
	}
}

func Test2Etudiant(t *testing.T) {
	res := terme(2)
	if res != -20 {
		t.Error("Vous dites que terme(2) vaut", res, "alors qu'il devrait valoir -20")
	}
}

// Ajouté après le test

type test struct {
	n  uint
	un int
}

var testSet []test = []test{
	{n: 0, un: 0},
	{n: 1, un: 5},
	{n: 2, un: -20},
	{n: 3, un: 105},
	{n: 4, un: -520},
	{n: 5, un: 2605},
	{n: 6, un: -13020},
	{n: 7, un: 65105},
	{n: 8, un: -325520},
	{n: 9, un: 1627605},
	{n: 10, un: -8138020},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {
		if terme(aTester.n) != aTester.un {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}
