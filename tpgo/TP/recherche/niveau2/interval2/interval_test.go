package interval

import "testing"

func TestVide(t *testing.T) {
	_, existe := dansInterval(0, 100, nil)
	if existe {
		t.Error("Vous indiquez qu'il existe un flottant entre 0 et 100 dans le tableau vide")
	}
}

func TestOk(t *testing.T) {
	pos, existe := dansInterval(0, 100, []float64{250, 150, 50})
	if !existe {
		t.Error("Vous dites qu'il n'existe pas de nombre entre 0 et 100 dans le tableau [250 150 50] alors qu'il y a le nombre 50")
	} else if pos != 2 {
		t.Error("Vous dites qu'il y a un nombre entre 0 et 100 dans le tableau [250 150 50] à la position", pos, "alors que le seul qui existe est à la position 2")
	}
}

func TestPasOk(t *testing.T) {
	_, existe := dansInterval(0, 25, []float64{250, 150, 50})
	if existe {
		t.Error("Vous dites qu'il existe un nombre entre 0 et 25 dans le tableau [250 150 50]")
	}
}

// Ajouté après le DS

type test struct {
	inf, sup   float64
	t          []float64
	pos1, pos2 int
	existe     bool
}

var testSet []test = []test{
	{inf: 0, sup: 100, t: nil, existe: false},
	{inf: 0, sup: 100, t: []float64{}, existe: false},
	{inf: 0, sup: 100, t: []float64{250, 150, 50}, existe: true, pos1: 2},
	{inf: 0, sup: 25, t: []float64{250, 150, 50}, existe: false},
	{inf: 0.35, sup: 0.45, t: []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1}, existe: true, pos1: 3},
	{inf: 0.35, sup: 0.55, t: []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1}, existe: true, pos1: 3, pos2: 4},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {
		pos, existe := dansInterval(aTester.inf, aTester.sup, aTester.t)
		if existe != aTester.existe || (existe && pos != aTester.pos1 && pos != aTester.pos2) {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}
