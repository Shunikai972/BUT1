package recherchedernier

import "testing"

func TestVide(t *testing.T) {
	_, existe := trouveDernier(1, nil)
	if existe {
		t.Error("Vous dites que l'entier 1 apparaît dans le tableau vide")
	}
}

func TestOk(t *testing.T) {
	pos, existe := trouveDernier(1, []int{1, 1, 1})
	if !existe {
		t.Error("Vous dites que l'entier 1 n'apparaît pas dans le tableau [1 1 1]")
	} else if pos != 2 {
		t.Error("Vous dites que la dernière apparition de 1 dans le tableau [1 1 1] est à la position", pos, "alors qu'elle est à la position 2")
	}
}

// Ajouté après le DS

type test struct {
	a      int
	t      []int
	pos    int
	existe bool
}

var testSet []test = []test{
	{a: 1, t: nil, existe: false},
	{a: 1, t: []int{}, existe: false},
	{a: 1, t: []int{1, 1, 1}, existe: true, pos: 2},
	{a: 5, t: []int{1, 3, 5, 7, 9, 5, 7, 3, 1}, existe: true, pos: 5},
	{a: 2, t: []int{1, 3, 5, 7, 9, 5, 7, 3, 1}, existe: false},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {
		pos, existe := trouveDernier(aTester.a, aTester.t)
		if existe != aTester.existe || (existe && pos != aTester.pos) {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}
