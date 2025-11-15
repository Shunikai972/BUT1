package factorielle

import "testing"

func TestFactorielle0(t *testing.T) {
	fac := factorielle(0)
	if fac != 1 {
		t.Fatal("pb avec factoirielle de 0")
	}
}

func TestFactorielle1(t *testing.T) {
	fac := factorielle(1)
	if fac != 1 {
		t.Fatal("pb avec factoirielle de 0")
	}
}

func TestFactorielle2(t *testing.T) {
	fac := factorielle(2)
	if fac != 2 {
		t.Fatal("pb avec factoirielle de 0")
	}
}

func TestFactorielle3(t *testing.T) {
	fac := factorielle(3)
	if fac != 6 {
		t.Fatal("pb avec factoirielle de 0")
	}
}
func TestFactorielle8(t *testing.T) {
	fac := factorielle(8)
	if fac != 40320 {
		t.Fatal("pb avec factoirielle de 0")
	}
}
