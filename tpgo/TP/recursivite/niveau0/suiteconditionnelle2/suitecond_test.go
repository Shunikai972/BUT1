package suitecond

import (
	"log"
	"testing"
)

func TestU0Etudiant(t *testing.T) {
	if u(0) != 5 {
		t.Fail()
	}
}

func TestU2Etudiant(t *testing.T) {
	if u(2) != 13 {
		t.Fail()
	}
}

// Ajoutés après le test

func TestU5(t *testing.T) {
	log.Print(u(5))
	if u(5) != 111 {
		t.Fail()
	}
}

func TestU10(t *testing.T) {
	log.Print(u(10))
	if u(10) != 973 {
		t.Fail()
	}
}
