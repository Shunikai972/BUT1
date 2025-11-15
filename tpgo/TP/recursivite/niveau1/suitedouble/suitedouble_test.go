package suitedouble

import (
	"testing"
)

func Test0Etudiant(t *testing.T) {
	res := terme(0)
	if res != 5 {
		t.Error("Vous dites que U(0) vaut", res, "alors qu'il vaut 5")
	}
}

func Test1Etudiant(t *testing.T) {
	res := terme(1)
	if res != 42 {
		t.Error("Vous dites que U(1) vaut", res, "alors qu'il vaut 42")
	}
}

// Ajouté après le DS

type test struct {
	n  uint
	un int
}

var testSet []test = []test{
	{n: 0, un: 5},
	{n: 1, un: 42},
	{n: 2, un: 153},
	{n: 3, un: 486},
	{n: 4, un: 1485},
	{n: 5, un: 4482},
	{n: 6, un: 13473},
	{n: 7, un: 40446},
	{n: 8, un: 121365},
	{n: 9, un: 364122},
	{n: 10, un: 1092393},
	{n: 20, un: 64505511405},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {
		if aTester.un != terme(aTester.n) {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}
