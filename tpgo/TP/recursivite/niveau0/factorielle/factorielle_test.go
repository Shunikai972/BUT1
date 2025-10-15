package factorielle

import (
	"testing"
)

func Test1Etudiant(t *testing.T) {
	if factorielle(0) != 1 {
		t.Fail()
	}
}

func Test2Etudiant(t *testing.T) {
	if factorielle(5) != 120 {
		t.Fail()
	}
}

func Test3Etudiant(t *testing.T) {
	if factorielle(10) != 3628800 {
		t.Fail()
	}
}
