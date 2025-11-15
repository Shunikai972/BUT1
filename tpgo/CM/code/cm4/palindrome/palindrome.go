package palindrome

func estPalindrome(s string) (b bool) {
	if len(s) <= 1 {
		return true
	}
	if s[0] != s[len(s)-1] {
		return false
	}
	return estPalindrome(s[1 : len(s)-1])
}
