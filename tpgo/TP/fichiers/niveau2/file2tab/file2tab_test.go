package file2tab

import "testing"

func TestEtudiant(t *testing.T) {
	res := file2tab("../../fichiers-tests/exemple.liste")
	if len(res) != 5 {
		t.FailNow()
	}
	if res[0] != "mot1" {
		t.Fatal(res[0], "au lieu de mot1")
	}
	if res[1] != "mot2" {
		t.Fatal(res[1], "au lieu de mot2")
	}
	if res[2] != "mot3" {
		t.Fatal(res[2], "au lieu de mot3")
	}
	if res[3] != "mot4" {
		t.Fatal(res[3], "au lieu de mot4")
	}
	if res[4] != "mot5" {
		t.Fatal(res[4], "au lieu de mot5")
	}
}

// Ajouté après le test machine
func Test2(t *testing.T) {
	res := file2tab("../../fichiers-tests/exemple2.liste")
	if len(res) != 7 {
		t.FailNow()
	}
	if res[0] != "bon" {
		t.FailNow()
	}
	if res[1] != "jour" {
		t.FailNow()
	}
	if res[2] != "a" {
		t.FailNow()
	}
	if res[3] != "tous" {
		t.FailNow()
	}
	if res[4] != "test" {
		t.FailNow()
	}
	if res[5] != "test" {
		t.FailNow()
	}
	if res[6] != "hop" {
		t.FailNow()
	}
}
