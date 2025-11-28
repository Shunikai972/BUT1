package tristruct

import (
	"math/rand"
	"slices"
	"testing"
)

func TestVide(t *testing.T) {
	var tab []chocolat
	tri(tab)
	if len(tab) != 0 {
		t.Fail()
	}
}

func TestExemple(t *testing.T) {
	var res []chocolat = []chocolat{
		{poids: 10, praline: true, noisette: true},
		{poids: 8, praline: true, noisette: true},
		{poids: 10, praline: true, noisette: false},
		{poids: 10, praline: false, noisette: true},
		{poids: 10, praline: false, noisette: false},
	}

	var tab []chocolat = []chocolat{
		res[2], res[0], res[4], res[1], res[3],
	}

	tri(tab)

	if !slices.Equal(res, tab) {
		t.Fail()
	}
}

// Ajoutés après le Ds

func TestUn(t *testing.T) {
	var res []chocolat = []chocolat{{poids: 10, praline: false, noisette: true}}
	var tab []chocolat = []chocolat{{poids: 10, praline: false, noisette: true}}
	tri(tab)
	if !slices.Equal(res, tab) {
		t.Fail()
	}
}
func TestGeneral(t *testing.T) {

	var res []chocolat = []chocolat{
		{poids: 25, praline: true, noisette: true},
		{poids: 15, praline: true, noisette: true},
		{poids: 10, praline: true, noisette: true},
		{poids: 9, praline: true, noisette: true},
		{poids: 8, praline: true, noisette: true},
		{poids: 7, praline: true, noisette: true},
		{poids: 5, praline: true, noisette: true},
		{poids: 34, praline: true, noisette: false},
		{poids: 30, praline: true, noisette: false},
		{poids: 28, praline: true, noisette: false},
		{poids: 20, praline: true, noisette: false},
		{poids: 12, praline: true, noisette: false},
		{poids: 9, praline: true, noisette: false},
		{poids: 26, praline: false, noisette: true},
		{poids: 25, praline: false, noisette: true},
		{poids: 24, praline: false, noisette: true},
		{poids: 23, praline: false, noisette: true},
		{poids: 20, praline: false, noisette: true},
		{poids: 7, praline: false, noisette: true},
		{poids: 19, praline: false, noisette: false},
		{poids: 17, praline: false, noisette: false},
		{poids: 14, praline: false, noisette: false},
		{poids: 13, praline: false, noisette: false},
		{poids: 5, praline: false, noisette: false},
		{poids: 4, praline: false, noisette: false},
	}

	for i := 0; i < 25; i++ {
		start := rand.Intn(len(res))
		end := start + rand.Intn(len(res)-start+1)
		expectedRes := res[start:end]
		tab := make([]chocolat, len(expectedRes))
		copy(tab, expectedRes)
		rand.Shuffle(len(tab), func(i, j int) { tab[i], tab[j] = tab[j], tab[i] })
		tri(tab)
		if !slices.Equal(tab, expectedRes) {
			t.Error("On attend ", expectedRes, " mais vous retournez ", tab)
		}
	}

}
