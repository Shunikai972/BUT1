package iterative

import "testing"

func TestVide1(t *testing.T) {
	position := rechercheDicho(1, nil)
	if position != -1 {
		t.Fatal("pb avec un tableau non initialise")
	}
}
func TestVide2(t *testing.T) {
	position := rechercheDicho(1, []int{})
	if position != -1 {
		t.Fatal("pb avec un tableau vide")
	}
}

func TestPresent1(t *testing.T) {
	position := rechercheDicho(1, []int{1})
	if position != 0 {
		t.Fatal("pb avec un element présent dans un tableau de taille 1")
	}
}

func TestAbsent1(t *testing.T) {
	position := rechercheDicho(2, []int{1})
	if position != -1 {
		t.Fatal("pb avec un element absent dans un tableau de taille 1")
	}
}

func TestPresent2(t *testing.T) {
	position := rechercheDicho(2, []int{1, 2})
	if position != 1 {
		t.Fatal("pb avec un element présent dans un tableau de taille 2 au début")
	}
}

func TestPresent3(t *testing.T) {
	position := rechercheDicho(1, []int{1, 2, 3})
	if position != 0 {
		t.Fatal("pb avec un element présent dans un tableau de taille 3 au debut")
	}
}

func TestPresent4(t *testing.T) {
	position := rechercheDicho(2, []int{1, 2, 3})
	if position != 1 {
		t.Fatal("pb avec un element présent dans un tableau de taille 3 au mileu")
	}
}

func TestPresent5(t *testing.T) {
	position := rechercheDicho(3, []int{1, 2, 3})
	if position != 2 {
		t.Fatal("pb avec un element présent dans un tableau de taille 3 à la fin")
	}
}
func TestAbsent2(t *testing.T) {
	position := rechercheDicho(3, []int{1, 2})
	if position != -1 {
		t.Fatal("pb avec un element absent dans un tableau de taille 1")
	}
}
