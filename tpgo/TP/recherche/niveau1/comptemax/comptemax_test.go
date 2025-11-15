package comptemax

import "testing"

func TestUn(t *testing.T) {
	val, nombre := compteMax([]int{1, 2, 3})
	if val != 3 {
		t.Error("Vous dites que le plus grand entier apparaissant dans le tableau [1 2 3] est", val, "alors que c'est 3")
	} else if nombre != 1 {
		t.Error("Vous dites que le plus grande entier (3) dans le tableau [1 2 3] apparaît", nombre, "fois alors qu'il apparaît 1 fois")
	}
}

func TestPlusieurs(t *testing.T) {
	val, nombre := compteMax([]int{1, 1, 2, 2, 3, 3})
	if val != 3 {
		t.Error("Vous dites que le plus grand entier apparaissant dans le tableau [1 1 2 2 3 3] est", val, "alors que c'est 3")
	} else if nombre != 2 {
		t.Error("Vous dites que le plus grande entier (3) dans le tableau [1 1 2 2 3 3] apparaît", nombre, "fois alors qu'il apparaît 2 fois")
	}
}

// Ajouté après le DS

type test struct {
	t      []int
	val    int
	nombre int
}

var testSet []test = []test{
	{t: nil, nombre: 0},
	{t: []int{}, nombre: 0},
	{t: []int{1, 2, 3}, val: 3, nombre: 1},
	{t: []int{1, 1, 2, 2, 3, 3}, val: 3, nombre: 2},
	{t: []int{3, 2, 1, 3, 2, 1, 3, 2, 1, 3}, val: 3, nombre: 4},
	{t: []int{3, 2, 1, 4, 3, 2, 1, 4, 3, 2, 1}, val: 4, nombre: 2},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {
		val, nombre := compteMax(aTester.t)
		if nombre != aTester.nombre || (nombre > 0 && val != aTester.val) {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}
