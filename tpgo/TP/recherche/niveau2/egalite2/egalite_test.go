package egalite

import "testing"

func TestEgal(t *testing.T) {
	if !egal([]int{1, 2, 3, 1}, []int{1, 1, 2, 3}) {
		t.Fail()
	}
}

func TestDifferent(t *testing.T) {
	if egal([]int{1, 2, 3}, []int{1, 1, 2, 3}) {
		t.Fail()
	}
}

// Ajoutés après le test

func TestEgal2(t *testing.T) {
	if !egal([]int{}, []int{}) {
		t.Fail()
	}
}

func TestEgal3(t *testing.T) {
	if !egal([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}) {
		t.Fail()
	}
}

func TestEgal4(t *testing.T) {
	if !egal([]int{3, 1, 1, 2, 2, 2, 99, 1, 1, 1, 1, 2, 2, 2, 2, 1, 1, 1, 1, 3, 1, 2, 2, 2, 1, 1, 1, 3, 1, 1}, []int{1, 1, 1, 1, 3, 1, 99, 1, 2, 1, 2, 2, 2, 1, 3, 3, 1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2, 1, 1}) {
		t.Fail()
	}
}

func TestDifferent2(t *testing.T) {
	if egal([]int{1, 1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}) {
		t.Fail()
	}
}

func TestDifferent3(t *testing.T) {
	if egal([]int{1, 1, 1, 2, 2, 2, 1, 1, 1}, []int{2, 1, 1, 2, 1, 2, 1, 1}) {
		t.Fail()
	}
}

func TestDifferent4(t *testing.T) {
	if egal([]int{1, 1, 1, 2, 2, 2, 1, 1, 1}, []int{2, 1, 1, 2, 1, 2, 1, 1, 3}) {
		t.Fail()
	}
}
