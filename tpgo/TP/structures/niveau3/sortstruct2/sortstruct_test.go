package sortstruct2

import (
	"testing"
)

func TestEtudiant(t *testing.T) {
	tab := []student{
		{firstName: "test", lastName: "un", age: 18, mark: 18.5},
		{firstName: "test", lastName: "deux", age: 18, mark: 17.5},
		{firstName: "test", lastName: "trois", age: 18, mark: 19},
		{firstName: "test", lastName: "quatre", age: 42, mark: 3.5},
		{firstName: "testa", lastName: "cinq", age: 18, mark: 16},
		{firstName: "testb", lastName: "cinq", age: 18, mark: 16},
		{firstName: "test", lastName: "sept", age: 18, mark: 15},
		{firstName: "test", lastName: "huit", age: 18, mark: 20},
		{firstName: "test", lastName: "neuf", age: 18, mark: 20},
		{firstName: "test", lastName: "dix", age: 18, mark: 14.5},
	}

	sortByMarkThenLastNameThenFirstName(tab)

	expected := []student{
		{firstName: "test", lastName: "huit", age: 18, mark: 20},
		{firstName: "test", lastName: "neuf", age: 18, mark: 20},
		{firstName: "test", lastName: "trois", age: 18, mark: 19},
		{firstName: "test", lastName: "un", age: 18, mark: 18.5},
		{firstName: "test", lastName: "deux", age: 18, mark: 17.5},
		{firstName: "testa", lastName: "cinq", age: 18, mark: 16},
		{firstName: "testb", lastName: "cinq", age: 18, mark: 16},
		{firstName: "test", lastName: "sept", age: 18, mark: 15},
		{firstName: "test", lastName: "dix", age: 18, mark: 14.5},
		{firstName: "test", lastName: "quatre", age: 42, mark: 3.5},
	}

	for i, stud := range expected {
		if tab[i] != stud {
			t.Fatal(i, stud, tab[i])
		}
	}
}

// Ajouté après le test machine
func Test2(t *testing.T) {
	tab := []student{
		{firstName: "a", lastName: "a", age: 18, mark: 1},
		{firstName: "c", lastName: "a", age: 18, mark: 1},
		{firstName: "b", lastName: "a", age: 18, mark: 1},
		{firstName: "d", lastName: "b", age: 42, mark: 2},
		{firstName: "e", lastName: "b", age: 18, mark: 3},
		{firstName: "f", lastName: "c", age: 18, mark: 4},
		{firstName: "g", lastName: "c", age: 18, mark: 4},
		{firstName: "h", lastName: "c", age: 18, mark: 5},
		{firstName: "h", lastName: "c", age: 18, mark: 6},
		{firstName: "i", lastName: "d", age: 18, mark: 0},
	}

	sortByMarkThenLastNameThenFirstName(tab)

	expected := []student{
		{firstName: "h", lastName: "c", age: 18, mark: 6},
		{firstName: "h", lastName: "c", age: 18, mark: 5},
		{firstName: "f", lastName: "c", age: 18, mark: 4},
		{firstName: "g", lastName: "c", age: 18, mark: 4},
		{firstName: "e", lastName: "b", age: 18, mark: 3},
		{firstName: "d", lastName: "b", age: 42, mark: 2},
		{firstName: "a", lastName: "a", age: 18, mark: 1},
		{firstName: "b", lastName: "a", age: 18, mark: 1},
		{firstName: "c", lastName: "a", age: 18, mark: 1},
		{firstName: "i", lastName: "d", age: 18, mark: 0},
	}

	for i, stud := range expected {
		if tab[i] != stud {
			t.Fatal(i, stud, tab[i])
		}
	}
}
