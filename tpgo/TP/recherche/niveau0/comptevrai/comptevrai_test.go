package comptevrai

import "testing"

func TestVide(t *testing.T) {
	nombre := compteVrai(nil)
	if nombre != 0 {
		t.Error("Dans le tableau vide vous indiquez que la valeur true apparaît", nombre, "fois alors qu'elle apparaît 0 fois")
	}
}

func TestPasVide(t *testing.T) {
	nombre := compteVrai([]bool{true, false, true, false})
	if nombre != 2 {
		t.Error("Dans le tableau [true false true false] vous indiquez que la valeur true apparaît", nombre, "fois alors qu'elle apparaît 2 fois")
	}
}

// Ajouté après le DS

type test struct {
	t      []bool
	nombre int
}

var testSet []test = []test{
	{t: nil, nombre: 0},
	{t: []bool{}, nombre: 0},
	{t: []bool{true, false, true, false}, nombre: 2},
	{t: []bool{false, false, false, false}, nombre: 0},
	{t: []bool{true, true, true, true}, nombre: 4},
	{t: []bool{true}, nombre: 1},
	{t: []bool{false}, nombre: 0},
	{t: []bool{false, false, true, false, false, false, true, false, false, true, true, false, false, false}, nombre: 4},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {
		if compteVrai(aTester.t) != aTester.nombre {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}
