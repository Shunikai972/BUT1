package fibonnaci

import "testing"

func TestFibonnacci0(t *testing.T) {
	fib := fibonnaci(0)
	if fib != 0 {
		t.Fatal("pb avec fib de 0")
	}
}

func TestFibonnacci1(t *testing.T) {
	fib := fibonnaci(1)
	if fib != 1 {
		t.Fatal("pb avec fib de 1")
	}
}

func TestFibonnacci2(t *testing.T) {
	fib := fibonnaci(1)
	if fib != 1 {
		t.Fatal("pb avec fib de 1")
	}
}

func TestFibonnacci15(t *testing.T) {
	fib := fibonnaci(15)
	if fib != 610 {
		t.Fatal("pb avec fib de 1")
	}
}
