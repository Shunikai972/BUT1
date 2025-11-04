package sortlist

import "testing"

func TestEtudiant(t *testing.T) {
	e1 := Element{V: 5}
	e2 := Element{V: 3, Next: &e1}
	e3 := Element{V: 2, Next: &e2}
	e4 := Element{V: 4, Next: &e3}
	e5 := Element{V: 1, Next: &e4}
	l := List{Content: &e5}

	ll := l.Sort()

	e := ll.Content
	for i := 1; i <= 5; i++ {
		if e == nil {
			t.Fatal("step1:", i)
		}
		if e.V != i {
			t.Fatal("step2:", i, "->", e.V, "instead of", i)
		}
		e = e.Next
	}

	e = l.Content
	tab := []int{1, 4, 2, 3, 5}
	for i := 0; i < 5; i++ {
		if e == nil {
			t.Fatal("step3:", i)
		}
		if e.V != tab[i] {
			t.Fatal("step4:", i)
		}
		e = e.Next
	}
}

// Ajouté après le test machine
func Test2Etudiant(t *testing.T) {
	e1 := Element{V: 5}
	e2 := Element{V: 3, Next: &e1}
	e3 := Element{V: 2, Next: &e2}
	e4 := Element{V: 4, Next: &e3}
	e5 := Element{V: 1, Next: &e4}
	e6 := Element{V: 6, Next: &e5}
	e7 := Element{V: 8, Next: &e6}
	e8 := Element{V: 7, Next: &e7}
	e9 := Element{V: 9, Next: &e8}
	e10 := Element{V: 10, Next: &e9}
	l := List{Content: &e10}

	ll := l.Sort()

	e := ll.Content
	for i := 1; i <= 10; i++ {
		if e == nil {
			t.Fatal("step1:", i)
		}
		if e.V != i {
			t.Fatal("step2:", i, "->", e.V, "instead of", i)
		}
		e = e.Next
	}

	e = l.Content
	tab := []int{10, 9, 7, 8, 6, 1, 4, 2, 3, 5}
	for i := 0; i < 10; i++ {
		if e == nil {
			t.Fatal("step3:", i)
		}
		if e.V != tab[i] {
			t.Fatal("step4:", i)
		}
		e = e.Next
	}
}

func Test3(t *testing.T) {
	l := List{}
	ll := l.Sort()

	if ll.Content != nil {
		t.FailNow()
	}
}
