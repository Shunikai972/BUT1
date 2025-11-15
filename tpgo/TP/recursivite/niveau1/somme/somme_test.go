package somme

import "testing"

func TestEtudiant(t *testing.T) {
	if somme([]int{1, 2, 3, 4}) != 10 {
		t.Fail()
	}
}

// Ajoutés après le test

type test struct {
	t []int
	s int
}

var testSet []test = []test{
	{t: []int{1, 2, 3, 4}, s: 10},
	{t: []int{0, 0, 0, 0}, s: 0},
	{t: []int{}, s: 0},
	{t: []int{-1, 2, 3}, s: 4},
	{t: []int{-1, -1, -1, -1, -1}, s: -5},
	{t: []int{0, 2, 100, 1, -7, 12}, s: 108},
	{t: []int{1, 1, 1}, s: 3},
	{t: []int{-2}, s: -2},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {
		s := somme(aTester.t)
		if aTester.s != s {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}
