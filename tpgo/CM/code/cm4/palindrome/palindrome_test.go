package palindrome

import "testing"

func TestPalindromeOK1(t *testing.T) {
	pal := estPalindrome("")
	if !pal {
		t.Fatal("pb avec \"\"")
	}
}

func TestPalindromeOK2(t *testing.T) {
	pal := estPalindrome("a")
	if !pal {
		t.Fatal("pb avec \"a\"")
	}
}

func TestPalindromeOK3(t *testing.T) {
	pal := estPalindrome("kayak")
	if !pal {
		t.Fatal("pb avec \"kayak\"")
	}
}

func TestPalindromeKO1(t *testing.T) {
	pal := estPalindrome("ayak")
	if pal {
		t.Fatal("pb avec \"ayak\"")
	}
}
