package multiplesdedeux

import "testing"

func TestVide(t *testing.T) {
	if !sontMultiples(nil) {
		t.Error("Vous dites que dans la tableau vide il existe des nombres qui ne sont pas multiples de 2")
	}
}

func TestOk(t *testing.T) {
	if !sontMultiples([]int{2, 4, 6}) {
		t.Error("Vous dites que dans le tableau [2 4 6] il existe des nombres qui ne sont pas multiples de 2")
	}
}

func TestPasOk(t *testing.T) {
	if sontMultiples([]int{1, 3, 5}) {
		t.Error("Vous dites que dans le tableau [1 3 5] tous les nombres sont multiples de 2")
	}
}

// Ajouté après le DS

type test struct {
	t  []int
	ok bool
}

var testSet []test = []test{
	{t: nil, ok: true},
	{t: []int{2, 4, 6}, ok: true},
	{t: []int{1, 3, 5}, ok: false},
	{t: []int{2}, ok: true},
	{t: []int{1}, ok: false},
	{t: []int{2, 4, 6, 2, 4, 6, 2, 4, 6}, ok: true},
	{t: []int{2, 4, 6, 2, 4, 6, 2, 4, 6, 1}, ok: false},
	{t: []int{2, 4, 6, 2, 4, 1, 6, 2, 4, 6}, ok: false},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {
		if sontMultiples(aTester.t) != aTester.ok {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}
