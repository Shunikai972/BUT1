package contient

import "testing"

func Test0Etudiant(t *testing.T) {
	e3 := element{valeur: 3}
	e2 := element{valeur: 2, suivant: &e3}
	e1 := element{valeur: 1, suivant: &e2}
	e0 := element{valeur: 0, suivant: &e1}
	l := liste{debut: &e0}

	if !l.contient(2) {
		t.Fail()
	}
}

func Test1Etudiant(t *testing.T) {
	e3 := element{valeur: 3}
	e2 := element{valeur: 2, suivant: &e3}
	e1 := element{valeur: 1, suivant: &e2}
	e0 := element{valeur: 0, suivant: &e1}
	l := liste{debut: &e0}

	if l.contient(4) {
		t.Fail()
	}
}

// Ajoutés après le test 2023-2024

type test struct {
	contenu  []int
	v        int
	contient bool
}

var testSet []test = []test{
	{contenu: []int{0, 1, 2, 3}, v: 2, contient: true},
	{contenu: []int{0, 1, 2, 3}, v: 4, contient: false},
	{contenu: []int{1, 5, 2, 0, 9, 2, 7, 5, 6, 3}, v: 9, contient: true},
	{contenu: []int{1, 5, 2, 0, 9, 2, 7, 5, 6, 3}, v: 12, contient: false},
	{contenu: []int{}, v: 1, contient: false},
}

func TestGeneralEtudiant(t *testing.T) {
	for numTest, aTester := range testSet {

		l1 := creerListe(aTester.contenu)
		l2 := creerListe(aTester.contenu)

		contient := l1.contient(aTester.v)
		if aTester.contient != contient || !egal(l1, l2) {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}

func creerListe(contenu []int) (l liste) {
	var e *element
	for i := len(contenu) - 1; i >= 0; i-- {
		newE := element{valeur: contenu[i], suivant: e}
		e = &newE
	}
	l.debut = e
	return
}

func egal(l1, l2 liste) bool {
	return egalElement(l1.debut, l2.debut)
}

func egalElement(e1, e2 *element) bool {
	if e1 == nil && e2 == nil {
		return true
	}

	if e1 == nil || e2 == nil {
		return false
	}

	if e1.valeur != e2.valeur {
		return false
	}

	return egalElement(e1.suivant, e2.suivant)
}

// Tests pour le DS 2024-2025

func TestExiste(t *testing.T) {

	var e3 element = element{valeur: 100, suivant: nil}
	var e2 element = element{valeur: 200, suivant: &e3}
	var e1 element = element{valeur: 300, suivant: &e2}
	var e0 element = element{valeur: 400, suivant: &e1}
	var l liste = liste{debut: &e0}

	if !l.contient(200) {
		t.Fail()
	}

}

func TestNExistePas(t *testing.T) {

	var e3 element = element{valeur: 100, suivant: nil}
	var e2 element = element{valeur: 200, suivant: &e3}
	var e1 element = element{valeur: 300, suivant: &e2}
	var e0 element = element{valeur: 400, suivant: &e1}
	var l liste = liste{debut: &e0}

	if l.contient(500) {
		t.Fail()
	}

}

// Ajoutés après le DS

func TestVide(t *testing.T) {

	var l liste = liste{debut: nil}

	if l.contient(0) {
		t.Fail()
	}

}

func TestExiste2(t *testing.T) {

	var e12 element = element{valeur: 100, suivant: nil}
	var e11 element = element{valeur: 150, suivant: &e12}
	var e10 element = element{valeur: 0, suivant: &e11}
	var e9 element = element{valeur: 3, suivant: &e10}
	var e8 element = element{valeur: 100, suivant: &e9}
	var e7 element = element{valeur: 112, suivant: &e8}
	var e6 element = element{valeur: 1024, suivant: &e7}
	var e5 element = element{valeur: 40, suivant: &e6}
	var e4 element = element{valeur: 100, suivant: &e5}
	var e3 element = element{valeur: 100, suivant: &e4}
	var e2 element = element{valeur: 200, suivant: &e3}
	var e1 element = element{valeur: 300, suivant: &e2}
	var e0 element = element{valeur: 400, suivant: &e1}
	var l liste = liste{debut: &e0}

	if !l.contient(1024) {
		t.Fail()
	}

}

func TestNExistePas2(t *testing.T) {

	var e12 element = element{valeur: 100, suivant: nil}
	var e11 element = element{valeur: 150, suivant: &e12}
	var e10 element = element{valeur: 0, suivant: &e11}
	var e9 element = element{valeur: 3, suivant: &e10}
	var e8 element = element{valeur: 100, suivant: &e9}
	var e7 element = element{valeur: 112, suivant: &e8}
	var e6 element = element{valeur: 1024, suivant: &e7}
	var e5 element = element{valeur: 40, suivant: &e6}
	var e4 element = element{valeur: 100, suivant: &e5}
	var e3 element = element{valeur: 100, suivant: &e4}
	var e2 element = element{valeur: 200, suivant: &e3}
	var e1 element = element{valeur: 300, suivant: &e2}
	var e0 element = element{valeur: 400, suivant: &e1}
	var l liste = liste{debut: &e0}

	if l.contient(1025) {
		t.Fail()
	}

}
