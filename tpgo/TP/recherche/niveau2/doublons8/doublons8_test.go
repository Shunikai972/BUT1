package doublons8

import "testing"

func Test1(t *testing.T) {
	if doublons([]int{2, 1, 3, 5, 6}) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	if !doublons([]int{6, 4, 4, 6, 5, 5}) {
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
	{tab: []int{6, 3, 4, 5, 2}, ok: true},
	{tab: []int{3}, ok: true},
	{tab: []int{-2, -3, -4}, ok: true},
	{tab: []int{-4, -3, -2}, ok: true},
	{tab: []int{-4, -3, -2, -1, 0, 1, 2}, ok: true},
	{tab: []int{-4, -3, -1, -1, 0, 1, 2}, ok: false},
	{tab: []int{1, 2, 3, 4, 5, 5, 6}, ok: false},
	{tab: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6}, ok: false},
	{tab: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6}, ok: true},
	{tab: []int{6, 6, 6, 6}, ok: true},
	{tab: []int{1, 2, 3, 1, 2, 3}, ok: true},
	{tab: []int{1, 2, 3, 2, 1, 2, 1, 3, 3}, ok: true},
	{tab: []int{1, 2, 3, 2, 1, 2, 1, 2, 3, 3, 2, 1, 1, 3}, ok: false},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {
		ok := doublons(aTester.tab)
		if aTester.ok != ok {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}
