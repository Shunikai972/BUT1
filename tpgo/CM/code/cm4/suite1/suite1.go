package suite1

func suite1(n int) int {
	if n == 0 {
		return 0
	}
	return 2*suite1(n-1) + 1
}
