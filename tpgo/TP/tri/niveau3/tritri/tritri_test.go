package tritri

import (
	"slices"
	"testing"
)

func TestExemple(t *testing.T) {
	var res [][]int = [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 6},
		{2, 3, 4},
		{4, 5, 6, 7},
	}

	var tab [][]int = [][]int{
		{3, 4, 2},
		{5, 1, 2, 4, 3},
		{7, 6, 5, 4},
		{1, 6, 2, 3, 4, 5},
	}

	tritri(tab)

	if !slices.EqualFunc(res, tab, func(t, tt []int) bool {
		return slices.Equal(t, tt)
	}) {
		t.Fail()
	}
}

// Ajoutés après le DS

func TestVide(t *testing.T) {
	var res [][]int = [][]int{}

	var tab [][]int = [][]int{}

	tritri(tab)

	if !slices.EqualFunc(res, tab, func(t, tt []int) bool {
		return slices.Equal(t, tt)
	}) {
		t.Fail()
	}
}

func TestVides(t *testing.T) {
	var res [][]int = [][]int{{}}

	var tab [][]int = [][]int{{}}

	tritri(tab)

	if !slices.EqualFunc(res, tab, func(t, tt []int) bool {
		return slices.Equal(t, tt)
	}) {
		t.Fail()
	}
}

func TestGeneral(t *testing.T) {
	var res [][]int = [][]int{
		{},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 6},
		{1, 3, 5},
		{2, 3, 4},
		{4, 5, 6, 7},
		{7},
		{8, 12, 100},
		{8, 100},
		{8, 101, 102, 103},
		{8, 101, 102, 103, 104},
	}

	var tab [][]int = [][]int{
		{3, 4, 2},
		{5, 1, 2, 4, 3},
		{1, 3, 5},
		{8, 104, 101, 103, 102},
		{8, 100},
		{7},
		{7, 6, 5, 4},
		{8, 101, 103, 102},
		{8, 100, 12},
		{1, 6, 2, 3, 4, 5},
		{},
	}

	tritri(tab)

	if !slices.EqualFunc(res, tab, func(t, tt []int) bool {
		return slices.Equal(t, tt)
	}) {
		t.Fail()
	}
}
