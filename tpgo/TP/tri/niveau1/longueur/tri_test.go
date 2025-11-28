package tri

import (
	"slices"
	"testing"
)

func TestPetit(t *testing.T) {
	attendu := []string{"a", "aa", "aaa"}
	res := classer([]string{"aa", "a", "aaa"})

	if !slices.Equal(attendu, res) {
		t.Fail()
	}
}

func TestGrand(t *testing.T) {
	attendu := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa"}
	res := classer([]string{"aaaaaa", "aa", "a", "aaaaa", "aaa", "aaaaaaa", "aaaa"})

	if !slices.Equal(attendu, res) {
		t.Fail()
	}
}

// Ajoutés après le test

func TestVide(t *testing.T) {
	attendu := []string{}
	res := classer([]string{})

	if !slices.Equal(attendu, res) {
		t.Fail()
	}
}

func TestAutre(t *testing.T) {
	attendu := []string{"un", "test", "cadeau", "bonjour", "abcdefghijklmnopqrstuvwxyz"}
	res := classer([]string{attendu[1], attendu[4], attendu[0], attendu[3], attendu[2]})

	if !slices.Equal(attendu, res) {
		t.Fail()
	}
}
