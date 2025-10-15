package minimum

func minimum(t []int) int {
	if len(t) == 0 {
		return -1
	}
	var posMin = 0
	for i := 1; i < len(t); i++ {
		if t[i] < t[posMin] {
			posMin = i
		}
	}
	return posMin
}
