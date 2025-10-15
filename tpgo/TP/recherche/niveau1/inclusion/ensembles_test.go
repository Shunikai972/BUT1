package ensembles

import "testing"

func TestInclus(t *testing.T) {
	if !inclus([]int{1}, []int{1, 2}) {
		t.Fail()
	}
}

func TestNonInclus(t *testing.T) {
	if inclus([]int{1, 2}, []int{2, 3}) {
		t.Fail()
	}
}

// Ajoutés après le test

func TestInclus2(t *testing.T) {
	if !inclus([]int{}, []int{}) {
		t.Fail()
	}
}

func TestInclus3(t *testing.T) {
	if !inclus([]int{}, []int{1, 2}) {
		t.Fail()
	}
}

func TestNonInclus2(t *testing.T) {
	if inclus([]int{1, 2}, []int{}) {
		t.Fail()
	}
}

func TestInclus4(t *testing.T) {
	if !inclus([]int{27, 45, 2, 5, 33, 1, 3, 65, 6, 9}, []int{27, 45, 6, 2, 99, 5, 33, 912, 3, 65, 9, 1, 112}) {
		t.Fail()
	}
}

func TestNonInclus3(t *testing.T) {
	if inclus([]int{27, 45, 2, 5, 33, 1, 3, 65, 6, 9}, []int{2, 27, 112, 224, 45, 5, 33, 1, 99, 3, 65, 9, 512}) {
		t.Fail()
	}
}
