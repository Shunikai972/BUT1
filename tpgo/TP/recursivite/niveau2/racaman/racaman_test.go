package racaman

import "testing"

func TestErreur1Etudiant(t *testing.T) {
	if racaman(0) != -1 {
		t.Fail()
	}
}

func TestErreur2Etudiant(t *testing.T) {
	if racaman(-1) != -1 {
		t.Fail()
	}
}

func TestValeur1Etudiant(t *testing.T) {
	if racaman(1) != 1 {
		t.Fail()
	}
}

func TestValeur2Etudiant(t *testing.T) {
	if racaman(2) != 3 {
		t.Fail()
	}
}

func TestValeur3Etudiant(t *testing.T) {
	if racaman(3) != 6 {
		t.Fail()
	}
}

func TestValeur4Etudiant(t *testing.T) {
	if racaman(4) != 2 {
		t.Fail()
	}
}

func TestValeur5Etudiant(t *testing.T) {
	if racaman(5) != 7 {
		t.Fail()
	}
}

func TestValeur6Etudiant(t *testing.T) {
	if racaman(6) != 13 {
		t.Fail()
	}
}

func TestValeur7Etudiant(t *testing.T) {
	if racaman(7) != 20 {
		t.Fail()
	}
}

func TestValeur8Etudiant(t *testing.T) {
	if racaman(8) != 12 {
		t.Fail()
	}
}

func TestValeur9Etudiant(t *testing.T) {
	if racaman(9) != 21 {
		t.Fail()
	}
}

func TestValeur10Etudiant(t *testing.T) {
	if racaman(10) != 11 {
		t.Fail()
	}
}

func TestValeur131(t *testing.T) {
	for i := 1; i < 131; i++ {
		if racaman(i) == 4 {
			t.Fail()
		}
	}
	if racaman(131) != 4 {
		t.Fail()
	}
}
