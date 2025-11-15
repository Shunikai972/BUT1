package yetanothersuite

import (
	"testing"
)

func Test0Etudiant(t *testing.T) {
	if terme(0) != 1 {
		t.Fail()
	}
}

func Test1Etudiant(t *testing.T) {
	if terme(3) != 26 {
		t.Fail()
	}
}

// Ajoutés après le test

type test struct {
	n  uint
	un int
}

var testSet []test = []test{
	{n: 0, un: 1},
	{n: 3, un: 26},
	{n: 5, un: 458330},
	{n: 7, un: 3275921592928522330},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {
		un := terme(aTester.n)
		if aTester.un != un {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}
