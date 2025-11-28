package decroissant

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func TestVide(t *testing.T) {
	tab := []int{}
	decroissant(tab)
	if len(tab) != 0 || tab == nil {
		t.Fail()
	}
	tab = nil
	decroissant(tab)
	if tab != nil {
		t.Fail()
	}
}

func TestDejaTrie(t *testing.T) {
	tab := []int{5, 4, 3, 2, 1}
	decroissant(tab)
	if len(tab) != 5 {
		t.Fail()
	} else if tab[0] != 5 ||
		tab[1] != 4 ||
		tab[2] != 3 ||
		tab[3] != 2 ||
		tab[4] != 1 {
		t.Fail()
	}
}

func TestCroissant(t *testing.T) {
	tab := []int{1, 2, 3, 4, 5}
	decroissant(tab)
	if len(tab) != 5 {
		t.Fail()
	} else if tab[0] != 5 ||
		tab[1] != 4 ||
		tab[2] != 3 ||
		tab[3] != 2 ||
		tab[4] != 1 {
		t.Fail()
	}
}

func TestMelange(t *testing.T) {
	tab := []int{3, 5, 2, 1, 4}
	decroissant(tab)
	if len(tab) != 5 {
		t.Fail()
	} else if tab[0] != 5 ||
		tab[1] != 4 ||
		tab[2] != 3 ||
		tab[3] != 2 ||
		tab[4] != 1 {
		t.Fail()
	}
}

func TestEgaux(t *testing.T) {
	tab := []int{3, 5, 3, 1, 4}
	decroissant(tab)
	if len(tab) != 5 {
		t.Fail()
	} else if tab[0] != 5 ||
		tab[1] != 4 ||
		tab[2] != 3 ||
		tab[3] != 3 ||
		tab[4] != 1 {
		t.Fail()
	}
}

// Cas de test pour le contrôle de 2022-2023

func TestVide2(t *testing.T) {
	res := []int{}
	tab := []int{}
	decroissant(tab)
	if !equal(tab, res) {
		t.Fail()
	}
}

func TestExemple(t *testing.T) {
	res := []int{7, 5, 4, 3, -1, -2, -6}
	tab := []int{3, 4, -1, 7, -2, -6, 5}
	decroissant(tab)
	if !equal(tab, res) {
		t.Fail()
	}
}

// Ajoutés après le test

func TestExemple2(t *testing.T) {
	res := []int{1}
	tab := []int{1}
	decroissant(tab)
	if !equal(tab, res) {
		t.Fail()
	}
}

func TestExemple3(t *testing.T) {
	tab := []int{1, 2, 3, 4, 5}
	res := []int{5, 4, 3, 2, 1}
	decroissant(tab)
	if !equal(tab, res) {
		t.Fail()
	}
}

func TestExemple4(t *testing.T) {
	res := []int{-1, -2, -3, -4, -5}
	tab := []int{-1, -2, -3, -4, -5}
	decroissant(tab)
	if !equal(tab, res) {
		t.Fail()
	}
}

// Fonction pour vérifier l'égalité de deux tableaux

func equal(t1, t2 []int) bool {
	if len(t1) != len(t2) {
		return false
	}
	for i := 0; i < len(t1); i++ {
		if t1[i] != t2[i] {
			return false
		}
	}
	return true
}

// Cas de test pour le contrôle de 2023-2024

func TestVide3(t *testing.T) {
	var tab []int
	decroissant(tab)
	if len(tab) != 0 {
		t.Fail()
	}
}

func TestExemple5(t *testing.T) {
	var tab []int = []int{1, 2, 3, 4, 5}
	var res []int = []int{5, 4, 3, 2, 1}
	decroissant(tab)
	if !slices.Equal(tab, res) {
		t.Fail()
	}
}

// Ajoutés après le DS

func Test1(t *testing.T) {
	var tab []int = []int{1}
	var res []int = []int{1}
	decroissant(tab)
	if !slices.Equal(tab, res) {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	var tab []int = []int{7, 3}
	var res []int = []int{7, 3}
	decroissant(tab)
	if !slices.Equal(tab, res) {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	var tab []int = []int{3, 7}
	var res []int = []int{7, 3}
	decroissant(tab)
	if !slices.Equal(tab, res) {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	var res []int = []int{125, 65, 45, 32, 25, 24, 23, 14, 12, 11, 9, 5, 4, 2, -1, -5, -10, -1000}
	var tab []int = make([]int, len(res))
	copy(tab, res)

	for i := 0; i < 25; i++ {
		rand.Shuffle(len(tab), func(i, j int) { tab[i], tab[j] = tab[j], tab[i] })
		decroissant(tab)
		if !slices.Equal(tab, res) {
			t.Errorf("J'attendais %d et j'ai eu %d", res, tab)
		}
	}
}
