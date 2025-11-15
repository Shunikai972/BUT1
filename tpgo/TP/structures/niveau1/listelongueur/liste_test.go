package liste

import "testing"

func TestExempleEtudiant(t *testing.T) {

	var e3 element = element{v: 100, next: nil}
	var e2 element = element{v: 200, next: &e3}
	var e1 element = element{v: 300, next: &e2}
	var e0 element = element{v: 400, next: &e1}
	var l listeChainee = listeChainee{head: &e0}

	if numElements(l) != 4 {
		t.Fail()
	}

}

// Ajoutés après le DS

func TestVide(t *testing.T) {

	var l listeChainee = listeChainee{head: nil}

	if numElements(l) != 0 {
		t.Fail()
	}

}

func TestUn(t *testing.T) {

	var e0 element = element{v: 400, next: nil}
	var l listeChainee = listeChainee{head: &e0}

	if numElements(l) != 1 {
		t.Fail()
	}

}

func TestExemple2(t *testing.T) {

	var e12 element = element{v: 100, next: nil}
	var e11 element = element{v: 150, next: &e12}
	var e10 element = element{v: 0, next: &e11}
	var e9 element = element{v: 3, next: &e10}
	var e8 element = element{v: 100, next: &e9}
	var e7 element = element{v: 112, next: &e8}
	var e6 element = element{v: 1024, next: &e7}
	var e5 element = element{v: 40, next: &e6}
	var e4 element = element{v: 100, next: &e5}
	var e3 element = element{v: 100, next: &e4}
	var e2 element = element{v: 200, next: &e3}
	var e1 element = element{v: 300, next: &e2}
	var e0 element = element{v: 400, next: &e1}
	var l listeChainee = listeChainee{head: &e0}

	if numElements(l) != 13 {
		t.Fail()
	}

}
