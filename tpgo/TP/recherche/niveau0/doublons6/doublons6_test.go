package doublons6

import "testing"

func Test1(t *testing.T) {
	if doublons([]int{1, 2, 3, 5, 6}) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	if !doublons([]int{2, 3, 4, 5, 6}) {
		t.Fail()
	}
}

// Ajoutés après le test

type test struct {
	tab []int
	ok  bool
}

var testSet []test = []test{
	{tab: []int{1, 2, 3, 5, 6}, ok: false},
	{tab: []int{2, 3, 4, 5, 6}, ok: true},
	{tab: []int{6, 3, 4, 5, 2}, ok: false},
	{tab: []int{3}, ok: true},
	{tab: []int{-2, -3, -4}, ok: false},
	{tab: []int{-4, -3, -2}, ok: true},
	{tab: []int{-4, -3, -2, -1, 0, 1, 2}, ok: true},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {
		ok := doublons(aTester.tab)
		if aTester.ok != ok {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}
