package sequences

import (
	"sort"
	"testing"
)

func TestEtudiant(t *testing.T) {
	res := sequences(6, 6)
	if len(res) != 1 {
		t.Fail()
	} else if len(res[0]) != 6 {
		t.Fail()
	} else if res[0][0] != 1 || res[0][1] != 2 || res[0][2] != 3 || res[0][3] != 4 || res[0][4] != 5 || res[0][5] != 6 {
		t.Fail()
	}
}

func Test1Etudiant(t *testing.T) {
	res := sequences(3, 5)
	if len(res) != 10 {
		t.Fail()
	} else {
		for _, s := range res {
			if len(s) != 3 {
				t.Fail()
			}
		}
	}
}

// Ajoutés après le test

type test struct {
	k   int
	n   int
	seq [][]int
}

var testSet []test = []test{
	{k: 6, n: 6, seq: [][]int{{1, 2, 3, 4, 5, 6}}},
	{k: 3, n: 5, seq: [][]int{{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4}, {1, 2, 5}, {1, 3, 5}, {2, 3, 5}, {1, 4, 5}, {2, 4, 5}, {3, 4, 5}}},
	{k: 1, n: 2, seq: [][]int{{1}, {2}}},
	{k: 0, n: 0, seq: [][]int{}},
	{k: 0, n: 5, seq: [][]int{}},
	{k: 1, n: 1, seq: [][]int{{1}}},
	{k: 7, n: 8, seq: [][]int{{1, 2, 3, 4, 5, 6, 7}, {1, 2, 3, 4, 5, 6, 8}, {1, 2, 3, 4, 5, 7, 8}, {1, 2, 3, 4, 6, 7, 8}, {1, 2, 3, 5, 6, 7, 8}, {1, 2, 4, 5, 6, 7, 8}, {1, 3, 4, 5, 6, 7, 8}, {2, 3, 4, 5, 6, 7, 8}}},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {
		seq := sequences(aTester.k, aTester.n)
		if !egal(aTester.seq, seq) {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}

func egal(seq1, seq2 [][]int) bool {

	if len(seq1) != len(seq2) {
		return false
	}

	sort.Slice(seq1, func(i, j int) bool {
		return avant(seq1[i], seq1[j])
	})

	sort.Slice(seq2, func(i, j int) bool {
		return avant(seq2[i], seq2[j])
	})

	for i, s1 := range seq1 {
		if !egalTab(s1, seq2[i]) {
			return false
		}
	}

	return true

}

func egalTab(s1, s2 []int) bool {
	// tableaux supposés en ordre croissant
	if len(s1) != len(s2) {
		return false
	}

	for i, v := range s1 {
		if s2[i] != v {
			return false
		}
	}

	return true
}

func avant(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return len(s1) < len(s2)
	}

	for i, v := range s1 {
		if s2[i] != v {
			return v < s2[i]
		}
	}

	return false
}
