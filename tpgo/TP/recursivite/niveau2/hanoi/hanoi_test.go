package hanoi

import (
	"slices"
	"testing"
)

func TestHanoi3Etudiant(t *testing.T) {
	expected := []int{1, 0, 5, 1, 2, 3, 1}
	res := hanoi(3)

	if !slices.Equal(expected, res) {
		t.Fail()
	}
}

func TestHanoi4Etudiant(t *testing.T) {
	expected := []int{0, 1, 3, 0, 4, 5, 0, 1, 3, 2, 4, 3, 0, 1, 3}
	res := hanoi(4)

	if !slices.Equal(expected, res) {
		t.Fail()
	}
}

// Ajoutés après le test

func TestHanoi1(t *testing.T) {
	expected := []int{1}
	res := hanoi(1)

	if !slices.Equal(expected, res) {
		t.Fail()
	}
}

func TestHanoi2(t *testing.T) {
	expected := []int{0, 1, 3}
	res := hanoi(2)

	if !slices.Equal(expected, res) {
		t.Fail()
	}
}

func TestHanoi7(t *testing.T) {
	expected := []int{1, 0, 5, 1, 2, 3, 1, 0, 5, 4, 2, 5, 1, 0, 5, 1, 2, 3, 1, 2, 5, 4, 2, 3, 1, 0, 5, 1, 2, 3, 1, 0, 5, 4, 2, 5, 1, 0, 5, 4, 2, 3, 1, 2, 5, 4, 2, 5, 1, 0, 5, 1, 2, 3, 1, 0, 5, 4, 2, 5, 1, 0, 5, 1, 2, 3, 1, 2, 5, 4, 2, 3, 1, 0, 5, 1, 2, 3, 1, 2, 5, 4, 2, 5, 1, 0, 5, 4, 2, 3, 1, 2, 5, 4, 2, 3, 1, 0, 5, 1, 2, 3, 1, 0, 5, 4, 2, 5, 1, 0, 5, 1, 2, 3, 1, 2, 5, 4, 2, 3, 1, 0, 5, 1, 2, 3, 1}
	res := hanoi(7)

	if !slices.Equal(expected, res) {
		t.Fail()
	}
}
