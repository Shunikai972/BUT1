package recherche

import "testing"

func TestVide(t *testing.T) {
	_, existe := trouve(0, nil)
	if existe {
		t.Error("recherche(nil, 0) retourne true mais devrait retourner false")
	}
	_, existe = trouve(0, []int{})
	if existe {
		t.Error("recherche([]int{}, 0) retourne true mais devrait retourner false")
	}
}

func TestAbsent(t *testing.T) {
	_, existe := trouve(6, []int{1, 2, 3, 4, 5})
	if existe {
		t.Error("recherche([]int{1, 2, 3, 4, 5}, 6) retourne true",
			"mais devrait retourner false")
	}
}

func TestPresent(t *testing.T) {
	pos, existe := trouve(4, []int{1, 2, 3, 4, 5})
	if !existe {
		t.Error("recherche([]int{1, 2, 3, 4, 5}, 4) retourne false",
			"mais devrait retourner true")
	} else if pos != 3 {
		t.Error("recherche([]int{1, 2, 3, 4, 5}, 4) retourne", pos,
			"mais devrait retourner 3")
	}
}

// Ajoutés après le test 1 de 2023-2024

type test struct {
	a      int
	t      []int
	pos1   int
	pos2   int
	existe bool
}

var testSet []test = []test{
	{a: 0, t: nil, existe: false},
	{a: 0, t: []int{}, existe: false},
	{a: 5, t: nil, existe: false},
	{a: 5, t: []int{}, existe: false},
	{a: 0, t: []int{0, 1, 2, 3, 4, 5}, existe: true, pos1: 0},
	{a: 0, t: []int{5, 4, 3, 2, 1, 0}, existe: true, pos1: 5},
	{a: 0, t: []int{1, 3, 2, 0, 4, 5}, existe: true, pos1: 3},
	{a: 5, t: []int{0, 1, 2, 3, 4, 5}, existe: true, pos1: 5},
	{a: 5, t: []int{5, 4, 3, 2, 1, 0}, existe: true, pos1: 0},
	{a: 5, t: []int{1, 3, 2, 0, 4, 5}, existe: true, pos1: 5},
	{a: 2, t: []int{1, 2, 3, 1, 2, 3}, existe: true, pos1: 1, pos2: 4},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {
		pos, existe := trouve(aTester.a, aTester.t)
		if aTester.existe != existe || (existe && aTester.pos1 != pos && aTester.pos2 != pos) {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}
