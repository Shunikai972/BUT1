package recursive

import "testing"

func TestRecherhceDicoVide1(t *testing.T) {
	position := rechercheDicho(1, nil)
	if position != -1 {
		t.Fatal("pb avec un tableau non initialise")
	}
}
func TestRecherhceDicoVide2(t *testing.T) {
	position := rechercheDicho(1, []int{})
	if position != -1 {
		t.Fatal("pb avec un tableau vide")
	}
}

func TestRecherhceDicoPresent1(t *testing.T) {
	position := rechercheDicho(1, []int{1})
	if position != 0 {
		t.Fatal("pb avec un element présent dans un tableau de taille 1")
	}
}

func TestRecherhceDicoAbsent1(t *testing.T) {
	position := rechercheDicho(2, []int{1})
	if position != -1 {
		t.Fatal("pb avec un element absent dans un tableau de taille 1", position)
	}
}

func TestRecherhceDicoPresent2(t *testing.T) {
	position := rechercheDicho(2, []int{1, 2})
	if position != 1 {
		t.Fatal("pb avec un element présent dans un tableau de taille 2 au début")
	}
}

func TestRecherhceDicoPresent3(t *testing.T) {
	position := rechercheDicho(1, []int{1, 2, 3})
	if position != 0 {
		t.Fatal("pb avec un element présent dans un tableau de taille 3 au debut")
	}
}

func TestRecherhceDicoPresent4(t *testing.T) {
	position := rechercheDicho(2, []int{1, 2, 3})
	if position != 1 {
		t.Fatal("pb avec un element présent dans un tableau de taille 3 au mileu")
	}
}

func TestRecherhceDicoPresent5(t *testing.T) {
	position := rechercheDicho(3, []int{1, 2, 3})
	if position != 2 {
		t.Fatal("pb avec un element présent dans un tableau de taille 3 à la fin")
	}
}
func TestRecherhceDicoAbsent2(t *testing.T) {
	position := rechercheDicho(3, []int{1, 2})
	if position != -1 {
		t.Fatal("pb avec un element absent dans un tableau de taille 1")
	}
}
