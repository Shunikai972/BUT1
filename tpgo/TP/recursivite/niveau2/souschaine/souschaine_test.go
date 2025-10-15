package souschaine

import "testing"

func Test1Etudiant(t *testing.T) {
	if !sousChaine("", "") {
		t.Fail()
	}
}

func Test2Etudiant(t *testing.T) {
	if sousChaine("", "a") {
		t.Fail()
	}
}

func Test3Etudiant(t *testing.T) {
	if !sousChaine("a", "") {
		t.Fail()
	}
}

func Test4Etudiant(t *testing.T) {
	if !sousChaine("abcde", "ace") {
		t.Fail()
	}
}

func Test5(t *testing.T) {
	if !sousChaine("abcdefghijklmnopqrstuvwxyz", "aeiouy") {
		t.Fail()
	}
}
