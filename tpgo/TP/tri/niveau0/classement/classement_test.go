package classement

import "testing"

func Test(t *testing.T) {
	res := classer([]bool{true, false, true, false})
	if len(res) != 4 || !res[0] || !res[1] || res[2] || res[3] {
		t.Error("Le classement de [true false true false] devrait retourner [true true false false] mais retourne", res)
	}
}

// Ajouté après le DS

type test struct {
	t       []bool
	len     int
	numTrue int
}

var testSet []test = []test{
	{t: nil, len: 0, numTrue: 0},
	{t: []bool{}, len: 0, numTrue: 0},
	{t: []bool{false}, len: 1, numTrue: 0},
	{t: []bool{true}, len: 1, numTrue: 1},
	{t: []bool{true, true, false, false}, len: 4, numTrue: 2},
	{t: []bool{true, false, true, false}, len: 4, numTrue: 2},
	{t: []bool{true, false, false, true}, len: 4, numTrue: 2},
	{t: []bool{false, true, true, false}, len: 4, numTrue: 2},
	{t: []bool{false, true, false, true}, len: 4, numTrue: 2},
	{t: []bool{false, false, true, true}, len: 4, numTrue: 2},
	{t: []bool{false, false, false, false, false, false, false}, len: 7, numTrue: 0},
	{t: []bool{true, true, true, true, true, true, true}, len: 7, numTrue: 7},
	{t: []bool{true, true, false, true, true, false, false}, len: 7, numTrue: 4},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {
		res := classer(aTester.t)
		if len(res) != aTester.len || !classe(res, aTester.numTrue) {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}

func classe(t []bool, numTrue int) bool {
	for i, v := range t {
		if i < numTrue && !v {
			return false
		}
		if i >= numTrue && v {
			return false
		}
	}
	return true
}
