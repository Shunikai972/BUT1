package factorielle

func factorielle(n int) int {
	if n < 2 {
		return 1
	}
	return n * factorielle(n-1)
}
