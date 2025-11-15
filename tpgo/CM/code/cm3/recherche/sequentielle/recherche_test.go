package sequentielle

import "testing"

func TestNonInitialise(t *testing.T) {
	trouve, position := rechercheElement(nil, 1)
	if trouve != false || position != -1 {
		t.Fatal("pb avec un tableau non initialise")
	}
}

func TestVide(t *testing.T) {
	trouve, position := rechercheElement([]int{}, 1)
	if trouve != false || position != -1 {
		t.Fatal("pb avec un tableau vide")
	}
}

func TestSingletonOK(t *testing.T) {
	trouve, position := rechercheElement([]int{1}, 1)
	if trouve != true || position != 0 {
		t.Fatal("pb avec un singleton OK")
	}
}

func TestSingletonK0(t *testing.T) {
	trouve, position := rechercheElement([]int{1}, 10)
	if trouve != false || position != -1 {
		t.Fatal("pb avec un singleton KO")
	}
}

func TestOK(t *testing.T) {
	trouve, position := rechercheElement([]int{12, 32, 7, 23, 9}, 7)
	if trouve != true || position != 2 {
		t.Fatal("pb avec un singleton OK")
	}
}

func TestK0(t *testing.T) {
	trouve, position := rechercheElement([]int{12, 32, 7, 23, 9}, 5)
	if trouve != false || position != -1 {
		t.Fatal("pb avec un singleton KO")
	}
}
