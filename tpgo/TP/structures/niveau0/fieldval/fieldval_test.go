package fieldval

import "testing"

func TestEtudiant(t *testing.T) {
	s := student{
		firstName: "test",
		lastName:  "student",
		age:       17,
	}

	v := getAge(s)
	if v != 17 {
		t.Fatal("L'age attendu est 17 mais vous retournez", v)
	}
}
