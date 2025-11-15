package iterative

func rechercheDicho(x int, t []int) int {
	debut := 0
	fin := len(t)
	var milieu int
	for debut < fin {
		milieu = (debut + fin) / 2
		if t[milieu] == x {
			return milieu
		} else {
			if t[milieu] > x {
				fin = milieu
			} else {
				debut = milieu + 1
			}
		}
	}
	return -1
}
