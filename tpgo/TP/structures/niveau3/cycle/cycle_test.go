package cycle

import "testing"

func TestAvecCycleEtudiant(t *testing.T) {

	n1 := noeud{valeur: 1, suivant: nil}
	n2 := noeud{valeur: 2, suivant: &n1}
	n3 := noeud{valeur: 3, suivant: &n2}
	n1.suivant = &n3

	if !estCyclique(liste{tete: &n3}) {
		t.Fail()
	}
}

func TestSansCycleEtudiant(t *testing.T) {

	n1 := noeud{valeur: 1, suivant: nil}
	n2 := noeud{valeur: 2, suivant: &n1}
	n3 := noeud{valeur: 3, suivant: &n2}

	if estCyclique(liste{tete: &n3}) {
		t.Fail()
	}
}

// Ajoutés après le DS

func TestVide(t *testing.T) {
	if estCyclique(liste{tete: nil}) {
		t.Fail()
	}
}

func TestSansCycle2(t *testing.T) {

	n1 := noeud{valeur: 1, suivant: nil}

	if estCyclique(liste{tete: &n1}) {
		t.Fail()
	}
}

func TestSansCycle3(t *testing.T) {

	n1 := noeud{valeur: 1, suivant: nil}
	n2 := noeud{valeur: 2, suivant: &n1}
	n3 := noeud{valeur: 3, suivant: &n2}
	n4 := noeud{valeur: 3, suivant: &n3}
	n5 := noeud{valeur: 3, suivant: &n4}
	n6 := noeud{valeur: 3, suivant: &n5}
	n7 := noeud{valeur: 3, suivant: &n6}
	n8 := noeud{valeur: 3, suivant: &n7}

	if estCyclique(liste{tete: &n8}) {
		t.Fail()
	}
}

func TestAvecCycle2(t *testing.T) {

	n1 := noeud{valeur: 1, suivant: nil}
	n2 := noeud{valeur: 2, suivant: &n1}
	n3 := noeud{valeur: 3, suivant: &n2}
	n4 := noeud{valeur: 3, suivant: &n3}
	n5 := noeud{valeur: 3, suivant: &n4}
	n6 := noeud{valeur: 3, suivant: &n5}
	n7 := noeud{valeur: 3, suivant: &n6}
	n8 := noeud{valeur: 3, suivant: &n7}

	n1.suivant = &n5

	if !estCyclique(liste{tete: &n8}) {
		t.Fail()
	}
}
