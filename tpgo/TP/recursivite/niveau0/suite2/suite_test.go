package suite

import "testing"

func TestU0Etudiant(t *testing.T) {
	u0 := terme(0)
	if u0 != 2 {
		t.Error("u0 doit valoir 2, mais vous retournez", u0)
	}
}

func TestU5Etudiant(t *testing.T) {
	u5 := terme(5)
	if u5 != 95 {
		t.Error("u5 doit valoir 95, mais vous retournez", u5)
	}
}
