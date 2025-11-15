package recursive

func rechercheDicho(x int, t []int) int {
	if len(t) == 0 {
		return -1
	}
	milieu := len(t) / 2
	if t[milieu] == x {
		return milieu
	}
	if t[milieu] > x {
		return rechercheDicho(x, t[:milieu])
	}
	tmp := rechercheDicho(x, t[milieu+1:])
	if tmp == -1 {
		return -1
	}
	return milieu + tmp
}
