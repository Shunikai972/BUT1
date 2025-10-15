package sequentielle

import "slices"

func rechercheElement(t []int, x int) (bool, int) {
	/*for i := 0; i < len(t); i++ {
		if t[i] == x {
			return true, i
		}
		i++
	}
	return false, -1*/

	pos := slices.Index(t, x)
	if pos == -1 {
		return false, -1
	}
	return true, pos
}
