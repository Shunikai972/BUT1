package recherche

import "testing"

func TestPresent(t *testing.T) {
	_, existe := trouveVrai([]bool{false, true, false})
	if !existe {
		t.Fail()
	}
}

func TestPresentPos(t *testing.T) {
	pos, _ := trouveVrai([]bool{false, true, false})
	if pos != 1 {
		t.Fail()
	}
}

func TestAbsent(t *testing.T) {
	_, existe := trouveVrai([]bool{false, false, false})
	if existe {
		t.Fail()
	}
}

// Ajoutés après le DS

func TestPresent2(t *testing.T) {
	_, existe := trouveVrai([]bool{false, false, false, true, false, true, true, true})
	if !existe {
		t.Fail()
	}
}

func TestPresentPos2(t *testing.T) {
	pos, _ := trouveVrai([]bool{false, false, false, true, false, true, true, true})
	if pos != 3 {
		t.Fail()
	}
}

func TestAbsent2(t *testing.T) {
	_, existe := trouveVrai([]bool{false, false, false, false, false, false, false, false, false, false, false})
	if existe {
		t.Fail()
	}
}

func TestVide(t *testing.T) {
	_, existe := trouveVrai([]bool{})
	if existe {
		t.Fail()
	}
}
