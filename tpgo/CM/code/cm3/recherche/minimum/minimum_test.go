package minimum

import "testing"

func TestNonInitialise(t *testing.T) {
	position := minimum(nil)
	if position != -1 {
		t.Fatal("pb avec un tableau non initialise")
	}
}

func TestVide(t *testing.T) {
	position := minimum([]int{})
	if position != -1 {
		t.Fatal("pb avec un tableau vide")
	}
}

func TestSingletonOK(t *testing.T) {
	position := minimum([]int{1})
	if position != 0 {
		t.Fatal("pb avec un singleton OK")
	}
}

func TestOK(t *testing.T) {
	position := minimum([]int{12, 32, 7, 23, 9})
	if position != 2 {
		t.Fatal("pb avec un singleton OK")
	}
}
