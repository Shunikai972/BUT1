package sortstruct

import (
	"testing"
)

func TestEtudiant(t *testing.T) {
	tab := []student{
		{firstName: "test", lastName: "1", age: 18, mark: 18.5},
		{firstName: "test", lastName: "2", age: 18, mark: 17.5},
		{firstName: "test", lastName: "3", age: 18, mark: 19},
		{firstName: "test", lastName: "4", age: 42, mark: 3.5},
		{firstName: "test", lastName: "5", age: 18, mark: 16},
		{firstName: "test", lastName: "6", age: 18, mark: 15.5},
		{firstName: "test", lastName: "7", age: 18, mark: 15},
		{firstName: "test", lastName: "8", age: 18, mark: 19.5},
		{firstName: "test", lastName: "9", age: 18, mark: 20},
		{firstName: "test", lastName: "10", age: 18, mark: 14.5},
	}

	sortByMark(tab)

	expected := []student{
		{firstName: "test", lastName: "9", age: 18, mark: 20},
		{firstName: "test", lastName: "8", age: 18, mark: 19.5},
		{firstName: "test", lastName: "3", age: 18, mark: 19},
		{firstName: "test", lastName: "1", age: 18, mark: 18.5},
		{firstName: "test", lastName: "2", age: 18, mark: 17.5},
		{firstName: "test", lastName: "5", age: 18, mark: 16},
		{firstName: "test", lastName: "6", age: 18, mark: 15.5},
		{firstName: "test", lastName: "7", age: 18, mark: 15},
		{firstName: "test", lastName: "10", age: 18, mark: 14.5},
		{firstName: "test", lastName: "4", age: 42, mark: 3.5},
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
		{firstName: "test", lastName: "1", age: 18, mark: 2},
		{firstName: "test", lastName: "2", age: 18, mark: 11},
		{firstName: "test", lastName: "3", age: 18, mark: 3},
		{firstName: "test", lastName: "4", age: 42, mark: 7},
		{firstName: "test", lastName: "5", age: 18, mark: 16},
		{firstName: "test", lastName: "6", age: 18, mark: 4},
		{firstName: "test", lastName: "7", age: 18, mark: 15},
		{firstName: "test", lastName: "8", age: 18, mark: 19.5},
		{firstName: "test", lastName: "9", age: 18, mark: 8},
		{firstName: "test", lastName: "10", age: 18, mark: 1.5},
	}

	sortByMark(tab)

	expected := []student{
		{firstName: "test", lastName: "8", age: 18, mark: 19.5},
		{firstName: "test", lastName: "5", age: 18, mark: 16},
		{firstName: "test", lastName: "7", age: 18, mark: 15},
		{firstName: "test", lastName: "2", age: 18, mark: 11},
		{firstName: "test", lastName: "9", age: 18, mark: 8},
		{firstName: "test", lastName: "4", age: 42, mark: 7},
		{firstName: "test", lastName: "6", age: 18, mark: 4},
		{firstName: "test", lastName: "3", age: 18, mark: 3},
		{firstName: "test", lastName: "1", age: 18, mark: 2},
		{firstName: "test", lastName: "10", age: 18, mark: 1.5},
	}

	for i, stud := range expected {
		if tab[i] != stud {
			t.Fatal(i, stud, tab[i])
		}
	}
}

func Test3(t *testing.T) {
	tab := []student{}

	sortByMark(tab)

	if len(tab) != 0 {
		t.Fail()
	}
}
