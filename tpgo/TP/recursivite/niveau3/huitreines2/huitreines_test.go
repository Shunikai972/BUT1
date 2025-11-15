package huitreines2

import "testing"

// Ces tests sont loins d'être exhaustifs, ils ne font que compter le nombre de
// solutions trouvées. Vous pouvez adapter les tests de huitreines si besoin pour
// vérifier un peu plus en profondeur que votre programme est correct.
func Test1Etudiant(t *testing.T) {
	if len(huitreines(1)) != 1 {
		t.Fail()
	}
}

func Test2Etudiant(t *testing.T) {
	if len(huitreines(2)) != 0 {
		t.Fail()
	}
}

func Test3Etudiant(t *testing.T) {
	if len(huitreines(3)) != 0 {
		t.Fail()
	}
}

func Test4Etudiant(t *testing.T) {
	if len(huitreines(4)) != 2 {
		t.Fail()
	}
}

func Test5Etudiant(t *testing.T) {
	if len(huitreines(5)) != 10 {
		t.Fail()
	}
}

func Test6(t *testing.T) {
	if len(huitreines(6)) != 4 {
		t.Fail()
	}
}

func Test8(t *testing.T) {
	if len(huitreines(8)) != 92 {
		t.Fail()
	}
}
