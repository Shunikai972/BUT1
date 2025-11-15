package plusgrand

import "testing"

func TestPlusGrand(t *testing.T) {
	if plusgrand([]int{1, 2, 3, 4, 5}) != 5 {
		t.Fail()
	}
}

// Ajoutés après le test

func TestPlusGrand2(t *testing.T) {
	if plusgrand([]int{5, 4, 3, 2, 1}) != 5 {
		t.Fail()
	}
}

func TestPlusGrand3(t *testing.T) {
	if plusgrand([]int{5, 6, 7, 2, 3, 9, 12, 1, 2, 5, 7}) != 12 {
		t.Fail()
	}
}

func TestPlusGrand4(t *testing.T) {
	if plusgrand([]int{1}) != 1 {
		t.Fail()
	}
}
